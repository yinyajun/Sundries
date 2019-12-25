import matplotlib.pyplot as plt
import numpy as np
from pylab import mpl

# matplotlib没有中文字体，动态解决
mpl.rcParams['font.sans-serif'] = ['DFKai-SB']  # 指定默认字体
mpl.rcParams['axes.unicode_minus'] = False  # 解决保存图像是负号'-'显示为方块的问题

fig = plt.figure()
x = np.arange(-10, 10, 0.1)
plt.subplot(121)
plt.plot(x, x ** 2)
plt.ylim(0, 15)
plt.title("Hessian最大特征值对应的特征向量方向[1,1]")
plt.xticks([])  # remove x alix scale
plt.annotate("曲率大，导数增长快", (2, 2))
#
plt.subplot(122)
plt.plot(x, x ** 2 / 5)
plt.ylim(0, 15)
plt.title("Hessian最小特征值对应特征向量方向[1,-1]")
plt.xticks([])  # remove x alix scale
plt.annotate("曲率小，导数增长慢", (4, 2))
#
plt.show()