#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2020/1/7 9:43
# @Author  : Yajun Yin
# @Note    :

import tensorflow as tf
import time
from tensorflow.contrib import lookup

num_true = 5
range_max = 50
num_sampled = 8
batch_size = 3

sess = tf.Session()
voc_table = lookup.index_table_from_file('a.csv', 0, 10, 0, key_dtype=tf.int64)
sess.run(tf.tables_initializer())


def parse(lines):
    sparse = tf.string_split([lines], ',')
    values = tf.string_to_number(sparse.values, tf.int64)
    return values


def id_index(values):
    ids = voc_table.lookup(values)
    return ids


def multi_hot(ids, values, dim):
    # validate_indices=true requires indices is increasing
    dense = tf.sparse_to_dense(ids, [dim], values, validate_indices=False)
    return dense


def negative_sample(true_classes, num_sample, min_range, max_range):
    samples = tf.random_uniform([num_sample], min_range, max_range, tf.int64)
    samples = tf.setdiff1d(samples, true_classes).out
    return samples


def process(lines):
    values = parse(lines)
    ids = id_index(values)
    vector = multi_hot(ids, tf.ones_like(ids), 20)
    samples = negative_sample(ids, 8, 1, 21)
    return values, ids, vector, samples, tf.size(samples)


data = tf.data.TextLineDataset("b.csv")
data = data.repeat(20)
data = data.map(process)
data = data.padded_batch(3, ([None], [None], [None], [None], ()))
iterator = data.make_initializable_iterator()
a, b, c, d, e = iterator.get_next()

sess.run(iterator.initializer)
t = time.time()
ret = sess.run([a, b, c, d, e])
for i in ret:
    print(i)
print(time.time() - t)