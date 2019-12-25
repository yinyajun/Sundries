#!/usr/bin/env python
# -*- coding: utf-8 -*-
import tensorflow as tf
import numpy as np
import os
from tqdm import tqdm
from scipy import misc

POS_NUM = 12
NEG_NUM = 3

POS_SAVE_DIR = './data/new_positive/'
NEG_SAVE_DIR = './data/new_negative/'
POS_DIR = './data/train/positive'
NEG_DIR = './data/train/negative'


# 给定一张图像， 随机调整图像的色彩。
def distort_color(image, color_ordering=0):
    if color_ordering == 0:
        image = tf.image.random_brightness(image, max_delta=8. / 255.)  # 亮度
        image = tf.image.random_saturation(image, lower=0.95, upper=1.05)  # 饱和度
        image = tf.image.random_hue(image, max_delta=0.02)  # 色相
        image = tf.image.random_contrast(image, lower=0.95, upper=1.05)  # 对比度
    elif color_ordering == 1:
        image = tf.image.random_brightness(image, max_delta=8. / 255.)  # 亮度
        image = tf.image.random_hue(image, max_delta=0.02)  # 色相
        image = tf.image.random_saturation(image, lower=0.95, upper=1.1)  # 饱和度
        image = tf.image.random_contrast(image, lower=0.95, upper=1.05)  # 对比度
    elif color_ordering == 2:
        image = tf.image.random_saturation(image, lower=0.95, upper=1.05)  # 饱和度
        image = tf.image.random_brightness(image, max_delta=8. / 255.)  # 亮度
        image = tf.image.random_hue(image, max_delta=0.02)  # 色相
        image = tf.image.random_contrast(image, lower=0.95, upper=1.05)  # 对比度
    elif color_ordering == 3:
        image = tf.image.random_hue(image, max_delta=0.02)  # 色相
        image = tf.image.random_saturation(image, lower=0.95, upper=1.05)  # 饱和度
        image = tf.image.random_brightness(image, max_delta=8. / 255.)  # 亮度
        image = tf.image.random_contrast(image, lower=0.95, upper=1.05)  # 对比度
    return tf.clip_by_value(image, 0.0, 1.0)  # 色彩调整的API可能导致像素的实数值超出0.0-1.0的范围，截断一个张量


def preprocess_for_train(image, height, width, bbox):
    # 如果没有提供标注框，则认为整个图像就是需要关注的部分
    if bbox is None:
        bbox = tf.constant([0.0, 0.0, 1.0, 1.0], dtype=tf.float32, shape=[1, 1, 4])
    # 转换图像张量的类型
    if image.dtype != tf.float32:
        image = tf.image.convert_image_dtype(image, dtype=tf.float32)
    # 随机截取图像，减少需要关注的物体大小对图像识别的影响
    bbox_begin, bbox_size, _ = tf.image.sample_distorted_bounding_box(tf.shape(image),
                                                                      bounding_boxes=bbox,
                                                                      min_object_covered=0.98,
                                                                      aspect_ratio_range=[0.54, 0.59],
                                                                      )
    distort_image = tf.slice(image, bbox_begin, bbox_size)
    # 将随机截图的图像调整为神经网络输入层的大小。大小调整的算法是随机的
    distort_image = tf.image.resize_images(distort_image, [height, width], method=np.random.randint(4))
    # 随机左右翻转图像
    distort_image = tf.image.random_flip_left_right(distort_image)
    # 使用一种随机的顺序调整图像色彩
    distort_image = distort_color(distort_image, np.random.randint(4))
    return distort_image


def image_augment(image, height, width, num):
    boxes = tf.constant([[[0.50, 0.15, 1.0, 0.85]]])
    res = []
    for _ in range(num):
        res.append(preprocess_for_train(image, height, width, boxes))
    return res


def generate_image(original_dir, save_dir, num):
    files = list(os.walk(original_dir))[0][2]
    for f in tqdm(files):
        image_path = os.path.join(original_dir, f)
        f_name = f.split('.')[0]
        # avoid memory leak
        tf.reset_default_graph()
        graph = tf.Graph()
        with graph.as_default() as g:
            image_raw_data = tf.gfile.FastGFile(image_path, 'rb').read()
            ima_data = tf.image.decode_jpeg(image_raw_data, channels=3)
            ret = image_augment(ima_data, 512, 288, num)
            with tf.Session(config=tf.ConfigProto(inter_op_parallelism_threads=0,
                                                  intra_op_parallelism_threads=0)) as sess:
                for i in range(num):
                    new_img = sess.run(ret[i])
                    new_f_name = f_name + '_new_' + str(i) + '.jpg'
                    new_img_path = os.path.join(save_dir, new_f_name)
                    misc.imsave(new_img_path, new_img)


generate_image(POS_DIR, POS_SAVE_DIR, POS_NUM)
generate_image(NEG_DIR, NEG_SAVE_DIR, NEG_NUM)
