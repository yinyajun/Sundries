# -*- coding: utf-8 -*-
# @Time    : 2018/2/7 15:55
# @Author  : Yajun Yin
# @Note    : modified based on Programming Collective Intelligence pp. 142
from collections import defaultdict
from PIL import Image, ImageDraw

import sys

my_data = [['slashdot', 'USA', 'yes', 18, 'None'],
           ['google', 'France', 'yes', 23, 'Premium'],
           ['digg', 'USA', 'yes', 24, 'Basic'],
           ['kiwitobes', 'France', 'yes', 23, 'Basic'],
           ['google', 'UK', 'no', 21, 'Premium'],
           ['(direct)', 'New Zealand', 'no', 12, 'None'],
           ['(direct)', 'UK', 'no', 21, 'Basic'],
           ['google', 'USA', 'no', 24, 'Premium'],
           ['slashdot', 'France', 'yes', 19, 'None'],
           ['digg', 'USA', 'no', 18, 'None'],
           ['google', 'UK', 'no', 18, 'None'],
           ['kiwitobes', 'UK', 'no', 19, 'None'],
           ['digg', 'New Zealand', 'yes', 12, 'Basic'],
           ['slashdot', 'UK', 'no', 21, 'None'],
           ['google', 'UK', 'yes', 18, 'Basic'],
           ['kiwitobes', 'France', 'yes', 19, 'Basic']]


class DecisionNode(object):
    def __init__(self, col=-1, value=None, results=None, tb=None, fb=None):
        """
        Each node in a tree.
        :param col:待检验的判断条件所对应的列索引值
        :param value:判断条件的阈值
        :param results:针对当前分支的结果，它是一个字典，除了叶子节点以外，在其他节点上该值都为None
        :param tb:tb也是decisionnode，对应于结果为true时，树上相对于当前节点的子树上的节点
        :param fb:fb也是decisionnode，对 应于结果为false时，树上相对于当前节点的子树上的节点
        """
        self.col = col
        self.value = value
        self.results = results
        self.tb = tb
        self.fb = fb


def divide_set(rows, column, value):
    """
    对某一个column上的数据拆分，既能够处理数值型数据，也能够处理离散型数据。
    :param rows:所有的数据行
    :param column:数据的某一列
    :param value:判断条件的阈值
    :return:返回split_function判断为真和假的两个集合
    """
    # split_function = None
    if isinstance(value, int) or isinstance(value, float):
        split_function = lambda row: row[column] >= value
    else:
        split_function = lambda row: row[column] == value
    set1 = [row for row in rows if split_function(row)]
    set2 = [row for row in rows if not split_function(row)]
    return set1, set2


# >>> d = treepredict.DecisionNode()
# >>> d
# <treepredict.DecisionNode object at 0x00000000011DDBE0>
# >>> d.divide_set(treepredict.my_data,2,'yes')
# ([['slashdot', 'USA', 'yes', 18, 'None'], ['google', 'France', 'yes', 23, 'Premium'], ['digg', 'USA', 'yes'
# , 24, 'Basic'], ['kiwitobes', 'France', 'yes', 23, 'Basic'], ['slashdot', 'France', 'yes', 19, 'None'], ['d
# igg', 'New Zealand', 'yes', 12, 'Basic'], ['google', 'UK', 'yes', 18, 'Basic'], ['kiwitobes', 'France', 'ye
# s', 19, 'Basic']], [['google', 'UK', 'no', 21, 'Premium'], ['(direct)', 'New Zealand', 'no', 12, 'None'], [
# '(direct)', 'UK', 'no', 21, 'Basic'], ['google', 'USA', 'no', 24, 'Premium'], ['digg', 'USA', 'no', 18, 'No
# ne'], ['google', 'UK', 'no', 18, 'None'], ['kiwitobes', 'UK', 'no', 19, 'None'], ['slashdot', 'UK', 'no', 2
# 1, 'None']])
# 随便的拆分方法并不理想，观察两个set的最终订阅情况，两边都混杂了各种情况，并不能说明问题。
def unique_counts(rows):
    """
    每行的最后一列就是最终的订阅情况，将集合内各种订阅情况计数到字典中
    :param rows:
    :return:
    """
    results = defaultdict(int)
    for row in rows:
        r = row[len(row) - 1]
        results[r] += 1
    return results


def gini_impurity1(rows):
    total = len(rows)
    counts = unique_counts(rows)
    imp = 0
    for k1 in counts:
        p1 = float(counts[k1]) / total
        for k2 in counts:
            if k1 == k2:
                continue
            p2 = float(counts[k2]) / total
            imp += p1 * p2
    return imp


