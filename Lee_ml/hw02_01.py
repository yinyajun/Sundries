#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/5/15 19:27
# @Author  : Yajun Yin
# @Note    :

import numpy as np
import matplotlib.pyplot as plt
from pylab import mpl

# matplotlib没有中文字体，动态解决
mpl.rcParams['font.sans-serif'] = ['DFKai-SB']  # 指定默认字体
mpl.rcParams['axes.unicode_minus'] = False  # 解决保存图像是负号'-'显示为方块的问题

# generate data
x_ = np.arange(0, 80, 0.1)
y_ = 1.5 * x_ + 2.2
x = np.arange(0, 80, 0.1)
np.random.shuffle(x)
m = 100
x_data = x[: m]
err = np.random.normal(0, 0.3, m)
y_data = 1.5 * x_data + 2.2 + err

x = np.arange(0, 5, 0.1)
y = np.arange(0, 5, 0.1)
Z = np.zeros((len(x), len(y)))
X, Y = np.meshgrid(x, y)
for i in range(len(x)):
    for j in range(len(y)):
        b = x[i]
        w = y[j]
        Z[j][i] = 0  # meshgrid吐出结果：y为行，x为列
        for n in range(len(x_data)):
            Z[j][i] += (y_data[n] - b - w * x_data[n]) ** 2
        Z[j][i] /= len(x_data)


# linear regression
def gradient_descent(x_data, y_data, w, b, lr):
    m = len(x_data)
    y_hat = w * x_data + b
    loss = np.dot(y_data - y_hat, y_data - y_hat) / m
    grad_b = -2.0 * np.sum(y_data - y_hat) / m
    grad_w = -2.0 * np.dot(y_data - y_hat, x_data) / m

    # update param
    b -= lr * grad_b
    w -= lr * grad_w
    return b, w, loss


def batch_gd():
    b = 4
    w = 4
    lr = 0.0001
    iteration = 70000
    b_history = [b]
    w_history = [w]
    loss_history = []
    # batch gradient descent
    for i in range(iteration):
        b, w, loss = gradient_descent(x_data, y_data, w, b, lr)
        b_history.append(b)
        w_history.append(w)
        loss_history.append(loss)
        if i % 10000 == 0:
            print("Step %d, w: %0.4f, b: %.4f, Loss: %.4f" % (i, w, b, loss))
    return b_history, w_history, loss_history


def mini_batch_gd(batch_size):
    b = 4
    w = 4
    lr = 0.0001
    iteration = 7000
    b_history = [b]
    w_history = [w]
    loss_history = []
    # batch gradient descent
    for i in range(iteration):
        # np.random.shuffle(x_data)
        assert batch_size < len(x_data)
        batches = len(x_data) // batch_size
        for j in range(batches):
            if j == batches - 1:
                x_d = x_data[j * batch_size:]
                y_d = y_data[j * batch_size:]
            else:
                x_d = x_data[j * batch_size:(j + 1) * batch_size]
                y_d = y_data[j * batch_size:(j + 1) * batch_size]
            b, w, loss = gradient_descent(x_d, y_d, w, b, lr)
            b_history.append(b)
            w_history.append(w)
            loss_history.append(loss)
        if i % 1000 == 0:
            print("Step %d, w: %0.4f, b: %.4f, Loss: %.4f" % (i, w, b, loss))
    return b_history, w_history, loss_history


def sgd():
    b = 4
    w = 4
    lr = 0.0001
    iteration = 700
    batch_size = 1
    b_history = [b]
    w_history = [w]
    loss_history = []
    # batch gradient descent
    for i in range(iteration):
        # np.random.shuffle(x_data)
        assert batch_size < len(x_data)
        batches = len(x_data) // batch_size
        for j in range(batches):
            if j == batches - 1:
                x_d = x_data[j * batch_size:]
                y_d = y_data[j * batch_size:]
            else:
                x_d = x_data[j * batch_size:(j + 1) * batch_size]
                y_d = y_data[j * batch_size:(j + 1) * batch_size]
            b, w, loss = gradient_descent(x_d, y_d, w, b, lr)
            b_history.append(b)
            w_history.append(w)
            loss_history.append(loss)
        if i % 100 == 0:
            print("Step %d, w: %0.4f, b: %.4f, Loss: %.4f" % (i, w, b, loss))
    return b_history, w_history, loss_history


def plot(b_history, w_history, model_name, color):
    C = plt.contourf(x, y, Z, 50, alpha=0.5, cmap=plt.get_cmap('jet'))  # 填充等高线
    # # plt.clabel(C, inline=True, fontsize=5)
    plt.plot([2.2], [1.5], 'x', ms=12, mew=3, color="orange")
    plt.plot(b_history, w_history, 'o-', ms=3, lw=0.5, color=color)
    plt.xlabel(r'$b$')
    plt.ylabel(r'$w$')
    plt.title(model_name)


#
b_history, w_history, loss_history = batch_gd()
b_history1, w_history1, loss_history1 = mini_batch_gd(10)
b_history2, w_history2, loss_history2 = sgd()

#
plt.subplot(1, 3, 1)
plot(b_history, w_history, "批量梯度下降, 70000epoch", 'black')
plt.subplot(1, 3, 2)
plot(b_history1, w_history1, "小批量梯度下降，7000epoch", 'red')
plt.subplot(1, 3, 3)
plot(b_history, w_history, "随机梯度下降，700epoch", "blue")

plt.show()
