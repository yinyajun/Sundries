#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/5/16 14:02
# @Author  : Yajun Yin
# @Note    :

'''使用Linux alarm信号来使限制函数运行时间的装饰器'''

import signal
import time


def timeout(seconds):
    def wrapper(func):
        def timeout_callback(signum, frame):
            raise TimeoutError("Run %s time out, limited time is %d s." % (func, seconds))

        def inner_wrapper(*args, **kwargs):
            signal.signal(signal.SIGALRM, timeout_callback)  # unix only
            signal.alarm(seconds)  # set alarm
            r = func(*args, **kwargs)
            signal.alarm(0)  # close alarm
            return r

        return inner_wrapper

    return wrapper


@timeout(5)
def test():
    print("start")
    time.sleep(6)
    print("end")


if __name__ == '__main__':
    try:
        c = test()
    except TimeoutError:
        print("function is timeout")
        pass
