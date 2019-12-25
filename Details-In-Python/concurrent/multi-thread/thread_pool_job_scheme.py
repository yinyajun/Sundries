#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/3/24 14:21
# @Author  : Yajun Yin
# @Note    :

import time
import threading
from queue import Queue
import subprocess
import random

RETRY_NUM = 1
READY = "Ready"
RUN = "Run"
RETRY = "Retry"
RESULT = "Result"

"""
任务执行框架
适用场景：大量任务需要跑，但每次只能跑有限个。如：用spark生产数据集的时候，有时需要跑30天的数据，但每次只能起3个，以免造成集群拥堵
原理：类似线程池。建立一个队列传递消息，起固定个数的线程。每个线程不停的从队列中拿消息来执行对应任务，直到队列里所有消息对应的任务完成，退出主线程
例子：见main
使用方法：
1. 建立工作线程：建立XXXJobThread类，重写default_cmd和default_job方法
2. 实例化Jobs类，如Jobs(2, SparkJobThread)，2代表几个线程，SparkJobThread是工作线程类
3. 使用submit方法提交。Jobs(2, SparkJobThread).submit(msgs)。msgs用来进入队列传递消息。
"""


class JobThread(threading.Thread):
    retry_time = RETRY_NUM

    def __init__(self, queue):
        threading.Thread.__init__(self)  # 父类必须第一个init
        self.queue = queue

    def run(self):
        while True:
            msg = self.queue.get()  # 阻塞程序，时刻监听队列是否有消息
            self.run_job(msg)
            self.queue.task_done()  # 标志队列中的任务完成，配合queue.join

    def pre_job(self, msg):
        return msg

    def do_job(self, msg, retry_time=0):
        cmd = self.default_cmd(msg)
        if not retry_time:
            self.log_info(RUN, cmd)
        else:
            self.log_info(RETRY + "#" + str(retry_time), cmd)
        status = self.default_job(msg)
        if not isinstance(status, int):
            raise ValueError("status should be int")

        return status

    def default_cmd(self, msg):
        cmd = msg
        return cmd

    def default_job(self, cmd):
        status = 0
        return status

    def post_job(self, msg, status):
        pass

    def run_job(self, msg):
        # pre_job
        msg = self.pre_job(msg)
        self.log_info(READY, "Job@{}".format(msg))
        # do_job
        status = self.do_job(msg)
        # retry
        if status != 0:
            retry_time = 1
            while status != 0 and retry_time <= self.retry_time:
                status = self.do_job(msg, retry_time=retry_time)
                retry_time += 1
        # post_job
        self.post_job(msg, status)
        self.log_info(RESULT, "Job@{}-{}".format(msg, "Success" if status == 0 else "Failed"))

    @staticmethod
    def log_info(status, info):
        print("[Thread-{}][{}]{}".format(threading.get_ident(), status, info))


class Jobs(object):
    def __init__(self, work_num, thread_cls):
        self.queue = Queue()
        self.work_num = work_num
        self.thread_cls = thread_cls
        assert (isinstance(self.thread_cls(self.queue), JobThread))

    def set_thread_pool(self, worker_nums, queue):
        for _ in range(worker_nums):
            t = self.thread_cls(queue)
            t.setDaemon(True)  # 设为守护线程，当主线程退出时，直接结束守护线程（因为JobThread中有死循环，必须设为守护线程）
            t.start()

    def submit(self, tasks):
        self.set_thread_pool(self.work_num, self.queue)
        for task in tasks:
            self.queue.put(task)
        self.queue.join()  # 只有队列中的所有任务完成才结束主线程


class SparkJobThread(JobThread):
    def __init__(self, queue):
        super().__init__(queue)

    def default_cmd(self, msg):
        return "spark-submit --executor-core=6 --executor-memory=12G taks.py --date={}".format(msg)

    def default_job(self, msg):
        cmd = self.default_cmd(msg)
        print(cmd)
        time.sleep(2)
        return random.randint(0, 1)


class PingJobThread(JobThread):
    def __init__(self, queue):
        super().__init__(queue)

    def default_cmd(self, msg):
        return "ping {}".format(msg)

    def default_job(self, msg):
        cmd = self.default_cmd(msg)
        return subprocess.call(cmd, shell=True)


if __name__ == "__main__":
    Jobs(2, SparkJobThread).submit(['20190321', '20190322', '20190323', '20190324'])
    # Jobs(2, PingJobThread).submit(['www.baidu.com', '127.1.1.1', 'www.sina.com.cn'])
