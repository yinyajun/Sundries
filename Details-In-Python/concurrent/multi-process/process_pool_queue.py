#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/5/30 16:44
# @Author  : Yajun Yin
# @Note    :

import os
import logging
from multiprocessing import Process, Queue, JoinableQueue

"""
task_queue: 类似生产者
res_queue: 类似消费者
"""

process_num = 3


def init_logger(name):
    logger = logging.getLogger(name)
    logger.setLevel(level=logging.INFO)
    handler = logging.FileHandler("{name}.log".format(name=name))
    handler.setLevel(logging.INFO)
    formatter = logging.Formatter(fmt='%(levelname)s %(asctime)s %(message)s', datefmt='%Y-%m-%d %H:%M:%S')
    handler.setFormatter(formatter)
    logger.addHandler(handler)
    return logger


def get_tasks():
    return [1, 2, 3]


class JobProcess(Process):
    def __init__(self, task_queue, res_queue):
        Process.__init__(self)  # 父类必须第一个init
        self.task_queue = task_queue
        self.res_queue = res_queue
        self.logger = None

    def run(self):
        self.logger = init_logger(str(os.getpid()))  # 每个子进程建立日志
        while True:
            task = self.task_queue.get()  # 阻塞程序，时刻监听队列是否有消息
            ret = self.do(task)
            self.res_queue.put(ret)
            self.task_queue.task_done()  # 标志队列中的任务完成，配合queue.join

    def do(self, task):
        # should be over-write
        return 1


task_queue = JoinableQueue()  # 实现了join和task_done的Queue
res_queue = Queue()

# set process pool
for _ in range(process_num):
    p = JobProcess(task_queue, res_queue)
    p.daemon = True
    p.start()

    all_tasks = get_tasks()
    for task in all_tasks:
        task_queue.put(task)

    task_queue.join()  # 只有队列中的所有任务完成才结束,task_queue.qsize==0
