#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/19 15:53
# @Author  : Yajun Yin
# @Note    :
import os
import tensorflow as tf
import mnist_inference

from tensorflow.contrib.tensorboard.plugins import projector
from tensorflow.examples.tutorials.mnist import input_data

BATCH_SIZE = 100
LEARNING_RATE_BASE = 0.8
LEARNING_RATE_DECAY = 0.99  # 指数衰减学习率的衰减率
REGULARIZATION_RATE = 0.0001
TRAINING_STEPS = 10000
MOVING_AVERAGE_DECAY = 0.99  # 滑动平均衰减率

LOG_DIR = 'logs'
SPRITE_FILE = 'mnist_sprite.jpg'
META_FILE = 'mnist_meta.tsv'
TENSOR_NAME = 'FINAL_LOGITS'


def train(mnist):
    with tf.name_scope("input"):
        x = tf.placeholder(tf.float32, [None, mnist_inference.INPUT_NODE], "x-input")
        y_ = tf.placeholder(tf.float32, [None, mnist_inference.OUTPUT_NODE], "y-input")

    regularizer = tf.contrib.layers.l2_regularizer(REGULARIZATION_RATE)
    y = mnist_inference.inference(x, regularizer)
    global_step = tf.Variable(0, trainable=False)

    with tf.name_scope("moving_average"):
        variable_averages = tf.train.ExponentialMovingAverage(MOVING_AVERAGE_DECAY, global_step)
        variables_average_op = variable_averages.apply(tf.trainable_variables())

    with tf.name_scope("loss_function"):
        cross_entropy = tf.nn.sparse_softmax_cross_entropy_with_logits(labels=tf.argmax(y_, 1), logits=y)
        cross_entropy_mean = tf.reduce_mean(cross_entropy)
        loss = cross_entropy_mean + tf.add_n(tf.get_collection('losses'))

    with tf.name_scope("train_step"):
        learning_rate = tf.train.exponential_decay(
            LEARNING_RATE_BASE,
            global_step,
            mnist.train.num_examples / BATCH_SIZE,
            LEARNING_RATE_DECAY
        )
        train_step = tf.train.GradientDescentOptimizer(learning_rate).minimize(loss, global_step=global_step)

        with tf.control_dependencies([train_step, variables_average_op]):
            train_op = tf.no_op(name='train')

    # train
    with tf.Session() as sess:
        init_op = tf.global_variables_initializer()
        sess.run(init_op)

        for i in range(TRAINING_STEPS):
            xs, ys = mnist.train.next_batch(BATCH_SIZE)
            _, loss_value, step = sess.run([train_op, loss, global_step], feed_dict={x: xs, y_: ys})

            if i % 1000 == 0:
                print("After %d training steps, loss on training batch is %g." % (step, loss_value))
                # global_step参数可以让每个被保存模型的文件名末尾加上训练的轮数，model.ckpt-1000

        final_result = sess.run(y, feed_dict={x: mnist.test.images})

    return final_result


def visualisation(final_result):
    y = tf.Variable(final_result, name=TENSOR_NAME)
    summary_writer = tf.summary.FileWriter(LOG_DIR)

    config = projector.ProjectorConfig()
    embedding = config.embeddings.add()
    embedding.tensor_name = y.name

    embedding.metadata_path = META_FILE
    embedding.sprite.image_path = SPRITE_FILE
    embedding.sprite.single_image_dim.extend([28, 28])

    projector.visualize_embeddings(summary_writer, config)

    sess = tf.InteractiveSession()
    sess.run(tf.global_variables_initializer())
    saver = tf.train.Saver()
    saver.save(sess, os.path.join(LOG_DIR, "model"), TRAINING_STEPS)

    summary_writer.close()


def main(argv=None):
    mnist = input_data.read_data_sets('data/mnist_data', one_hot=True)
    final_result = train(mnist)
    visualisation(final_result)


if __name__ == '__main__':
    main()
