# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/5 15:39
# @Author  : Yajun Yin
# @Note    : copied in <Tensorflow实战Google深度学习框架>


import tensorflow as tf
import numpy as np
import threading
import time


def MyLoop(coord, worker_id):
    while not coord.should_stop():
        # 随机停止所有线程
        if np.random.rand() < 0.1:
            print("Stopping from id: %d\n" % worker_id)
            # 通知其他线程停止
            coord.request_stop()
        else:
            print("Working on id: %d\n" % worker_id)
        time.sleep(1)


coord = tf.train.Coordinator()
# 申明5个线程
threads = [threading.Thread(target=MyLoop, name="Loop thread", args=(coord, i)) for i in range(5)]
for t in threads:
    t.start()
coord.join(threads=threads)

