#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/2 14:07
# @Author  : Yajun Yin
# @Note    : copied in <Tensorflow实战Google深度学习框架>

import tensorflow as tf
from preprocessing_image import preprocess_for_train

train_files = tf.train.match_filenames_once("train_file_*")
test_files = tf.train.match_filenames_once("test_file_*")


# parse TFRecord
def parser(record):
    features = tf.parse_single_example(record, features={'image': tf.FixedLenFeature([], tf.string),
                                                         'label': tf.FixedLenFeature([], tf.int64),
                                                         'height': tf.FixedLenFeature([], tf.int64),
                                                         'width': tf.FixedLenFeature([], tf.int64),
                                                         'channel': tf.FixedLenFeature([], tf.int64)})
    # deserialize string to array
    decoded_image = tf.decode_raw(features['image'], tf.uint8)
    decoded_image.set_shape([features['height'], features['width'], features['channel']])
    label = features['label']
    return decoded_image, label


image_size = 299
batch_size = 100
shuffle_buffer = 10000  # min_after_dequeue
NUM_EPOCHS = 10

# *************** train **************

'''NO.1 step: generate dataset'''
dataset = tf.contrib.data.TFRecordDataset(train_files)
dataset = dataset.map(parser)
# preprocessing
dataset = dataset.map(lambda image, label: (preprocess_for_train(image, image_size, image_size, None), label))
# shuffle
dataset = dataset.shuffle(shuffle_buffer).batch(batch_size)
# repeat for epochs
dataset = dataset.repeat(NUM_EPOCHS)

'''No.2 step: initialize iterator'''
iterator = dataset.make_initializable_iterator()

'''No.3 step: get input from iterator'''
image_batch, label_batch = iterator.get_next()

# neural network
learning_rate = 0.01
logit = inference(image_batch)
loss = calc_loss(logit, label_batch)
train_step = tf.train.GradientDescentOptimizer(learning_rate).minimize(loss)

# *************** test **************
test_dataset = tf.contrib.data.TFRecordDataset(test_files)
test_dataset = test_dataset.map(parser).map(
    lambda image, label: (tf.image.resize_images(image, [image_size, image_size]), label))
test_dataset = test_dataset.batch(batch_size)

test_iterator = test_dataset.make_initializable_iterator()

test_image_batch, test_label_batch = test_iterator.get_next()

test_logit = inference(test_image_batch)
predictions = tf.argmax(test_logit, axis=-1, output_type=tf.int32)

with tf.Session() as sess:
    sess.run([tf.global_variables_initializer(), tf.local_variables_initializer()])

    sess.run(iterator.initializer)

    while True:
        try:
            sess.run(train_step)
        except tf.errors.OutOfRangeError as e:
            break

    sess.run(test_iterator.initializer)

    test_results = []
    test_labels = []
    while True:
        try:
            pred, label = sess.run([predictions, test_label_batch])
            test_results.extend(pred)
            test_labels.extend(label)
        except tf.errors.OutOfRangeError:
            break
    correct = [float(y == y_) for (y, y_) in zip(test_results, test_labels)]
    accuracy = sum(correct) / len(correct)
    print("Test accuracy is：", accuracy)
