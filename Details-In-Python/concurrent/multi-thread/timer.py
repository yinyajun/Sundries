#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/5/30 10:43
# @Author  : Yajun Yin
# @Note    :

"""
可以和decorator中的timer对比
"""

import threading
import random
import time

TIMER = None
FAILURE_CNT = 0
ALARM_INTERVAL = 15


def check_timer():
    global TIMER, FAILURE_CNT
    msg = None
    if FAILURE_CNT:
        msg = "mission failed"
    if msg:
        print("alarm:", msg, FAILURE_CNT)
        FAILURE_CNT = 0
    TIMER = threading.Timer(ALARM_INTERVAL, check_timer)
    TIMER.start()


def task():
    time.sleep(3)
    if not random_error():
        print("task done")
    else:
        print('task failed')


def random_error():
    global FAILURE_CNT
    s = random.randint(0, 10)
    if s >= 8:
        FAILURE_CNT += 1
        return True
    return False


if __name__ == '__main__':
    TIMER = threading.Timer(1, check_timer)
    TIMER.start()
    while True:
        task()