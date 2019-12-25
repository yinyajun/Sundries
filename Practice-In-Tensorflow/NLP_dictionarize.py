#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/15 10:34
# @Author  : Yajun Yin
# @Note    :

"""transfer word to its id"""
import collections
from operator import itemgetter
import sys
import os


def make_voc(raw_data):
    dir_name = os.path.dirname(raw_data)

    vocab_name = '.'.join(['vocab'] + os.path.basename(raw_data).split('.'))
    vocab_output = os.path.join(dir_name, vocab_name)
    return vocab_output


def word_cnt(raw_data, voc_size):
    # collections.Counter likes dict
    # it has a useful method: most_common()
    counter = collections.Counter()
    vocab_output = make_voc(raw_data)

    with open(raw_data, "r", encoding="utf-8") as fin:
        for line in fin:
            for word in line.strip().split():
                counter[word] += 1

    sorted_word_to_cnt = sorted(counter.items(), key=itemgetter(1), reverse=True)
    # By this way, [x[0] for x in counter.most_common(1000)] can get the same list
    sorted_words = [x[0] for x in sorted_word_to_cnt]

    # add "<eos>", "<sos>", "<unk>"
    sorted_words = ["<unk>", "<sos>", "<eos>"] + sorted_words
    if len(sorted_words) > voc_size:
        sorted_words = sorted_words[:voc_size]

    with open(vocab_output, "w", encoding="utf-8") as fout:
        for word in sorted_words:
            fout.write(word + '\n')


def dictionarize(raw_data, output):
    vocab_output = make_voc(raw_data)
    # read vocab, map vocab to id
    with open(vocab_output, 'r', encoding="utf-8") as f_vocab:
        vocab = [w.strip() for w in f_vocab.readlines()]
    # zip() can also do the same thing
    word_to_id = {v: k for (k, v) in enumerate(vocab)}

    def get_id(word):
        return word_to_id[word] if word in word_to_id else word_to_id["<unk>"]

    with open(raw_data, 'r', encoding='utf-8') as fin:
        with open(output, 'w', encoding='utf-8') as fout:
            for line in fin:
                words = line.strip().split() + ["<eos>"]
                out_line = ' '.join([str(get_id(w)) for w in words]) + '\n'
                fout.write(out_line)


def main(argv=None):
    if argv is None or len(argv) != 4:
        print("argv: raw_data, output, voc_size")
    if argv is not None:
        print(argv)
        raw_data = argv[1]
        output = argv[2]
        voc_size = int(argv[3])
        word_cnt(raw_data, voc_size)
        dictionarize(raw_data, output)


if __name__ == '__main__':
    main(sys.argv)
