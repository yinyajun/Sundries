#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/1/25 14:27
# @Author  : Yajun Yin
# @Note    :


import tensorflow as tf


def test_embedding_lookup_gradient():
    """ Toy example to show the result of gradient through embedding_lookup ops.
        Conclusion: regularization of embedding lookup can not be set as
        ||embedding_matrix||F, because Optimizer ops will compute gradients
        w.r.t. every indices of embedding matrix and updates.
    """

    x = tf.placeholder(tf.int64, shape=(1,))
    y = tf.placeholder(tf.float32, shape=(1, 1))

    w = tf.get_variable("w1", [10, 1], initializer=tf.truncated_normal_initializer(stddev=0.1))
    xx = tf.nn.embedding_lookup(w, x)

    y_ = 200 * xx + 5
    # print(y_)
    reg = tf.reshape(tf.reduce_sum(tf.multiply(w, w)), [1, 1])
    # print(reg)
    y_ += reg

    loss = tf.losses.mean_squared_error(y_, y)
    opt = tf.train.GradientDescentOptimizer(1.0)
    grads_and_vars = opt.compute_gradients(loss)

    sess = tf.Session()
    sess.run(tf.global_variables_initializer())

    ret = sess.run(grads_and_vars, feed_dict={x: [0], y: [[0.0]]})
    print(tf.get_collection("trainable_variables"))
    print(ret)
