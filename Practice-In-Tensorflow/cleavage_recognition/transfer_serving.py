#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# @Time    : 2019/5/27 19:35
# @Author  : Yajun Yin
# @Note    :

import os
import traceback
import datetime
import requests
import numpy as np
import tensorflow as tf
import pymysql
import logging
import time
import json
from multiprocessing import Process, Queue, JoinableQueue
from tensorflow.python.platform import gfile

# log目录
LOG_DIR = ''

# Inception-v3模型瓶颈层的节点个数
BOTTLENECK_TENSOR_SIZE = 2048

# Inception-v3模型中代表瓶颈层结果的张量名称。
# 在谷歌提出的Inception-v3模型中，这个张量名称就是'pool_3/_reshape:0'。
# 在训练模型时，可以通过tensor.name来获取张量的名称。
BOTTLENECK_TENSOR_NAME = 'pool_3/_reshape:0'

# 图像输入张量所对应的名称。
JPEG_DATA_TENSOR_NAME = 'DecodeJpeg/contents:0'

# 下载的谷歌训练好的Inception-v3模型文件目录
MODEL_DIR = 'inception_model/'

# 下载的谷歌训练好的Inception-v3模型文件名
MODEL_FILE = 'classify_image_graph_def.pb'

# LR模型目录
LR_MODEL_DIR = 'lr_model'

# LR模型文件名
LR_MODEL_FILE = 'transfer_coef.npz'

IMG_NUMS = 1

# 分数门限
THRESH = 0.45

# 并发数
CONCURRENCY_NUM = 5

IMG_LIST_API = None  # 不能提供


def init_logger(name):
    logger = logging.getLogger(name)
    logger.setLevel(level=logging.INFO)
    handler = logging.FileHandler(LOG_DIR + "{name}.log".format(name=name))
    handler.setLevel(logging.INFO)
    formatter = logging.Formatter(fmt='%(levelname)s %(asctime)s %(message)s', datefmt='%Y-%m-%d %H:%M:%S')
    handler.setFormatter(formatter)
    logger.addHandler(handler)
    return logger


def get_id_list():
    return list(np.random.randint(0, 100, np.random.randint(10, 20)))


def get_image_url(id):
    params = {id: 'id'}
    response = requests.get(IMG_LIST_API, params=params)
    if response.status_code != 200:
        return None
    d = json.loads(response.text)
    try:
        data = d["data"]
        img_url = data['url']
        return img_url
    except KeyError:
        return None


def get_image_data(url):
    try:
        response = requests.get(url, timeout=(1, 3))  # timeout for connect and read
    except:
        raise ConnectionError
    if response.status_code != 200:
        raise ConnectionError
    content = response.content
    if content is None or len(content) == 0:
        raise ConnectionError
    return content


def get_bottleneck(sess, image_data, bottleneck_tensor, jpeg_data_tensor):
    bottleneck_values = sess.run(bottleneck_tensor, feed_dict={jpeg_data_tensor: image_data})
    bottleneck_values = np.squeeze(bottleneck_values)
    return bottleneck_values


def load_inception_model():
    with gfile.FastGFile(os.path.join(MODEL_DIR, MODEL_FILE), 'rb') as f:
        graph_def = tf.GraphDef()
        graph_def.ParseFromString(f.read())
        bottleneck_tensor, jpeg_data_tensor = tf.import_graph_def(graph_def,
                                                                  return_elements=[BOTTLENECK_TENSOR_NAME,
                                                                                   JPEG_DATA_TENSOR_NAME])
    return bottleneck_tensor, jpeg_data_tensor


def load_lr_model():
    lr_model = np.load(os.path.join(LR_MODEL_DIR, LR_MODEL_FILE))
    w = lr_model['w']
    b = lr_model['b']
    return w, b


def softmax(logits):
    a = np.exp(logits)
    b = np.sum(a)
    return a / b


def predict(inception_output, w, b):
    logits = inception_output.dot(w) + b
    # ['positive', 'negative']
    predictions = softmax(logits)
    return predictions[0]


