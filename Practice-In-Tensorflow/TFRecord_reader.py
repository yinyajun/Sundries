#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/2 14:07
# @Author  : Yajun Yin
# @Note    : copied in <Tensorflow实战Google深度学习框架>

import tensorflow as tf

# 创建一个reader来读取TFRecord中的样例
reader = tf.TFRecordReader()

# 创建一个临时队列用于维护输入文件列表
filename_queue = tf.train.string_input_producer(["E:/path/to/output.tfrecords"])

# 从文件中读取一个样例。read_up_to函数用于一次性读取多个样例
_, serialized_example = reader.read(filename_queue)
# 解析读入的样例。如果需要解析多个样例，用parse_example
features = tf.parse_single_example(serialized_example, features={
    # TensorFlow提供了两种属性解析方法
    # 1.tf.FixedLenFeature解析结果为Tensor.
    # 2.tf.VarLenFeature,这种解析方法解析为一个SparseTensor,用于处理稀疏矩阵
    # 格式需要与上面的写入数据的格式相一致
    'image_raw': tf.FixedLenFeature([], tf.string),
    'pixels': tf.FixedLenFeature([], tf.int64),
    'label': tf.FixedLenFeature([], tf.int64),
})

# tf.decode_raw将字符串解析为对应的像素数组
images = tf.decode_raw(features['image_raw'], tf.uint8)
labels = tf.cast(features['label'], tf.int32)
pixels = tf.cast(features['pixels'], tf.int32)

sess = tf.Session()
# 启动多线程处理输入数据
coord = tf.train.Coordinator()
threads = tf.train.start_queue_runners(sess=sess, coord=coord)

# 每次读取一个样例，当所有样例读取完之后，从头读取。
for i in range(10):
    image, label, pixel = sess.run([images, labels, pixels])
    print(label, pixel, images)
