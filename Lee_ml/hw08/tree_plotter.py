import matplotlib.pyplot as plt
from pylab import mpl

# matplotlib没有中文字体，动态解决
mpl.rcParams['font.sans-serif'] = ['DFKai-SB']  # 指定默认字体
mpl.rcParams['axes.unicode_minus'] = False  # 解决保存图像是负号'-'显示为方块的问题

# fc: facecolor, ec:edgecolor
decision_node = dict(boxstyle='round,pad=0.5', fc='yellow', ec='k', lw=1, alpha=0.4)  # annotate bbox
leaf_node = dict(fc='lightblue', alpha=0.6)  # annotate bbox
arrow_args = dict(arrowstyle='<|-', connectionstyle='arc3', color='orange')  # arrow type


# https://blog.csdn.net/u013457382/article/details/50956459
# https://blog.csdn.net/helunqu2017/article/details/78659490
# https://blog.csdn.net/qq_39131592/article/details/79056255
def plot_node(node_txt, center_pt, parent_pt, node_type):
    arrow_type = arrow_args if center_pt != parent_pt else None
    create_plot.ax1.annotate(node_txt,  # 注释文本
                             xy=parent_pt,  # 被注释的坐标点
                             xycoords='data',  # 指定数据点和注释内容的坐标系
                             textcoords='data',
                             xytext=center_pt,  # 注释文字的坐标位置
                             va="bottom",  # 注释点位置
                             ha="center",
                             bbox=node_type,  # 方框外形
                             arrowprops=arrow_type  # 箭头外形
                             )


def get_num_leafs(my_tree):
    """
    确定叶子节点个数，以便确定x轴长度
    :param my_tree:
    :return:
    """
    num_leafs = 0
    first_str = list(my_tree.keys())[0]
    second_dict = my_tree[first_str]
    for key in second_dict.keys():
        if type(second_dict[key]).__name__ == 'dict':
            num_leafs += get_num_leafs(second_dict[key])
        else:
            num_leafs += 1
    return num_leafs


def get_tree_depth(my_tree):
    """
    递归确定树的高度，以便确定y轴高度
    :param my_tree:
    :return:
    """
    max_depth = 0
    first_str = list(my_tree.keys())[0]
    second_dict = my_tree[first_str]
    for key in second_dict.keys():
        if type(second_dict[key]).__name__ == 'dict':
            thisDepth = get_tree_depth(second_dict[key]) + 1
        else:
            thisDepth = 1
        if thisDepth > max_depth:
            max_depth = thisDepth
    return max_depth


def plot_mid_text(cntr_pt, parent_pt, txt_string):
    """
    在父子结点之间添加文本信息(特征的取值)
    :param cntr_pt:
    :param parent_pt:
    :param txt_string:
    :return:
    """
    x_mid = (parent_pt[0] - cntr_pt[0]) / 2.0 + cntr_pt[0]
    y_mid = (parent_pt[1] - cntr_pt[1]) / 2.0 + cntr_pt[1]
    create_plot.ax1.text(x_mid, y_mid, txt_string,
                         size=8,
                         color="r",
                         style="italic",
                         weight="light",
                         horizontalalignment='center',
                         bbox=dict(boxstyle='round', facecolor="r", alpha=0.2))


def plot_tree(my_tree, parent_pt, node_txt):
    """
    递归绘制树
    :param my_tree:
    :param parent_pt:
    :param node_txt:
    :return:
    """
    # 计算树的宽度
    num_leafs = get_num_leafs(my_tree)

    first_str = list(my_tree.keys())[0]
    cntr_pt = update_internal_node(num_leafs)
    plot_mid_text(cntr_pt, parent_pt, node_txt)  # 标记子节点的属性（特征的值）
    plot_node(first_str, cntr_pt, parent_pt, decision_node)

    second_dict = my_tree[first_str]
    plot_tree.y_off = plot_tree.y_off - 1.0 / plot_tree.total_d

    for key in second_dict.keys():
        if type(second_dict[key]).__name__ == 'dict':
            plot_tree(second_dict[key], cntr_pt, str(key))
        else:
            pt = update_leaf_node()
            plot_mid_text(pt, cntr_pt, str(key))
            plot_node(second_dict[key], (plot_tree.x_off, plot_tree.y_off), cntr_pt, leaf_node)

    plot_tree.y_off = plot_tree.y_off + 1.0 / plot_tree.total_d  # 回溯


def update_internal_node(num_leafs):
    cntr_pt = (plot_tree.x_off + (1.0 + num_leafs) / (2.0 * plot_tree.total_w), plot_tree.y_off)
    return cntr_pt


def update_leaf_node():
    plot_tree.x_off = plot_tree.x_off + 1.0 / plot_tree.total_w
    return plot_tree.x_off, plot_tree.y_off


def create_plot(in_tree):
    """
    调用递归函数plot_tree
    :param in_tree:
    :return:
    """
    fig = plt.figure(1, facecolor='white')
    fig.clf()
    axprops = dict(xticks=[], yticks=[])  # 去除坐标
    create_plot.ax1 = plt.subplot(111, frameon=False, **axprops)  # 去除方框
    plot_tree.total_w = float(get_num_leafs(in_tree))  # 叶子节点个数
    plot_tree.total_d = float(get_tree_depth(in_tree))  # 树的深度
    plot_tree.x_off = -0.5 / plot_tree.total_w  # 已经遍历的叶子节点横坐标
    plot_tree.y_off = 1.0  # 已经遍历的叶子节点纵坐标
    root_pt = (0.5, 1.0)  # 根节点坐标
    plot_tree(in_tree, root_pt, '')
    plt.show()
