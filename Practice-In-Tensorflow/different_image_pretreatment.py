#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/5 10:07
# @Author  : Yajun Yin
# @Note    :

import matplotlib.pyplot as plt
import tensorflow as tf

image_raw_data = tf.gfile.FastGFile("/path/to/data/lena.jpeg", 'rb').read()

with tf.Session() as sess:
    # decode: tensor
    img_data1 = tf.image.decode_jpeg(image_raw_data)

    # 像素的数据类型由int转为实数
    # 这些都是tensor，由tensor.eval()转为矩阵出图
    img_data = tf.image.convert_image_dtype(img_data1, dtype=tf.float32)
    # print(img_data.eval())
    plt.subplot(3, 3, 1)
    plt.imshow(img_data.eval())

    # resize
    print(img_data.eval().shape)  # (512,512,3)
    resized = tf.image.resize_images(img_data, [300, 300], method=0)
    print(resized.get_shape())  # (300, 300, ?)
    plt.subplot(3, 3, 2)
    plt.imshow(resized.eval())

    # crop or padding
    croped = tf.image.resize_image_with_crop_or_pad(img_data, 300, 300)
    padded = tf.image.resize_image_with_crop_or_pad(img_data, 1000, 1000)
    plt.subplot(3, 3, 3)
    plt.imshow(croped.eval())
    plt.subplot(3, 3, 4)
    plt.imshow(padded.eval())

    # flip
    flipped = tf.image.flip_up_down(img_data)
    plt.subplot(3, 3, 5)
    plt.imshow(flipped.eval())

    flipped = tf.image.flip_left_right(img_data)
    plt.subplot(3, 3, 6)
    plt.imshow(flipped.eval())
    # flipped = tf.image.random_flip_up_down(img_data)

    # lightness
    # 这里使用的tensor是tf.int的图片，因为这个api对tf.float的tensor需要归一化才能处理
    adjusted = tf.image.adjust_brightness(img_data1, -0.5)
    plt.subplot(3, 3, 7)
    plt.imshow(adjusted.eval())

    # contrast
    adjusted = tf.image.random_contrast(img_data1, 0, 5)
    plt.subplot(3, 3, 8)
    plt.imshow(adjusted.eval())

    # whitening
    std = tf.image.per_image_standardization(img_data1)
    # std后有负值，imshow不支持RGB图中有负值的显示

    # bounding box

    print(img_data)
    print(img_data.eval().shape)
    batched = tf.expand_dims(img_data, 0)
    print(batched)
    print(batched.eval().shape)
    boxes = tf.constant([[[0.05, 0.05, 0.9, 0.7], [0.35, 0.47, 0.5, 0.56]]])
    result = tf.image.draw_bounding_boxes(batched, boxes)
    print(result.eval().shape)
    plt.subplot(3, 3, 9)
    plt.imshow(result.eval()[0])

    plt.show()
