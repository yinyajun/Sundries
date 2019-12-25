#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/15 17:53
# @Author  : Yajun Yin
# @Note    :
import jieba
import tensorflow as tf
import seq_2_seq

CHECKPOINT_PATH = "D:\work\python_test\save\seq2seq_ckpt-4400"

HIDDEN_SIZE = 1024
NUM_LAYERS = 2
SRC_VOCAB_SIZE = 10000
TRG_VOCAB_SIZE = 4000
SHARE_EMB_AND_SOFTMAX = True

SOS_ID = 1
EOS_ID = 2


class NMTModel(object):
    def __init__(self):
        # encoder & decoder
        self.enc_cell = tf.nn.rnn_cell.MultiRNNCell(
            [tf.nn.rnn_cell.BasicLSTMCell(HIDDEN_SIZE) for _ in range(NUM_LAYERS)])
        self.dec_cell = tf.nn.rnn_cell.MultiRNNCell(
            [tf.nn.rnn_cell.BasicLSTMCell(HIDDEN_SIZE) for _ in range(NUM_LAYERS)])

        # embedding vector for source and target language
        self.src_embedding = tf.get_variable("src_emb", [SRC_VOCAB_SIZE, HIDDEN_SIZE])
        self.trg_embedding = tf.get_variable("trg_emb", [TRG_VOCAB_SIZE, HIDDEN_SIZE])

        if SHARE_EMB_AND_SOFTMAX:
            self.softmax_weight = tf.transpose(self.trg_embedding)
        else:
            self.softmax_weight = tf.get_variable("weight", [HIDDEN_SIZE, TRG_VOCAB_SIZE])
        self.softmax_bias = tf.get_variable("softmax_bias", [TRG_VOCAB_SIZE])

    def inference(self, src_input):
        src_size = tf.convert_to_tensor([len(src_input)], dtype=tf.int32)
        src_input = tf.convert_to_tensor([src_input], dtype=tf.int32)
        src_emb = tf.nn.embedding_lookup(self.src_embedding, src_input)
        MAX_DEC_LEN = 100

        with tf.variable_scope("encoder"):
            enc_outputs, enc_state = tf.nn.dynamic_rnn(self.enc_cell, inputs=src_emb, sequence_length=src_size,
                                                       dtype=tf.float32)

        with tf.variable_scope("decoder/rnn/multi_rnn_cell"):
            init_array = tf.TensorArray(dtype=tf.int32, size=0, dynamic_size=True, clear_after_read=False)
            init_array = init_array.write(0, SOS_ID)
            init_loop_var = (enc_state, init_array, 0)

            def continue_loop_cond(state, trg_ids, step):
                return tf.reduce_all(tf.logical_and(
                    tf.not_equal(trg_ids.read(step), EOS_ID),
                    tf.less(step, MAX_DEC_LEN - 1)
                ))

            def loop_body(state, trg_ids, step):
                trg_input = [trg_ids.read(step)]
                trg_emb = tf.nn.embedding_lookup(self.trg_embedding, trg_input)

                dec_outputs, next_state = self.dec_cell.call(state=state, inputs=trg_emb)
                output = tf.reshape(dec_outputs, [-1, HIDDEN_SIZE])
                logits = (tf.matmul(output, self.softmax_weight) + self.softmax_bias)

                next_id = tf.argmax(logits, axis=1, output_type=tf.int32)
                trg_ids = trg_ids.write(step + 1, next_id[0])
                return next_state, trg_ids, step + 1

            state, trg_ids, step = tf.while_loop(continue_loop_cond, loop_body, init_loop_var)
            return trg_ids.stack()


def sentence_to_ids(sentence, voc_file):
    with open(voc_file, 'r', encoding='utf-8') as f_vocab:
        vocab = [w.strip() for w in f_vocab.readlines()]
    word_to_id = {v: k for k, v in enumerate(vocab)}
    get_id = lambda w: word_to_id[w] if w in word_to_id else word_to_id["<unk>"]
    test_sentence = jieba.cut(sentence)
    test_sentence = [x for x in test_sentence if x != ' ']
    sentence_id = [get_id(w) for w in test_sentence]
    # add SOS
    sentence_id = [SOS_ID] + sentence_id
    return sentence_id


def ids_to_sentence(ids, voc_file):
    with open(voc_file, 'r', encoding='utf-8') as f_vocab:
        vocab = [w.strip() for w in f_vocab.readlines()]
    id_to_word = {k: v for k, v in enumerate(vocab)}
    sentence = [id_to_word[i] for i in ids]
    print()
    sentence = ''.join(sentence)
    return sentence


def main():
    with tf.variable_scope("mnt_model", reuse=None):
        model = NMTModel()
    test = "This is a test."
    test_sentence = sentence_to_ids(test, "corpus/vocab.train.txt.en")
    print(test_sentence)
    output_op = model.inference(test_sentence)
    sess = tf.Session()
    saver = tf.train.Saver()
    saver.restore(sess, CHECKPOINT_PATH)

    output = sess.run(output_op)
    print(output)
    output = ids_to_sentence(output, "corpus/vocab.train.txt.zh")
    print(output)
    sess.close()


if __name__ == '__main__':
    main()
