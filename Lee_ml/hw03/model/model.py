#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/5/20 21:33
# @Author  : Yajun Yin
# @Note    :


# 由于github上ipython显示不稳定
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt

pd.set_option('display.max_columns', None)


# 标准化
class standardScaler(object):
    def __init__(self, train_x):
        m_ = train_x.apply(lambda x: x.mean())
        s_ = train_x.apply(lambda x: x.std())
        self.m_ = np.array(m_, float)
        self.s_ = np.array(s_, float)

    def fit(self, df):
        df = df.sub(self.m_).divide(self.s_)
        return df


class LinearRegression(object):
    def __init__(self):
        self.ss, self.train_x, self.train_y = self.train_data()
        self.test_x, self.test_y = self.test_data()

    def train_data(self):
        df = pd.read_csv("../data/train.csv")
        df = df.loc[df.observation == 'PM2.5']
        data = df.filter(regex="[^a-z]")
        data = data.drop(['Date'], axis=1)
        data = data.convert_objects(convert_numeric=True)

        train_x = []
        train_y = []
        for i in range(15):
            x = data.iloc[:, i:i + 9]
            x.columns = np.array(
                range(9))  # notice if we don't set columns name, it will have different columns name in each iteration
            #     d=data.iloc[:,-2:]
            #     d.columns=[9,10]
            #     x=x.join(d)
            y = data.iloc[:, i + 9]
            y.columns = np.array(range(1))
            train_x.append(x)
            train_y.append(y)

        train_x = pd.concat(train_x)
        train_y = pd.concat(train_y)
        # 将含有负值的记录删除
        cond1 = train_x.apply(lambda x: x > 0).all(axis=1)
        x = train_x[cond1]
        y = train_y[cond1]
        cond2 = y.apply(lambda x: x > 0)
        x = x[cond2]
        y = y[cond2]
        train_x = x
        train_y = y

        ss = standardScaler(train_x)
        train_x = ss.fit(train_x)

        train_x = np.array(train_x, float)
        train_y = np.array(train_y, float)
        # add bias to train dataset
        train_x = np.concatenate((np.ones((train_x.shape[0], 1)), train_x), axis=1)

        return ss, train_x, train_y

    def test_data(self):
        df_test = pd.read_csv('../data/test(1).csv')
        df_test = df_test[df_test['AMB_TEMP'] == 'PM2.5']
        test_data = df_test.iloc[:, 2:]
        test_data = test_data.convert_objects(convert_numeric=True)
        test_data = self.ss.fit(test_data)
        test_x = np.array(test_data, float)
        test_x = np.concatenate((np.ones((test_x.shape[0], 1)), test_x), axis=1)
        real = pd.read_csv("../data/answer.csv")
        test_y = np.array(real.value, float)
        return test_x, test_y

    def fit(self):
        # 初始化
        w = np.random.normal(0, 0.001, len(self.train_x[0]))
        lr = 1.0
        epoch_num = 20000
        # 训练, AdaGrad
        loss_history = []
        sum_grad = np.zeros(len(self.train_x[0]))
        for i in range(epoch_num):
            y_hat = np.dot(self.train_x, w)  # MxN * Nx1 = Mx1
            loss = ((self.train_y - y_hat) ** 2).mean()
            grad = np.dot(self.train_x.transpose(), y_hat - self.train_y)  # NxM * Mx1 = Nx1
            sum_grad += grad ** 2
            ada_lr = lr / np.sqrt(sum_grad)
            w -= ada_lr * grad

            loss_history.append(loss)
            if i % 1000 == 0:
                print("Step: {0}, Loss: {1}".format(i, loss))

        self.w = w
        plt.title("Loss in first 1000 step")
        plt.xlabel("step")
        plt.ylabel("MSE")
        plt.plot(loss_history[4:1000])
        plt.show()

    def predict(self):
        y_hat = np.dot(self.test_x, self.w)
        return y_hat

    @staticmethod
    def regression_metric(y, y_hat):
        sse = ((y_hat - y) ** 2).sum()
        mse = ((y_hat - y) ** 2).mean()
        rmse = np.sqrt(mse)
        mae = np.abs(y_hat - y).mean()

        ssr = ((y_hat - y.mean()) ** 2).sum()
        sst = ((y - y.mean()) ** 2).sum()
        r2 = 1 - sse / sst

        from sklearn.metrics import r2_score
        r2_score(y, y_hat)
        print("MSE: {0}, RMSE:{1}, MAE: {2}, R2: {3}".format(mse, rmse, mae, r2))


lin_regression = LinearRegression()
lin_regression.fit()
y_hat = lin_regression.predict()
lin_regression.regression_metric(lin_regression.test_y, y_hat)