def gini_impurity2(rows):
    total = len(rows)
    counts = unique_counts(rows)
    tmp = 0
    for k in counts:
        p = float(counts[k]) / total
        tmp += p * p
    imp = 1 - tmp
    return imp


def gini_impurity3(rows):
    total = len(rows)
    counts = unique_counts(rows)
    imp = 0
    for k in counts:
        p = float(counts[k]) / total
        imp += p * (1 - p)
    return imp


def entropy(rows):
    from math import log
    log2 = lambda x: log(x) / log(2)
    results = unique_counts(rows)
    ent = 0.0
    for r in results:
        p = float(results[r]) / len(rows)
        ent += (-p * log2(p))
    return ent


# gini和entropy的区别：归一化后的entropy是gini的上界，说明熵的“惩罚”更大一点；entropy的变化相对于gini稍微缓慢一点。
# 为了衡量一个特征的好坏：好的特征能够使信息增益大
# 信息增益：原数据集上的熵E1，分区后的平均熵E2，Gain=E1-E2；意思是使数据集的混乱程度降低最多，意味着分区后的数据集更一致。

def build_tree(rows, split='entropy'):
    if len(rows) == 0:
        return DecisionNode()
    split_func = {'entropy': entropy, 'gini': gini_impurity2, 'var': variance}
    try:
        current_score = split_func.get(split)(rows)
    except TypeError:
        print("Input correct split type: 'entropy' or 'gini' or 'var', default:'entropy'")
        # sys.exit(1)

    # set up some variables to record best split scores
    best_gain = 0.0
    best_criteria = None
    best_sets = None

    column_count = len(rows[0]) - 1
    for col in range(column_count):
        column_values = set()
        for row in rows:
            column_values.add(row[col])
        for value in column_values:
            set1, set2 = divide_set(rows, col, value)

            # calculate the entropy of set1 and set2 and their average entropy
            # the ratio of set1
            p = len(set1) / len(rows)
            gain = current_score - p * split_func.get(split)(set1) - (1 - p) * split_func.get(split)(set2)
            # 节点不能为空：这点很重要;因为节点为空，相当于不用划分，这样best_gain还是0
            if gain > best_gain and len(set1) > 0 and len(set2) > 0:
                best_gain = gain
                best_criteria = (col, value)
                best_sets = (set1, set2)

    # create sub-tree
    # best_gain=0有两种情况：1.gain=0 2.不用划分了，就是整个set的result都一样,best_gain=0就是递归的退出条件
    if best_gain > 0:
        trueBranch = build_tree(best_sets[0])
        falseBranch = build_tree(best_sets[1])
        return DecisionNode(col=best_criteria[0], value=best_criteria[1], tb=trueBranch, fb=falseBranch)
    else:
        return DecisionNode(results=unique_counts(rows))


def print_tree(tree, indent='  '):
    # 前序遍历
    # it is a leaf
    if tree.results is not None:
        print('Leaf***', dict(tree.results))
    else:
        # print judge col and value
        print('feature' + str(tree.col) + ':' + str(tree.value) + '?')

        # print branch
        print(indent + 'T:-->', end="")
        print_tree(tree.tb, indent + '    ')
        print(indent + 'F:-->', end="")
        print_tree(tree.fb, indent + '    ')


def get_width(tree):
    if tree.results is not None:
        return 1
    else:
        return get_width(tree.tb) + get_width(tree.fb)


def get_depth(tree):
    if tree.results is not None:
        return 0
    else:
        return max(get_depth(tree.tb), get_depth(tree.fb)) + 1


def draw_tree(tree, filename="decision_tree.jpeg"):
    """
    合理的尺寸画布，将画布和根节点传递给draw_node
    :param tree:
    :param filename:
    :return:
    """
    w = get_width(tree) * 100
    h = get_depth(tree) * 100 + 120

    img = Image.new('RGB', (w, h), (255, 255, 255))
    draw = ImageDraw.Draw(img)

    draw_node(draw, tree, w / 2, 20)
    img.save(filename, 'JPEG')


def draw_node(draw, tree, x, y):
    """
    实际绘制决策树的节点，递归工作
    :param draw:
    :param tree:
    :param x:
    :param y:
    :return:
    """
    if tree.results is None:
        w1 = get_width(tree.tb) * 100
        w2 = get_width(tree.fb) * 100

        left = x - (w1 + w2) / 2
        right = x + (w1 + w2) / 2

        draw.text((x - 20, y - 10), ('feature' + str(tree.col) + ':' + str(tree.value)), fill=(0, 0, 0))
        # print('feature' + str(tree.col) + ':' + str(tree.value))

        # 分支的连线
        draw.line((x, y, left + w1 / 2, y + 100), fill=(255, 0, 0))
        draw.line((x, y, right - w2 / 2, y + 100), fill=(255, 0, 0))

        draw_node(draw, tree.tb, left + w1 / 2, y + 100)
        draw_node(draw, tree.fb, right - w2 / 2, y + 100)
    else:
        txt = u' \n'.join(['%s: %d' % v for v in tree.results.items()])
        # print(txt)
        draw.text((x - 20, y), txt, fill=(0, 0, 0))


