#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/6/14 12:33
# @Author  : Yajun Yin
# @Note    :

# Gini Coefficient versus AOC
# Conclusion: Gini = 2*AOC -1
# https://stats.stackexchange.com/questions/155310/what-is-the-difference-between-gini-and-auc-curve-interpretation
import numpy as np
import matplotlib.pyplot as plt

predictions = np.asarray([0.9, 0.3, 0.8, 0.75, 0.65, 0.6, 0.78, 0.7, 0.05, 0.4, 0.4, 0.05, 0.5, 0.1, 0.1])
actual = np.asarray([1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0])

x = np.linspace(0, 1, len(actual))

# get gini line
print("Original Actual Label:", actual)
idx = np.argsort(predictions)
y = actual[idx]
print("After sort according to predictions, Actual Label:", y)
y = np.cumsum(y)  # calculate cumulative values
y = y / y[-1]  # normalization

# get 45 degree line
y2 = x

# plot
ax = plt.subplot(111)
plt.fill_between(x, y)
plt.fill_between(x, y, y2)

#
plt.title("Gini Coeffients")
plt.xlabel("Cumulative share of Predictions")
plt.ylabel("Cumulative share of Actual Values")
# ax.axis('equal')  # it will disable xlim and ylim
ax.set_aspect('equal', adjustable='box')  # square-like figure

# annotate
plt.annotate("gini", xy=(0.5, 0.4), xytext=(0.5, 0.6),
             arrowprops=dict(arrowstyle='-|>', connectionstyle='arc3', color='black'),
             bbox=dict(boxstyle='round,pad=0.5', fc='yellow', ec='k', lw=1, alpha=0.4)
             )

plt.show()
