#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/15 10:12
# @Author  : Yajun Yin
# @Note    :

# """split words"""
# import jieba
#
# with open("data/en-zh/train.tags.en-zh.zh", 'r', encoding='utf-8') as fin:
#     for line in fin:
#         separated_line = ' '.join(jieba.cut(line))
#         with open('train.txt.zh', 'a', encoding='utf-8') as fout:
#             fout.write(separat"""split words"""ed_line)

"""make dataset"""
import tensorflow as tf
import NLP_dictionarize

MAX_LEN = 50
SOS_ID = 1


def make_dataset(file_path):
    """
    make dataset
    :param file_path: words in file have been mapped to ids
    :return: (sentence ids, sentence_len)
    """
    dataset = tf.contrib.data.TextLineDataset(file_path)
    dataset = dataset.map(lambda string: tf.string_split([string]).values)
    dataset = dataset.map(lambda string: tf.string_to_number(string, tf.int32))
    dataset = dataset.map(lambda x: (x, tf.size(x)))

    return dataset


def make_src_trg_dataset(src_path, trg_path, batch_size):
    """

    :param src_path: source file
    :param trg_path: target file
    :param batch_size:
    :return:
    """
    src_data = make_dataset(src_path)
    trg_data = make_dataset(trg_path)
    dataset = tf.contrib.data.Dataset.zip((src_data, trg_data))

    # filter empty sentence or long sentence
    def filter_length(src_tuple, trg_tuple):
        ((src_input, src_len), (trg_label, trg_len)) = (src_tuple, trg_tuple)
        # 1< len < Max_Len
        src_len_ok = tf.logical_and(tf.greater(src_len, 1), tf.less_equal(src_len, MAX_LEN))
        trg_len_ok = tf.logical_and(tf.greater(trg_len, 1), tf.less_equal(trg_len, MAX_LEN))
        return tf.logical_and(src_len_ok, trg_len_ok)

    dataset = dataset.filter(filter_length)

    def make_trg_input(src_tuple, trg_tuple):
        """
        we have "x y z <eos>", but decoder input need to be "<sos> x y z"
        :param src_tuple:
        :param trg_tuple:
        :return:
        """
        ((src_input, src_len), (trg_label, trg_len)) = (src_tuple, trg_tuple)
        trg_input = tf.concat([[SOS_ID], trg_label[:-1]], axis=0)
        return (src_input, src_len), (trg_input, trg_label, trg_len)

    dataset = dataset.map(make_trg_input)
    dataset = dataset.shuffle(10000)

    # batch
    padded_shapes = ((tf.TensorShape([None]),
                      tf.TensorShape([])),
                     (tf.TensorShape([None]),
                      tf.TensorShape([None]),
                      tf.TensorShape([])))

    batched_dataset = dataset.padded_batch(batch_size, padded_shapes=padded_shapes)
    return batched_dataset


def main(argv):
    split_file_zh = "train.txt.zh"
    split_file_en = "train.txt.en"
    NLP_dictionarize.main([split_file_en, "vocab.en", "train.dict.en"])
    NLP_dictionarize.main([split_file_zh, "vocab.zh", "train.dict.zh"])
