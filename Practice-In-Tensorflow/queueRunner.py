"""http://wiki.jikexueyuan.com/project/tensorflow-zh/how_tos/threading_and_queues.html"""

#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/3/5 11:56
# @Author  : Yajun Yin
# @Note    :


import tensorflow as tf

example = tf.random_normal([2])

# Create a queue, and an op that enqueues examples one at a time in the queue.
"""If the `shapes` argument is specified, each component of a queue
    element must have the respective fixed shape. If it is
    unspecified, different queue elements may have different shapes,
    but the use of `dequeue_many` is disallowed."""
queue = tf.RandomShuffleQueue(10, 5, "float", shapes=[(2,)])
enqueue_op = queue.enqueue(example)

# Create a training graph that starts by dequeuing a batch of examples.
inputs = queue.dequeue_many(3)
train_op = inputs
# train_op = ...use 'inputs' to build the training part of the graph...

# Create a queue runner that will run 4 threads in parallel to enqueue
# examples.
qr = tf.train.QueueRunner(queue, [enqueue_op] * 4)

# Launch the graph.
sess = tf.Session()
# Create a coordinator, launch the queue runner threads.
coord = tf.train.Coordinator()
enqueue_threads = qr.create_threads(sess, coord=coord, start=True)  # 启动入队线程
# Run the training loop, controlling termination with the coordinator.
# Queue runner support catch exception.
# noinspection PyBroadException
try:
    for step in range(3):
        if coord.should_stop():
            break
        print(sess.run(train_op))
except Exception as e:
    coord.request_stop(e)

# Terminate as usual.  It is innocuous to request stop twice.
# When done, ask the threads to stop.
coord.request_stop()
# And wait for them to actually do it.
coord.join(enqueue_threads)

