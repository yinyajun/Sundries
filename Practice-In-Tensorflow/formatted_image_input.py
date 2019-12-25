#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/6 18:10
# @Author  : Yajun Yin
# @Note    :

import tensorflow as tf
import preprocessing_image

files = tf.train.match_filenames_once("D:/path/to/mnist_output.tfrecords")
filename_queue = tf.train.string_input_producer(files, shuffle=True)

reader = tf.TFRecordReader()
_, serialized_example = reader.read(filename_queue)
features = tf.parse_single_example(serialized_example, features={'image': tf.FixedLenFeature([], tf.string),
                                                                 'label': tf.FixedLenFeature([], tf.int64),
                                                                 'height': tf.FixedLenFeature([], tf.int64),
                                                                 'width': tf.FixedLenFeature([], tf.int64),
                                                                 'channel': tf.FixedLenFeature([], tf.int64)})
image, label = features['image'], features['label']
height, width = features['height'], features['width']
channel = features['channel']

decoded_image = tf.decode_raw(image, tf.uint8)
decoded_image.set_shape([height, width, channel])

distorted_image = preprocessing_image.preprocess_for_train(decoded_image, height, width, None)
image_size = 299

min_after_dequeue = 10000
batch_size = 100
capacity = min_after_dequeue + 3 * batch_size
image_batch, label_batch = tf.train.shuffle_batch([distorted_image, label], batch_size, capacity, min_after_dequeue,
                                                  num_threads=3)

learning_rate = 0.01
logit = inference(image_batch)
loss = calc_loss(logit, label_batch)
train_step = tf.train.GradientDescentOptimizer(learning_rate).minimize(loss)

with tf.Session() as sess:
    sess.run((tf.global_variables_initializer(), tf.local_variables_initializer()))
    coord = tf.train.Coordinator()
    threads = tf.train.start_queue_runners(sess, coord)

    TRAINING_STEPS = 5000
    for i in range(TRAINING_STEPS):
        sess.run(train_step)

    coord.request_stop()
    coord.join(threads)
