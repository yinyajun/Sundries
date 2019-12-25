#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/5/16 11:31
# @Author  : Yajun Yin
# @Note    :
"""将函数动态添加cron功能，能够完成：
1. 每天8点跑
2. 每隔2小时跑
"""

import datetime


def cron(interval, rem):
    def wrapper(func):
        def inner_wrapper(*args, **kwargs):
            cur = datetime.datetime.now().strftime("%H:%M")
            cur_hour, cur_min = cur.split(":")

            if int(cur_hour) % interval == rem:
                return func(*args, **kwargs)
            else:
                return

        return inner_wrapper

    return wrapper


# 每天12点跑
@cron(24, 12)
def task1():
    pass


# 每隔2小时跑
@cron(2, 0)
def task2():
    pass