def save_ret_to_mysql(ret):
    t = datetime.datetime.today().strftime('%Y%m')
    sql = "insert into table{p} (live_id, live_uid, score) values (%s, %s, %s)".format(p=t)
    params = [(r[0], r[1], "%0.4f" % r[2]) for r in ret]
    MYSQL = {}
    db = pymysql.connect(**MYSQL)
    try:
        with db.cursor() as cursor:
            cursor.executemany(sql, params)
            db.commit()
    except pymysql.Error as e:
        exec_str = traceback.format_exc()
        logger.error('save failed, %s, %s' % e, exec_str)
        db.rollback()
        raise pymysql.Error
    finally:
        db.close()


class JobProcess(Process):
    def __init__(self, task_queue, res_queue):
        Process.__init__(self)  # 父类必须第一个init
        self.task_queue = task_queue  # 生产者队列
        self.res_queue = res_queue  # 消费者队列
        self.logger = None

    def run(self):
        # 子进程的入口
        self.logger = init_logger(str(os.getpid()))  # 每个子进程的日志
        # 先加载模型
        w, b = load_lr_model()
        bottleneck_tensor, jpeg_data_tensor = load_inception_model()
        # 一个子进程只起一个session，因为没有新加入的Op，所以不会有内存泄漏的风险
        with tf.Session() as sess:
            while True:
                id = self.task_queue.get()  # 阻塞程序，时刻监听队列是否有消息
                s = float(self.get_score(id, w, b, bottleneck_tensor, jpeg_data_tensor, sess))
                if s >= THRESH:
                    self.res_queue.put((id, s))
                self.task_queue.task_done()  # 标志队列中的任务完成，配合queue.join

    def get_score(self, id, w, b, output_tensor, input_tensor, sess):
        start = time.time()
        url = get_image_url(id)
        t1 = time.time() - start
        if url is None:
            return -1.0
        s2 = time.time()
        try:
            image_data = get_image_data(url)
            t2 = time.time() - s2
            s3 = time.time()
            bottleneck_values = get_bottleneck(sess, image_data, output_tensor, input_tensor)
            t3 = time.time() - s3
            s4 = time.time()
            prediction = predict(bottleneck_values, w, b)
            t4 = time.time() - s4
            msg = "get_score ok, id: {id}, score: {score:0.4f}, {t1:0.4f}, {t2:0.4f}, {t3:0.4f}, {t4:0.4f}, " \
                  "{t5:0.4f}".format(id=id, url=url, score=prediction, t1=t1, t2=t2, t3=t3, t4=t4,
                                     t5=time.time() - start)
            self.logger.info(msg)
            return prediction
        except ConnectionError:
            return -2.0


def main(process_num):
    task_queue = JoinableQueue()  # 实现了task_done和join的方法
    res_queue = Queue()

    # set process pool
    for _ in range(process_num):
        p = JobProcess(task_queue, res_queue)
        p.daemon = True  # 子进程无限循环，设为守护进程后，主进程结束，子进程也结束
        p.start()
    # Loop for detecting
    while task_queue.empty():
        start = time.time()

        # get living list
        id_list = get_id_list()
        s2 = time.time()

        # put job into task_queue
        for id in id_list:
            task_queue.put(id)

        task_queue.join()  # 只有队列中的所有任务完成才结束
        t2 = time.time() - s2
        s3 = time.time()
        ret = []
        while not res_queue.empty():
            ret.append(res_queue.get())

        # save ret to mysql
        t3 = time.time() - s3
        s4 = time.time()
        try:
            # save_ret_to_mysql(ret)
            print(ret)
            t4 = time.time() - s4
            t5 = time.time() - start
            msg = "save ok, len:{len}, {t1:0.4f},{t2:0.4f},{t3:0.4f},{t4:0.4f},{t5:0.4f}".format(len=len(ret), t1=t1,
                                                                                                 t2=t2, t3=t3, t4=t4,
                                                                                                 t5=t5)
            logger.info(msg)
        except pymysql.Error as e:
            msg = "save fail, %s" % e
            logger.info(msg)
        # sleep 5s before another detecting epoch
        time.sleep(5)


if __name__ == '__main__':
    t = datetime.datetime.today().strftime('%m%d%H')
    logger = init_logger(t)
    main(CONCURRENCY_NUM)
