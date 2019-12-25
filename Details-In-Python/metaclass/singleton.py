#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/6/4 14:52
# @Author  : Yajun Yin
# @Note    :

# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/6/4 12:15
# @Author  : Yajun Yin
# @Note    :

import inspect
from pprint import pprint


class Singleton(object):
    def __new__(cls, *args, **kwargs):
        if not hasattr(cls, '_instance'):
            print(super(Singleton, cls))
            cls._instance = super(Singleton, cls).__new__(cls, *args, **kwargs)
        return cls._instance


class MetaB(type):
    def __new__(cls, name, bases, attrs):
        print("meta", cls, "metaclass创建类")
        cls._instance = None
        return type.__new__(cls, name, bases, attrs)

    def __init__(cls, *args, **kwargs):
        print("init", cls, "metaclass初始化")
        cls._meta_value = 4327
        super().__init__(*args, **kwargs)

    def __call__(cls, *args, **kwargs):
        print('call', cls, super(), 'args:', *args, 'kwargs:', *kwargs)
        if cls._instance is None:
            cls._instance = super().__call__(*args, **kwargs)
        return cls._instance


class bbb(object, metaclass=MetaB):
    def __init__(self, a, b=None):
        self.a = a
        self.b = b


# bbb = MetaB的实例化
# bbb的实例化，相当于调用MetaB的__call__方法

pprint(MetaB.__dict__)
pprint(bbb.__dict__)

print(bbb._instance)

#
print("###########t1##################")
t1 = bbb(5, b=66)
t2 = bbb(6, b=77)
print(t1 is t2, t1.a, t2.a, t1._meta_value)
pprint(bbb._instance)
