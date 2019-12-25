# -*- coding: utf-8 -*-
# @Time    : 2018/3/1 11:16
# @Author  : Yajun Yin
# @Note    : copied in <Tensorflow实战Google深度学习框架>

from tensorflow.examples.tutorials.mnist import input_data
import tensorflow as tf
import pandas as pd

INPUT_NODE = 784
OUTPUT_NODE = 10

# 神经网络的参数
LAYER1_NODE = 500  # 隐藏层节点数
BATCH_SIZE = 100
LEARNING_RATE_BASE = 0.8
LEARNING_RATE_DECAY = 0.99  # 指数衰减学习率的衰减率
REGULARIZATION_RATE = 0.0001
TRAINING_STEPS = 30000
MOVING_AVERAGE_DECAY = 0.99  # 滑动平均衰减率


def inference(input_tensor, avg_class, weights1, biases1, weights2, biases2):
    if avg_class is None:
        layer1 = tf.nn.relu(tf.matmul(input_tensor, weights1) + biases1)
        return tf.matmul(layer1, weights2) + biases2
    else:
        layer1 = tf.nn.relu(tf.matmul(input_tensor, avg_class.average(weights1)) + avg_class.average(biases1))
        return tf.matmul(layer1, avg_class.average(weights2)) + avg_class.average(biases2)

def train(mnist):
    # 收集分析数据
    analysis_data = {}
    iterations = []
    validate_acc_iter = []
    test_acc_iter = []

    x = tf.placeholder(dtype=tf.float32, shape=[None, INPUT_NODE], name="x-input")
    y_ = tf.placeholder(tf.float32, [None, OUTPUT_NODE], name="y-input")

    # 隐藏层参数
    weights1 = tf.Variable(
        tf.truncated_normal([INPUT_NODE, LAYER1_NODE], stddev=0.1))
    biases1 = tf.Variable(tf.constant(0.1, shape=[LAYER1_NODE]))

    # 输出层参数
    weights2 = tf.Variable(
        tf.truncated_normal([LAYER1_NODE, OUTPUT_NODE], stddev=0.1))
    biasses2 = tf.Variable(tf.constant(0.1, shape=[OUTPUT_NODE]))

    y = inference(x, None, weights1, biases1, weights2, biasses2)

    global_step = tf.Variable(0, trainable=False)

    variable_averages = tf.train.ExponentialMovingAverage(MOVING_AVERAGE_DECAY, global_step)

    variables_average_op = variable_averages.apply(tf.trainable_variables())

    average_y = inference(x, variable_averages, weights1, biases1, weights2, biasses2)

    cross_entropy = tf.nn.sparse_softmax_cross_entropy_with_logits(labels=tf.argmax(y_, 1), logits=y)
    cross_entropy_mean = tf.reduce_mean(cross_entropy)

    regularizer = tf.contrib.layers.l2_regularizer(REGULARIZATION_RATE)
    regularization = regularizer(weights1) + regularizer(weights2)

    loss = cross_entropy_mean + regularization

    learning_rate = tf.train.exponential_decay(
        LEARNING_RATE_BASE,
        global_step,
        mnist.train.num_examples / BATCH_SIZE,
        LEARNING_RATE_DECAY
    )

    train_step = tf.train.GradientDescentOptimizer(learning_rate).minimize(loss, global_step=global_step)

    train_op = tf.group(train_step, variables_average_op)

    correct_prediction = tf.equal(tf.argmax(average_y, 1), tf.argmax(y_, 1))
    accuracy = tf.reduce_mean(tf.cast(correct_prediction, tf.float32))

    with tf.Session() as sess:
        init_op = tf.global_variables_initializer()
        sess.run(init_op)
        validate_feed = {x: mnist.validation.images, y_: mnist.validation.labels}
        test_feed = {x: mnist.test.images, y_: mnist.test.labels}

        for i in range(TRAINING_STEPS):
            if i % 1000 == 0:
                validate_acc = sess.run(accuracy, feed_dict=validate_feed)
                #
                test_acc = sess.run(accuracy, feed_dict=test_feed)
                iterations.append(i)
                validate_acc_iter.append(validate_acc)
                test_acc_iter.append(test_acc)
                print("After %d training steps, validation accuracy using average model is %g, "
                      "test accuracy using average model is %g" % (i, validate_acc, test_acc))

            xs, ys = mnist.train.next_batch(BATCH_SIZE)
            sess.run(train_op, feed_dict={x: xs, y_: ys})

        test_acc = sess.run(accuracy, feed_dict=test_feed)
        print("After %d training steps, test accuracy ""using average model is %g" % (TRAINING_STEPS, test_acc))

        analysis_data["iter"] = iterations
        analysis_data["test_acc_iter"] = test_acc_iter
        analysis_data["validate_acc_iter"] = validate_acc_iter
        analysis_pd = pd.DataFrame(analysis_data)
        analysis_pd.to_csv("D:/work/python_test/tf_a.csv")


def main(argv=None):
    mnist = input_data.read_data_sets("/path/to/MNIST_data", one_hot=True)
    train(mnist)


if __name__ == '__main__':
    tf.app.run()

