#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/7 14:51
# @Author  : Yajun Yin
# @Note    :

import numpy as np
import tensorflow as tf

"""LSTM with ease"""
lstm = tf.nn.rnn_cell.BasicLSTMCell(lstm_hidden_size)

# state.c and state.h
state = lstm.zero_state(batch_size, tf.float32)

loss = 0.0

# num_steps is length of series
for i in range(num_steps):
    if i > 0:
        # reuse defined variables
        tf.get_variable_scope().reuse_variables()

        lstm_output, state = lstm(current_input, state)

        final_output = fully_connected(lstm_output)

        loss += calc_loss(final_output, expected_output)
