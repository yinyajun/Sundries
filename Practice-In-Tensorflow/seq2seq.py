#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/15 14:02
# @Author  : Yajun Yin
# @Note    :

import tensorflow as tf
import seq2seq_dataset

SRC_TRAIN_DATA = "corpus/train.en"
TRG_TRAIN_DATA = "corpus/train.zh"
CHECKPOINT_PATH = "save/seq2seq_ckpt"
HIDDEN_SIZE = 1024
NUM_LAYERS = 2
SRC_VOCAB_SIZE = 10000
TRG_VOCAB_SIZE = 4000
BATCH_SIZE = 100
NUM_EPOCH = 5
KEEP_PROB = 0.8
MAX_GRAD_NORM = 5
SHARE_EMB_AND_SOFTMAX = True


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

    def forward(self, src_input, src_size, trg_input, trg_label, trg_size):
        batch_size = tf.shape(src_input)[0]

        src_emb = tf.nn.embedding_lookup(self.src_embedding, src_input)
        trg_emb = tf.nn.embedding_lookup(self.trg_embedding, trg_input)

        src_emb = tf.nn.dropout(src_emb, KEEP_PROB)
        trg_emb = tf.nn.dropout(trg_emb, KEEP_PROB)

        with tf.variable_scope("encoder"):
            enc_outputs, enc_state = tf.nn.dynamic_rnn(self.enc_cell, inputs=src_emb, sequence_length=src_size,
                                                       dtype=tf.float32)

        with tf.variable_scope("decoder"):
            dec_outputs, _ = tf.nn.dynamic_rnn(self.dec_cell, inputs=trg_emb, sequence_length=trg_size,
                                               initial_state=enc_state)

        # calculate perplexity
        output = tf.reshape(dec_outputs, [-1, HIDDEN_SIZE])
        logits = tf.matmul(output, self.softmax_weight) + self.softmax_bias
        loss = tf.nn.sparse_softmax_cross_entropy_with_logits(labels=tf.reshape(trg_label, [-1]), logits=logits)

        label_weights = tf.sequence_mask(trg_size, maxlen=tf.shape(trg_label)[1], dtype=tf.float32)
        label_weights = tf.reshape(label_weights, [-1])
        cost = tf.reduce_sum(loss * label_weights)
        cost_per_token = cost / tf.reduce_sum(label_weights)

        trainable_variables = tf.trainable_variables()
        grads = tf.gradients(cost / tf.to_float(batch_size), trainable_variables)
        grads, _ = tf.clip_by_global_norm(grads, MAX_GRAD_NORM)
        optimizer = tf.train.GradientDescentOptimizer(learning_rate=1.0)
        train_op = optimizer.apply_gradients(zip(grads, trainable_variables))
        return cost_per_token, train_op


def run_epoch(sess, cost_op, train_op, saver, step):
    while True:
        try:
            cost, _ = sess.run([cost_op, train_op])
            if step % 10 == 0:
                print("After %d steps, per token cost is %g" % (step, cost))
            if step % 200 == 0:
                saver.save(sess, CHECKPOINT_PATH, global_step=step)
            step += 1
        except tf.errors.OutOfRangeError:
            break
    return step


def main():
    initializer = tf.random_uniform_initializer(-0.05, 0.05)

    with tf.variable_scope("mnt_model", reuse=None, initializer=initializer):
        train_model = NMTModel()

    data = seq2seq_dataset.make_src_trg_dataset(SRC_TRAIN_DATA, TRG_TRAIN_DATA, BATCH_SIZE)
    iterator = data.make_initializable_iterator()
    (src, src_size), (trg_input, trg_label, trg_size) = iterator.get_next()

    cost_op, train_op = train_model.forward(src, src_size, trg_input, trg_label, trg_size)

    saver = tf.train.Saver()
    step = 0
    with tf.Session() as sess:
        tf.global_variables_initializer().run()
        ckpt = tf.train.get_checkpoint_state(CHECKPOINT_PATH)
        if ckpt and ckpt.model_checkpoint_path:
            saver.restore(sess, ckpt.model_checkpoint_path)
            # model_checkpoint_path: 'D:\\work\\python_test\\save\\seq2seq_ckpt-4400'
            step = ckpt.model_checkpoint_path.split('/')[-1].split('-')[-1]
            step = tf.cast(step, tf.float32)
            for i in range(9010):
                sess.run(iterator.initializer)
                step = run_epoch(sess, cost_op, train_op, saver, step)
        else:
            for i in range(NUM_EPOCH):
                print("In iteration: %d" % (i + 1))
                sess.run(iterator.initializer)
                step = run_epoch(sess, cost_op, train_op, saver, step)


if __name__ == '__main__':
    main()

# After 3720 steps, per token cost is 3.76949
# After 3730 steps, per token cost is 3.71416
# After 3740 steps, per token cost is 3.79179
# After 3750 steps, per token cost is 3.73866
# After 3760 steps, per token cost is 3.81567
# After 3770 steps, per token cost is 3.7651
# After 3780 steps, per token cost is 3.52548
# After 3790 steps, per token cost is 3.69416
# After 3800 steps, per token cost is 3.5164
# After 3810 steps, per token cost is 3.69632
# After 3820 steps, per token cost is 3.68578
# After 3830 steps, per token cost is 3.85428
# After 3840 steps, per token cost is 3.75729
# After 3850 steps, per token cost is 3.74639
# After 3860 steps, per token cost is 3.59444
