#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/16 18:23
# @Author  : Yajun Yin
# @Note    :

import os
from tensorflow.python import pywrap_tensorflow
CHECKPOINT_PATH = "D:\work\python_test\save\seq2seq_ckpt-4400"
reader = pywrap_tensorflow.NewCheckpointReader(CHECKPOINT_PATH)
var_to_shape_map = reader.get_variable_to_shape_map()
for key in var_to_shape_map:
    print("tensor_name:", key)



# tensor_name: encoder/rnn/multi_rnn_cell/cell_1/basic_lstm_cell/bias
# tensor_name: decoder/rnn/multi_rnn_cell/cell_0/basic_lstm_cell/bias
# tensor_name: mnt_model/src_emb
# tensor_name: encoder/rnn/multi_rnn_cell/cell_0/basic_lstm_cell/bias
# tensor_name: encoder/rnn/multi_rnn_cell/cell_1/basic_lstm_cell/kernel
# tensor_name: mnt_model/softmax_bias
# tensor_name: decoder/rnn/multi_rnn_cell/cell_0/basic_lstm_cell/kernel
# tensor_name: decoder/rnn/multi_rnn_cell/cell_1/basic_lstm_cell/kernel
# tensor_name: encoder/rnn/multi_rnn_cell/cell_0/basic_lstm_cell/kernel
# tensor_name: decoder/rnn/multi_rnn_cell/cell_1/basic_lstm_cell/bias
# tensor_name: mnt_model/trg_emb
