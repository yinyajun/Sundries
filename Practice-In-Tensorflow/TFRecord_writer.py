#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/2 14:07
# @Author  : Yajun Yin
# @Note    : copied in <Tensorflow实战Google深度学习框架>
import tensorflow as tf
from tensorflow.examples.tutorials.mnist import input_data
import numpy as np


# 注意Int64List的value参数必须是一个list
# 生成整数型的属性
def _int64_feature(value):
    return tf.train.Feature(int64_list=tf.train.Int64List(value=[value]))


# 生成字符串型的属性
def _bytes_feature(value):
    return tf.train.Feature(bytes_list=tf.train.BytesList(value=[value]))


mnist = input_data.read_data_sets("/path/to/mnist/data", dtype=tf.uint8, one_hot=True)
images = mnist.train.images
# label作为Example的一个属性
labels = mnist.train.labels
# 像素值可以作为Example的一个属性
pixels = images.shape[1]
num_examples = mnist.train.num_examples

# 输出TFRecord文件的地址
# 这里需要写绝对路径，否则报failed to create a newwriteablefile
filename = "E:/path/to/output.tfrecords"

# 创建一个writer来写TFRecord文件
writer = tf.python_io.TFRecordWriter(filename)
for index in range(num_examples):
    # 数组要序列化为字符串
    image_raw = images[index].tostring()
    # 将一个样例转化为Example Protocol Buffer,并将所有信息写入这个结构
    example = tf.train.Example(features=tf.train.Features(feature={
        'pixels': _int64_feature(pixels),
        'label': _int64_feature(np.argmax(labels[index])),
        'image_raw': _bytes_feature(image_raw)
    }))
    # 写入TFRecord
    writer.write(example.SerializeToString())

writer.close()
