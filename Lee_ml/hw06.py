#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/6/02 19:27
# @Author  : Yajun Yin
# @Note    :


import numpy as np
import matplotlib.pyplot as plt
from pylab import mpl

# matplotlib没有中文字体，动态解决
mpl.rcParams['font.sans-serif'] = ['DFKai-SB']  # 指定默认字体
mpl.rcParams['axes.unicode_minus'] = False  # 解决保存图像是负号'-'显示为方块的问题

POSITIVE_SIZE = 100
NEGATIVE_SIZE = 100
THRESH = 0.5


def generate_data():
    # Create red points centered at (-1, -1)
    red_points = np.random.randn(POSITIVE_SIZE, 2) - 1 * np.ones((POSITIVE_SIZE, 2))

    # Create blue points centered at (1, 1)
    blue_points = np.random.randn(NEGATIVE_SIZE, 2) + 1 * np.ones((NEGATIVE_SIZE, 2))
    X = np.concatenate((blue_points, red_points))
    Y = np.concatenate([np.ones(POSITIVE_SIZE), np.zeros(NEGATIVE_SIZE)], axis=0)
    Y_ = np.concatenate([np.array([[1, 0]] * POSITIVE_SIZE), np.array([[0, 1]] * NEGATIVE_SIZE)], axis=0)

    random_index = np.random.permutation(POSITIVE_SIZE + NEGATIVE_SIZE)
    X = X[random_index]
    Y = Y[random_index]
    Y_ = Y_[random_index]
    return X, Y, Y_, red_points, blue_points


def dataset_split(dataset, train_ratio, test_ratio=None):
    size = len(dataset)
    # train, test, val
    if test_ratio:
        return np.split(dataset, [int(train_ratio * size), int((test_ratio + train_ratio) * size)])
    # train, test
    return np.split(dataset, [int(train_ratio * size)])


# Stand Scaler
class StandScaler(object):
    def __init__(self, X):
        self.m_ = np.mean(X, axis=0)
        self.s_ = np.std(X, axis=0)

    def fit(self, data):
        assert data.shape[1] == self.m_.shape[0]
        return (data - self.m_) / self.s_


# Logistic Regression
class LogisticRegression(object):
    def __init__(self, lr, epoch):
        self.lr = lr
        self.epoch = epoch

        self.X = None
        self.Y = None
        self.W = None
        self.Y_hat = None

    @staticmethod
    def _sigmoid(a):
        return 1.0 / (1 + np.exp(-a))

    def logit(self, X, w):
        # mxd * dx1 = mx1
        z = np.dot(X, w)
        return self._sigmoid(z)

    def predict(self, X=None):
        if X is None:
            X = self.X
        Y_hat = self.logit(X, self.W)
        return Y_hat

    def loss(self, Y=None, Y_hat=None):
        if Y_hat is None or Y is None:
            Y_hat = self.predict()
            Y = self.Y
        return np.mean(-1.0 * (Y * np.log(Y_hat) + (1.0 - Y) * np.log(1.0 - Y_hat)))

    def compute_grad(self):
        self.Y_hat = self.predict()
        # dxm * mx1 = dx1
        grad = np.dot(self.X.transpose(), self.Y_hat - self.Y)
        return grad

    def fit(self, X, Y):
        self.X = np.concatenate((np.ones((X.shape[0], 1)), X), axis=1)
        self.W = np.random.normal(0, 0.001, len(self.X[0]))
        self.Y = Y
        loss_history = []
        sum_grad = np.zeros(len(self.X[0]))
        for i in range(self.epoch):
            loss = self.loss()
            grad = self.compute_grad()
            sum_grad += grad ** 2
            ada_lr = self.lr / np.sqrt(sum_grad)
            self.W -= ada_lr * grad
            loss_history.append(loss)

            if i % 100 == 0:
                print("Step: {0}, Loss: {1}".format(i, loss))

    def metric(self, Y_hat=None, Y=None):
        if Y_hat is None or Y is None:
            Y = self.Y
            Y_hat = self.predict()
        # acc
        prediction = Y_hat > THRESH
        label = Y > THRESH
        correct_prediction = np.equal(prediction, label)
        acc = np.mean(correct_prediction.astype(float))
        # True Positive
        tp = np.sum(np.logical_and(prediction == True, label == True).astype(int))
        # False Negative
        fn = np.sum(np.logical_and(prediction == False, label == True).astype(int))
        # False Positive
        fp = np.sum(np.logical_and(prediction == True, label == False).astype(int))
        # True Negative
        tn = np.sum(np.logical_and(prediction == False, label == False).astype(int))
        accuracy = (tp + tn) / (tp + tn + fn + fp)
        precision = tp / (tp + fp)
        recall = tp / (tp + fn)
        f_score = 2 * precision * recall / (precision + recall)
        return acc, accuracy, precision, recall, f_score


if __name__ == '__main__':
    # train
    X, Y, _, red, blue = generate_data()
    X_train, X_test = dataset_split(X, 0.7)
    Y_train, Y_test = dataset_split(Y, 0.7)
    lr = LogisticRegression(0.1, 4000)
    lr.fit(X_train, Y_train)
    _, accuracy, precision, recall, f_score = lr.metric()
    print("[Train] accuracy:{acc:0.4f}, precision:{prec:0.4f},recall:{rec:0.4f}, f_score:{f:0.4f}".format(
        acc=accuracy, prec=precision, rec=recall, f=f_score))

    # test
    X_test = np.concatenate((np.ones((X_test.shape[0], 1)), X_test), axis=1)
    Y_test_pred = lr.predict(X_test)
    test_loss = lr.loss(Y_test, Y_test_pred)
    print(test_loss)
    _, accuracy, precision, recall, f_score = lr.metric(Y_test_pred, Y_test)
    print("[Test] accuracy:{acc:0.4f}, precision:{prec:0.4f},recall:{rec:0.4f}, f_score:{f:0.4f}".format(
        acc=accuracy, prec=precision, rec=recall, f=f_score))

    plt.scatter(red[:, 0], red[:, 1], color='red')
    plt.scatter(blue[:, 0], blue[:, 1], color='blue')
    W = lr.W
    print(W)
    x = np.linspace(-4, 4, 100)
    y = (-W[0] - W[1] * x) / W[2]
    plt.plot(x, y)
    plt.title("逻辑回归")
    plt.show()
