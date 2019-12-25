# -*- coding: utf-8 -*-
from math import log
import pandas as pd
import operator
import tree_plotter


def read_data():
    """
    创建样本数据集。
    :return:数据集，特征名列表
    """
    df = pd.read_csv("./watermelon_3a.csv", index_col=0)
    columns = list(df.columns)
    discrete_cols = columns[:-3]
    continuous_cols = columns[-3:-1]
    discrete_df = df.iloc[:, :-3]
    label_df = df.iloc[:, -1]
    df = pd.concat([discrete_df, label_df], axis=1)
    dataset = df.values
    dataset = dataset.tolist()
    return dataset, discrete_cols


def calc_shannon_ent(data_set):
    """
    计算数据集的信息熵
    :param data_set: 如： [[1, 1, 'yes'], [1, 1, 'yes'], [1, 0, 'no'], [0, 1, 'no'], [0, 1, 'no']]
    :return:
    """
    num = len(data_set)  # n rows
    # 为所有的分类类目创建字典
    label_counts = {}
    for feat_vec in data_set:
        current_label = feat_vec[-1]  # 取得最后一列数据
        if current_label not in label_counts.keys():
            label_counts[current_label] = 0
        label_counts[current_label] += 1

    # 计算香浓熵
    shannon_ent = 0.0
    for key in label_counts:
        prob = float(label_counts[key]) / num
        shannon_ent = shannon_ent - prob * log(prob, 2)
    return shannon_ent


def split_data_set(data_set, axis, value):
    """
    划分数据集子集。从原数据集中划分出axis列的特征值为value的子集，划分完后，将axis特征丢弃。
    :param data_set:  待划分的数据集
    :param axis: 特征索引
    :param value: 特征的具体值
    :return:
    """
    ret_data_set = []
    for feat_vec in data_set:
        if feat_vec[axis] == value:
            reduce_feat_vec = feat_vec[:axis]
            reduce_feat_vec.extend(feat_vec[axis + 1:])
            ret_data_set.append(reduce_feat_vec)
    return ret_data_set


def calculate_conditonal_ent(feature_idx, data_set):
    """
    计算条件熵。当数据集根据某特征划分出多个子集后，分别计算子集的信息熵，然后求期望。
    H(D|A) = \sum p(Di)*H(Di)
    :param feature_idx:
    :param data_set:
    :return:
    """
    feature_val_list = [number[feature_idx] for number in data_set]  # 得到某个特征下所有值（某列）
    unique_feature_val_list = set(feature_val_list)  # 获取无重复的属性特征值，能划分为多少个sub_data_set
    new_entropy = 0
    for feature_val in unique_feature_val_list:
        sub_data_set = split_data_set(data_set, feature_idx, feature_val)
        prob = len(sub_data_set) / float(len(data_set))  # sub_data_set的占比
        new_entropy += prob * calc_shannon_ent(sub_data_set)  # 对各sub_data_set（给定特征）香农熵(条件熵)求期望
    return new_entropy


def choose_best_feature_to_split(data_set):
    """
    按照最大信息增益划分数据
    :param data_set: 样本数据，如： [[1, 1, 'yes'], [1, 1, 'yes'], [1, 0, 'no'], [0, 1, 'no'], [0, 1, 'no']]
    :return:
    """
    num_feature = len(data_set[0]) - 1  # 特征个数，label不算特征
    base_entropy = calc_shannon_ent(data_set)  # 经验熵H(D)
    best_info_gain = 0
    best_feature_idx = -1
    for feature_idx in range(num_feature):
        new_entropy = calculate_conditonal_ent(feature_idx, data_set)
        info_gain = base_entropy - new_entropy  # 计算信息增益，g(D,A)=H(D)-H(D|A)
        # 最大信息增益
        if info_gain > best_info_gain:
            best_info_gain = info_gain
            best_feature_idx = feature_idx

    return best_feature_idx


