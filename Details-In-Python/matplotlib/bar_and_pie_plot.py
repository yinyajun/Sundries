#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/5/10 14:02
# @Author  : Yajun Yin
# @Note    :

import matplotlib.pyplot as plt
from pylab import mpl

# matplotlib没有中文字体，动态解决
mpl.rcParams['font.sans-serif'] = ['DFKai-SB']  # 指定默认字体
mpl.rcParams['axes.unicode_minus'] = False  # 解决保存图像是负号'-'显示为方块的问题

fig = plt.figure()
a = fig.add_subplot(1, 2, 1)
a.set_title("0509的*尾号")
labels = ["早于date", "晚于date"]
sizes = [114, 5985]
explode = (0.02, 0.0)
colors = ['yellowgreen', 'lightskyblue']
plt.pie(sizes, explode=explode, labels=labels, colors=colors,
        labeldistance=1.05, autopct='%3.1f%%', shadow=False,
        startangle=90, pctdistance=0.6,
        wedgeprops={"edgecolor": "k", 'linewidth': 0.5, 'linestyle': 'dashed', 'antialiased': True})

# 设置x，y轴刻度一致，这样饼图才能是圆的
plt.axis('equal')
plt.legend()

###########

b = fig.add_subplot(1, 2, 2)
b.set_title("0410的*尾号")
labels = ["早于date", "晚于date"]
sizes = [313, 6142]
explode = (0.02, 0.0)
colors = ['yellowgreen', 'lightskyblue']
plt.pie(sizes, explode=explode, labels=labels, colors=colors,
        labeldistance=1.05, autopct='%3.1f%%', shadow=False,
        startangle=90, pctdistance=0.6,
        wedgeprops={"edgecolor": "k", 'linewidth': 0.5, 'linestyle': 'dashed', 'antialiased': True})

# 设置x，y轴刻度一致，这样饼图才能是圆的
plt.axis('equal')
plt.legend()

plt.show()


def autolabel(rects):
    for rect in rects:
        height = rect.get_height()
        plt.text(rect.get_x() + rect.get_width() / 2. - 0.2, 1.01 * height, '%s' % int(height))


fig = plt.figure()
a = fig.add_subplot(1, 2, 1)
a.set_title("0509的*尾号")
labels = ["A评级", "B评级", "C评级", "D评级", "else评级"]
sizes1 = [1008, 2443, 2524, 7, 14]
sizes2 = [161, 782, 909, 4, 10]
x = list(range(len(labels)))
total_width, n = 0.8, 2
width = total_width / n
autolabel(plt.bar(x, sizes1, width=width, label='次数', fc='lightskyblue'))
for i in range(len(x)):
    x[i] += width
autolabel(plt.bar(x, sizes2, width=width, label='数目', tick_label=labels, fc='yellowgreen'))
plt.legend()

###########

a = fig.add_subplot(1, 2, 2)
a.set_title("0410的*尾号")
labels = ["A评级", "B评级", "C评级", "D评级", "else评级"]
sizes1 = [623, 2458, 932, 439, 2003]
sizes2 = [189, 927, 359, 212, 1212]
x = list(range(len(labels)))
total_width, n = 0.8, 2
width = total_width / n
autolabel(plt.bar(x, sizes1, width=width, label='次数', fc='lightskyblue'))
for i in range(len(x)):
    x[i] += width
autolabel(plt.bar(x, sizes2, width=width, label='数目', tick_label=labels, fc='yellowgreen'))
plt.legend()

plt.show()
