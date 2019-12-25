#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/5/13 19:19
# @Author  : Yajun Yin
# @Note    :
import numpy as np
import matplotlib.pyplot as plt
from pylab import mpl

# matplotlib没有中文字体，动态解决
mpl.rcParams['font.sans-serif'] = ['DFKai-SB']  # 指定默认字体
mpl.rcParams['axes.unicode_minus'] = False  # 解决保存图像是负号'-'显示为方块的问题

fig = plt.figure()
plt.subplot(1, 2, 1)
# generate data
x_ = np.arange(-1, 10, 0.1)
y_ = 7.5 * x_ + 3.2
x = np.arange(0, 10, 0.1)
np.random.shuffle(x)
m = 30
x = x[: m]
y = 7.5 * x + 3.2
err = np.random.normal(0, 3, m)
y = y + err

plt.plot(x_, y_, 'k')
plt.scatter(x, y, c='y')

# linear regression
# model: y_hat = w*x + b
w = 1.0
b = 1.0
alpha = 0.00001
iteration = 11000
loss_history = []

for i in range(iteration):
    m = float(len(x))
    y_hat = w * x + b
    loss = np.dot(y - y_hat, y - y_hat) / m
    loss_history.append(loss)
    grad_w = -1.0 * np.dot(y - y_hat, x) / m
    grad_b = -1.0 * np.sum(y - y_hat) / m

    # update param
    w -= alpha * grad_w
    b -= alpha * grad_b
    if i % 100 == 0:
        print("Step %i, w: %0.4f, b: %.4f, Loss: %.4f" % (i, w, b, loss))

y_hat = w * x_ + b
plt.plot(x_, y_hat, "b--")
plt.legend(('baseline', 'predict', 'data'))
plt.title("线性回归")

plt.subplot(1, 2, 2)
plt.plot(np.arange(iteration), loss_history)
plt.title("损失")

plt.show()