def majority_cnt(class_list):
    """
    多数表决，返回出现次数最多的类别标签
    :param class_list: 类数组
    :return:
    """
    class_count = {}
    for vote in class_list:
        if vote not in class_count.keys():
            class_count[vote] = 0
        class_count[vote] += 1
    sorted_class_count = sorted(class_count.items(), key=operator.itemgetter(1), reversed=True)
    print(sorted_class_count[0][0])
    return sorted_class_count[0][0]


def create_tree(data_set, labels):
    """
    构建决策树
    决策树的结点结构使用的是两层嵌套的dict，
    {当前节点(特征)：{通向子节点边1(特征具体值)：子节点1,...,通向子节点边n：子节点n}}
    :param data_set: 数据集合，如： [[1, 1, 'yes'], [1, 1, 'yes'], [1, 0, 'no'], [0, 1, 'no'], [0, 1, 'no']]
    :param labels: 标签数组，如：['no surfacing', 'flippers']
    :return:
    """
    # 递归的基本情形
    class_list = [sample[-1] for sample in data_set]  # ['yes', 'yes', 'no', 'no', 'no']
    # 类别相同，停止划分
    if class_list.count(class_list[-1]) == len(class_list):
        return class_list[-1]
    # 遍历完所有特征，返回出现次数最多的类别
    if len(data_set[0]) == 1:
        return majority_cnt(class_list)

    # 递归的递推情形
    # 按照信息增益最高选取分类特征属性
    best_feature_idx = choose_best_feature_to_split(data_set)  # 返回分类的特征的数组索引
    best_feat_label = labels[best_feature_idx]  # 该特征的label
    my_tree = {best_feat_label: {}}  # 构建树的字典
    del (labels[best_feature_idx])  # 找到最佳划分特征后将其从列表中删除
    feature_values = [example[best_feature_idx] for example in data_set]
    unique_feature_values = set(feature_values)

    for feature_value in unique_feature_values:
        sub_labels = labels[:]  # 子标签数组（去除当前划分的feature_idx）
        # 构建数据的子集合，并进行递归
        sub_data_set = split_data_set(data_set, best_feature_idx, feature_value)  # 待划分的子数据集
        my_tree[best_feat_label][feature_value] = create_tree(sub_data_set, sub_labels)  # 连接子节点
    return my_tree


def classify(input_tree, feat_labels, test_vec):
    """
    决策树分类
    由于结点是两层嵌套的dict，分类时需要判断使用的是哪层的key
    :param input_tree: 决策树
    :param feat_labels: 特征标签
    :param test_vec: 测试的数据
    :return:
    """
    first_str = list(input_tree.keys())[0]  # 第一层的字典的key，代表特征
    second_dict = input_tree[first_str]  # 第二层字典，key是分支（特征具体值），value是子树
    feat_index = feat_labels.index(first_str)  # 获取特征在feat_labels中的位置
    for key in second_dict.keys():
        # test_vec中feature值指向该分支
        if test_vec[feat_index] == key:
            # 该节点仍然是dict，就是非叶子结点
            if type(second_dict[key]).__name__ == 'dict':
                class_label = classify(second_dict[key], feat_labels, test_vec)
            else:
                class_label = second_dict[key]
            return class_label


def storeTree(inputTree, filename):
    """
    持久化决策树，其实就是序列化字典
    :param inputTree:
    :param filename:
    :return:
    """
    import pickle
    with open(filename, 'wb') as f:
        pickle.dump(inputTree, f)


def loadTree(filename):
    """
    加载决策树
    :param filename:
    :return:
    """
    import pickle
    with open(filename, 'rb') as f:
        d = pickle.load(f)
    return d


if __name__ == '__main__':
    import copy

    # create decision tree
    data_set, labels = read_data()
    labels_cpy = copy.deepcopy(labels)
    decision_tree = create_tree(data_set, labels)

    # save decision tree
    storeTree(decision_tree, "watermelon_dt.a")

    # load decision tree
    dt = loadTree('watermelon_dt.a')
    print("==Decision Tree==", dt)

    # visualize decision tree
    tree_plotter.create_plot(dt)

    # predict
    prediction = classify(dt, labels_cpy, ['black', 'stiff', '', 'blur', '', 'soft_stick'])
    print(prediction)
