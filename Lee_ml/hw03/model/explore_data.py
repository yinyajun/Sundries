#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/5/23 9:40
# @Author  : Yajun Yin
# @Note    :

import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

# 数据读取与预处理
train_data = pd.read_csv("../data/train.csv")
train_data.drop(['Date', 'stations'], axis=1, inplace=True)
column = train_data['observation'].unique()
# print(column)
new_train_data = pd.DataFrame(np.zeros([24*240, 18]), columns=column)

for i in column:
    train_data1 = train_data[train_data['observation'] == i]
    # Be careful with the inplace, as it destroys any data that is dropped!
    train_data1.drop(['observation'], axis=1, inplace=True)
    train_data1 = np.array(train_data1)
    train_data1[train_data1 == 'NR'] = '0'
    train_data1 = train_data1.astype('float')
    train_data1 = train_data1.reshape(1, 5760)
    train_data1 = train_data1.T
    new_train_data[i] = train_data1

# label = np.array(new_train_data['PM2.5'][9:], dtype='float32')

# 探索性数据分析 EDA
# 最简单粗暴的方式就是根据 HeatMap 热力图分析各个指标之间的关联性
f, ax = plt.subplots(figsize=(9, 6))
sns.heatmap(new_train_data.corr(), fmt="d", linewidths=0.5, ax=ax)
plt.show()

