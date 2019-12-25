#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/2 14:07
# @Author  : Yajun Yin
# @Note    : copied in <Tensorflow实战Google深度学习框架>
import os
import tensorflow as tf
from tensorflow.examples.tutorials.mnist import input_data
import mnist_inference
import numpy as np

BATCH_SIZE = 100
LEARNING_RATE_BASE = 0.01
LEARNING_RATE_DECAY = 0.99  # 指数衰减学习率的衰减率
REGULARIZATION_RATE = 0.0001
TRAINING_STEPS = 30000
MOVING_AVERAGE_DECAY = 0.99  # 滑动平均衰减率
# 模型保存路径和文件名
MODEL_SAVE_PATH = "/path/to/model/"
MODEL_NAME = "model.ckpt"


def train(mnist):
    x = tf.placeholder(tf.float32, [BATCH_SIZE,
                                    mnist_inference.IMAGE_SIZE,
                                    mnist_inference.IMAGE_SIZE,
                                    mnist_inference.NUM_CHANNELS], "x-input")

    # x = tf.placeholder(tf.float32, [None, mnist_inference.INPUT_NODE], "x-input")
    y_ = tf.placeholder(tf.float32, [None, mnist_inference.OUTPUT_NODE], "y-input")

    regularizer = tf.contrib.layers.l2_regularizer(REGULARIZATION_RATE)

    y = mnist_inference.inference(x, True, regularizer)

    global_step = tf.Variable(0, trainable=False)

    variable_averages = tf.train.ExponentialMovingAverage(MOVING_AVERAGE_DECAY, global_step)

    variables_average_op = variable_averages.apply(tf.trainable_variables())

    cross_entropy = tf.nn.sparse_softmax_cross_entropy_with_logits(labels=tf.argmax(y_, 1), logits=y)
    cross_entropy_mean = tf.reduce_mean(cross_entropy)

    loss = cross_entropy_mean + tf.add_n(tf.get_collection("losses"))

    learning_rate = tf.train.exponential_decay(
        LEARNING_RATE_BASE,
        global_step,
        mnist.train.num_examples / BATCH_SIZE,
        LEARNING_RATE_DECAY
    )
    train_step = tf.train.GradientDescentOptimizer(learning_rate).minimize(loss, global_step=global_step)

    with tf.control_dependencies([train_step, variables_average_op]):
        train_op = tf.no_op(name='train')

    # 持久化
    saver = tf.train.Saver()
    with tf.Session() as sess:
        init_op = tf.global_variables_initializer()
        sess.run(init_op)

        for i in range(TRAINING_STEPS):
            xs, ys = mnist.train.next_batch(BATCH_SIZE)
            reshaped_xs = np.reshape(xs, (BATCH_SIZE,
                                 mnist_inference.IMAGE_SIZE,
                                 mnist_inference.IMAGE_SIZE,
                                 mnist_inference.NUM_CHANNELS))
            _, loss_value, step = sess.run([train_op, loss, global_step], feed_dict={x: reshaped_xs, y_: ys})

            if i % 1000 == 0:
                print("After %d training steps, loss on training batch is %g." % (step, loss_value))
                # global_step参数可以让每个被保存模型的文件名末尾加上训练的轮数，model.ckpt-1000
                saver.save(sess, os.path.join(MODEL_SAVE_PATH, MODEL_NAME), global_step=global_step)


def main(argv=None):
    mnist = input_data.read_data_sets("/path/to/MNIST_data", one_hot=True)
    train(mnist)


if __name__ == '__main__':
    tf.app.run()
