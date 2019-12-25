#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/3/15 14:21
# @Author  : Yajun Yin
# @Note    :


"""metaclass编程，动态定义类。廖雪峰教程中使用ORM作为例子，熟悉这个例子还可以理清实例方法和类方法的区别。"""


class Field(object):
    """
    存储table的字段名字和类型
    """

    def __init__(self, name, column_type):
        self.name = name
        self.column_type = column_type


class StringField(Field):
    def __init__(self, name):
        super().__init__(name, "varchar(100)")


class IntegerField(Field):
    def __init__(self, name):
        super().__init__(name, "bigint")


class ModelMetaClass(type):
    def __new__(mcs, name, bases, attrs):
        # 对于Model类，不做修改，直接创建。

        if name == "Model":
            return type.__new__(mcs, name, bases, attrs)
        print(attrs.items())
        print("Found model: %s" % name)
        column_maps = {}

        # 这里attrs是使用该元类修饰的类本身的类属性
        for k, v in attrs.items():
            if isinstance(v, Field):  # 利用多态判断
                print("Found %s ==> %s" % (k, v))
                column_maps[k] = v
        for k in column_maps.keys():
            attrs.pop(k)

        # 为类添加属性
        attrs['__table__'] = name  # 类名作为表名
        attrs['__col_maps__'] = column_maps
        # 通过修改attrs来创建动态的类
        return type.__new__(mcs, name, bases, attrs)


class Model(dict, metaclass=ModelMetaClass):
    """
    Model类继承 字典，实例化的时候直接使用dict的实例化方式
    """

    def __init__(self, **kwargs):
        super().__init__(**kwargs)

    def __getattr__(self, item):
        try:
            return self[item]  # 实例属性
        except KeyError:
            raise AttributeError("'Model' object has no attribute '%s'" % item)

    def save(self):

        cols = []
        vals = []
        args = []
        for k, v in self.__col_maps__.items():  # 类属性
            cols.append(v.name)
            vals.append('?')
            args.append(getattr(self, k, None))  # 获取实例属性(要求类属性col_map中的key和实例属性同名)
        cols = ','.join(cols)
        vals = ','.join(vals)
        sql = "insert %s (%s) values (%s)" % (self.__table__, cols, vals)
        print("SQL: %s" % sql)
        print("Args: %s" % str(args))


class UserTable(Model):
    id = IntegerField("userid")
    name = StringField("username")
    pwd = IntegerField("password")


u = UserTable(id=888, name='Michael', pwd=123456)
u.save()

