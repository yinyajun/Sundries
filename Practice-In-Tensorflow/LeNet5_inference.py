#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/2 13:17
# @Author  : Yajun Yin
# @Note    : copied in <Tensorflow实战Google深度学习框架>

import tensorflow as tf

# # 卷积层只和过滤器的尺寸，深度和当前节点矩阵的深度有关。
# # filer_weight的参数是4维张量:1-2/过滤器尺寸；3：当前层深度；4：过滤器深度
# filter_weight = tf.get_variable("weight", [5, 5, 3, 16], initializer=tf.truncated_normal_initializer(stddev=0.1))
# biases = tf.get_variable("biases", [16], initializer=tf.constant_initializer(0.1))
#
# conv = tf.nn.conv2d(input, filter_weight, strides=[1, 1, 1, 1], padding='SAME')
# bias = tf.nn.bias_add(conv, biases)
#
# actived_conv = tf.nn.relu(bias)
#
# pool = tf.nn.max_pool(actived_conv, ksize=[1, 3, 3, 1], strides=[1, 2, 2, 1], padding='SAME')
INPUT_NODE = 784
OUTPUT_NODE = 10

IMAGE_SIZE = 28
NUM_CHANNELS = 1
NUM_LABELS = 10

CONV1_DEEP = 32
CONV1_SIZE = 5

CONV2_DEEP = 64
CONV2_SIZE = 5

FC_SIZE = 512


# train参数用来区别训练和测试，因为训练时用到dropout
def inference(input_tensor, train, regularizer):
    with tf.variable_scope('layer1-conv1'):
        conv1_weights = tf.get_variable("weights", [CONV1_SIZE, CONV1_SIZE, NUM_CHANNELS, CONV1_DEEP],
                                        initializer=tf.truncated_normal_initializer(stddev=0.1))
        conv1_bias = tf.get_variable("biases", [CONV1_DEEP], initializer=tf.constant_initializer(0.0))

        conv1 = tf.nn.conv2d(input_tensor, conv1_weights, strides=[1, 1, 1, 1], padding='SAME')
        relu1 = tf.nn.relu(tf.nn.bias_add(conv1, conv1_bias))

    with tf.name_scope('layer2-pool'):
        pool1 = tf.nn.max_pool(relu1, ksize=[1, 2, 2, 1], strides=[1, 2, 2, 1], padding="SAME")

    with tf.variable_scope("layer3-conv2"):
        conv2_weights = tf.get_variable("weights", [CONV2_SIZE, CONV2_SIZE, CONV1_DEEP, CONV2_DEEP],
                                        initializer=tf.truncated_normal_initializer(stddev=0.1))
        conv2_bias = tf.get_variable("biases", [CONV2_DEEP], initializer=tf.constant_initializer(0.0))
        conv2 = tf.nn.conv2d(pool1, conv2_weights, strides=[1, 1, 1, 1], padding="SAME")
        relu2 = tf.nn.relu(tf.nn.bias_add(conv2, conv2_bias))

    with tf.name_scope('layer4-pool'):
        pool2 = tf.nn.max_pool(relu2, ksize=[1, 2, 2, 1], strides=[1, 2, 2, 1], padding="SAME")
        # 自动得到当前层的size，pool_shape[0]= batch_size。
        pool_shape = pool2.get_shape().as_list()
        nodes = pool_shape[1] * pool_shape[2] * pool_shape[3]
        # 因为全连接层的输入格式为向量，这里reshape
        reshaped = tf.reshape(pool2, [pool_shape[0], nodes])

    with tf.variable_scope("layer5-fc1"):
        fc1_weight = tf.get_variable("weights", [nodes, FC_SIZE],
                                     initializer=tf.truncated_normal_initializer(stddev=0.1))
        # 只有全连接层需要正则化
        if regularizer is not None:
            tf.add_to_collection("losses", regularizer(fc1_weight))
        fc1_bias = tf.get_variable("biases", [FC_SIZE], initializer=tf.constant_initializer(0.1))

        fc1 = tf.nn.relu(tf.matmul(reshaped, fc1_weight) + fc1_bias)
        if train:
            fc1 = tf.nn.dropout(fc1, 0.5)

    with tf.variable_scope("layer6-fc2"):
        fc2_weight = tf.get_variable("weights", [FC_SIZE, NUM_LABELS],
                                     initializer=tf.truncated_normal_initializer(stddev=0.1))
        if regularizer is not None:
            tf.add_to_collection("losses", regularizer(fc2_weight))
        fc2_bias = tf.get_variable("biases", [NUM_LABELS], initializer=tf.constant_initializer(0.1))
        logit = tf.matmul(fc1, fc2_weight) + fc2_bias

    return logit
