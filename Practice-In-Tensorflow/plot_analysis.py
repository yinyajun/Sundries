#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/1 15:27
# @Author  : Yajun Yin
# @Note    :
import pandas as pd
import matplotlib.pyplot as plt


def test_vs_validate_plot():
    # 读取迭代次数，每次的test_acc和validate_acc
    df = pd.read_csv("D:/work/python_test/tf_a.csv")
    # 通过pandas画图，再通过plt修改图
    df.index = df["iter"]
    df = df.iloc[:, range(2, 4)]
    df.plot()
    # 修改图的属性
    plt.xlabel("Iterations")
    plt.ylabel("Accuracy")
    plt.title("Test accuracy vs Validate accuracy under different iterations")
    # 重新设置坐标值范围
    plt.axis([0, 30000, 0.974, 0.986])
    # 获得当前坐标实例对象
    ax = plt.gca()
    # 设置x坐标的精细程度
    ax.set_xticks(range(1000, 30000, 3000))
    plt.savefig("test_vs_validate.jpg")
    plt.show()


def different_model_plot():
    # 激活函数，网络结构（层数）
    # 优化：学习率（指数衰减），正则项（系数和正则项形式），滑动平均，优化函数
    # 迭代次数：early stopping
    pass
