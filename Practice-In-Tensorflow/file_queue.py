#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/5 11:56
# @Author  : Yajun Yin
# @Note    :


import tensorflow as tf

# 使用tf.train.match_filenames_once函数获取符合一个正则表达式的所有文件
# 得到的文件列表可以通过tf.train.string_input_producer函数有效管理
files = tf.train.match_filenames_once("/path/to/data.tfrecords-*")

# tf.train.string_input_producer函数创建输入队列
# num_epochs：计算一轮后自动停止
filename_queue = tf.train.string_input_producer(files, shuffle=False, num_epochs=1)

reader = tf.TFRecordReader()
_, serialized_example = reader.read(filename_queue)
features = tf.parse_single_example(serialized_example, features={
    'i': tf.FixedLenFeature([], tf.int64),
    'j': tf.FixedLenFeature([], tf.int64)})

with tf.Session() as sess:
    init_op = tf.local_variables_initializer()
    init_op.run()

    print(files.eval())

    coord = tf.train.Coordinator()
    threads = tf.train.start_queue_runners(sess, coord)

    for i in range(4):
        print(sess.run([features['i'], features['j']]))
    coord.request_stop()
    coord.join(threads)