def predict(observation, tree):
    # 每个observation有条确定路径，相当于链表
    if tree.results is not None:
        print(dict(tree.results))
        # 这种方法求dict中最大value的key太麻烦
        # return list(tree.results.keys())[list(tree.results.values()).index(max(tree.results.values()))]
        return max(tree.results.items(), key=lambda x: x[1])[0]
    else:
        v = observation[tree.col]
        # 判断这个特征是数值型还是离散型
        if isinstance(v, int) or isinstance(v, float):
            if v >= tree.value:
                branch = tree.tb
            else:
                branch = tree.fb
        else:
            if v == tree.value:
                branch = tree.tb
            else:
                branch = tree.fb
        return predict(observation, branch)


# 剪枝：目前的完全生成树只有当无法降低熵，就是best_gain==0时，才会停止创建分支。
# 一种策略是，设置一个熵减少的阈值，就是信息增益的阈值，当小于这个阈值的时候，就不再分支。
# 另一种策略是，先构建好整个完全生成树，然后再尝试消除多余节点，即剪枝。合并具有相同父节点的一组节点，合并后的熵增加
# 是否小于指定阈值。
def prune(tree, min_gain):
    """
    剪枝，当信息增益小于min_gain时，不再分支
    :param tree:
    :param min_gain: 最小阈值
    :return:
    """
    # 后序遍历
    # 判断左右分支是不是叶子节点
    # 左分支不是叶子节点，右分支不是叶子节点.
    # 这里的逻辑是（左存在 and 左不是叶子）or（右存在 and 右不是叶子）
    if tree.tb.results is None:
        prune(tree.tb, min_gain)
    elif tree.fb.results is None:
        prune(tree.fb, min_gain)
    # 左右都是叶子，看看是否需要合并叶子
    if tree.tb.results is not None and tree.fb.results is not None:
        tb, fb = [], []
        for v, c in tree.tb.results.items():
            tb += [[v]] * c
        for v, c in tree.fb.results.items():
            fb += [[v]] * c
        total = len(tb + fb)
        p = float(len(tb)) / total
        delta_gain = entropy(tb + fb) - p * entropy(tb) - (1 - p) * entropy(fb)
        if delta_gain < min_gain:
            tree.tb, tree.fb = None, None
            tree.results = unique_counts(tb + fb)


# 处理缺失值
# 缺失了某些数据，这些数据是确定分支走向所必需的，可以选择两个分支都走，加权平均。
# 如果有特征缺失，则每个分支对应的结果都会计算一遍，最终结果乘以对应的权重
def md_predict(observation, tree):
    if tree.results is not None:
        # print('leaf',dict(tree.results))
        return dict(tree.results)
    else:
        v = observation[tree.col]
        # 如果v的值缺失
        if v is None:
            # print(tree.col, tree.value)
            tr, fr = md_predict(observation, tree.tb), md_predict(observation, tree.fb)
            # print('tr', tr)
            # print('fr', fr)
            tcount = sum(tr.values())
            fcount = sum(fr.values())
            tw = float(tcount) / (tcount + fcount)
            fw = float(fcount) / (tcount + fcount)
            result = {}
            for k, v in tr.items():
                result[k] = v * tw
            for k, v in fr.items():
                if k not in result:
                    result[k] = 0
                result[k] += v * fw
            # print(result)
            return result
        else:
            if isinstance(v, int) or isinstance(v, float):
                if v >= tree.value:
                    branch = tree.tb
                else:
                    branch = tree.fb
            else:
                if v == tree.value:
                    branch = tree.tb
                else:
                    branch = tree.fb
            return md_predict(observation, branch)


def variance(rows):
    """
    高的方差说明结果分散，低的方差说明结果比较接近。对于回归，使用方差作为分裂判断的依据。
    :param rows:
    :return:
    """
    if len(rows) == 0:
        return 0
    # data中的最后一列就是结果
    data = [float(row[len(row) - 1]) for row in rows]
    mean = sum(data) / len(data)
    var = sum([(d - mean) ** 2 for d in data]) / len(data)
    return var
