#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/5 11:56
# @Author  : Yajun Yin
# @Note    :


import tensorflow as tf


def _int64_feature(value):
    return tf.train.Feature(int64_list=tf.train.Int64List(value=[value]))


# 文件数和每个文件的数据量
num_shards = 2
instances_per_shard = 2
for i in range(num_shards):
    filename = ('/path/to/data.tfrecords-%.5d-of-%.5d' % (i, num_shards))
    writer = tf.python_io.TFRecordWriter(filename)
    # 将数据封装成example结构
    for j in range(instances_per_shard):
        example = tf.train.Example(features=tf.train.Features(feature={
            'i': _int64_feature(i),
            'j': _int64_feature(j)}))
        writer.write(example.SerializeToString())
    writer.close()
