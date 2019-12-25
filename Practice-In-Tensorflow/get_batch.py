#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/6 14:41
# @Author  : Yajun Yin
# @Note    :

import tensorflow as tf
import threading

# 使用tf.train.match_filenames_once函数获取符合一个正则表达式的所有文件
# 得到的文件列表可以通过tf.train.string_input_producer函数有效管理
a = threading.active_count()
files = tf.train.match_filenames_once("/path/to/data.tfrecords-*")

# tf.train.string_input_producer函数创建输入队列
# num_epochs：计算一轮后自动停止
filename_queue = tf.train.string_input_producer(files, shuffle=False)

reader = tf.TFRecordReader()
_, serialized_example = reader.read(filename_queue)
features = tf.parse_single_example(serialized_example, features={
    'i': tf.FixedLenFeature([], tf.int64),
    'j': tf.FixedLenFeature([], tf.int64)})

# 文件列表中可以读取单个样例，这些样例通过预处理，组成batch，作为神经网络的输入层
# tf.train.batch和tf.train.shuffle_batch函数将单个样例组成batch
example, label = features['i'], features['j']

batch_size = 3
capacity = 1000 + 3 * batch_size

example_batch, label_batch = tf.train.batch([example, label], batch_size=batch_size, capacity=capacity)
# example_batch, label_batch = tf.train.shuffle_batch([example, label], batch_size=batch_size, capacity=capacity,
#                                                    min_after_dequeue=30)

with tf.Session() as sess:
    tf.global_variables_initializer().run()
    tf.local_variables_initializer().run()

    coord = tf.train.Coordinator()
    threads = tf.train.start_queue_runners(sess, coord)

    for i in range(2):
        cur_example_batch, cur_label_batch = sess.run([example_batch, label_batch])
        print("No. %d batch:" % (i + 1), cur_example_batch, cur_label_batch)

    coord.request_stop()
    coord.join(threads)
