#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/3/26 14:14
# @Author  : Yajun Yin
# @Note    :

import datetime
import random
import time

"""
python协程版生产者消费者模式：
生产者：1. 产生物品。 2. 将物品传给消费者消费 3. 根据消费结果，进一步处理 
消费者：1. 有效物品，进行消费，返回消费结果 2. 无效物品，直接返回消费失败
"""

CONSUMER_OK = 1
CONSUMER_ERR = 0
PRODUCE_OK = 1
PRODUCE_ERR = 0
CONSUMER = 'CONSUMER'
PRODUCER = 'PRODUCER'


def log_info(identity, info):
    now = datetime.datetime.now().strftime("%Y%m%d %H:%M:%S")
    print("[{identity}@{time}] {info}".format(identity=identity, time=now, info=info))


def coroutine(func):
    """预激生成器"""

    def wrapper(*args, **kwargs):
        gen = func(*args, **kwargs)
        next(gen)
        return gen

    return wrapper


class ProducerConsumerModel(object):
    def __init__(self):
        pass

    @staticmethod
    def do_consume(production):
        if production == PRODUCE_OK:
            log_info(CONSUMER, "consume successful product")
            time.sleep(3)
            consume_result = random.randint(0, 1)
            if consume_result == 0:
                log_info(CONSUMER, "consume successfully")
                return CONSUMER_OK
        else:
            log_info(CONSUMER, "consume failed product")
        log_info(CONSUMER, "consume failed")
        return CONSUMER_ERR

    @staticmethod
    def do_produce():
        print("")
        log_info(PRODUCER, "do_produce")
        time.sleep(1)
        product = random.randint(0, 1)
        if product == PRODUCE_OK:
            log_info(PRODUCER, "produce successfully")
        else:
            log_info(PRODUCER, "produce failed")
        return product

    @staticmethod
    def treat_consumer_result(result):
        if result == CONSUMER_OK:
            log_info(PRODUCER, "treat successful consume result")
        else:
            log_info(PRODUCER, "treat failed consume result")
        time.sleep(2)

    @coroutine
    def consumer(self):
        result = None
        while True:
            production = yield result
            result = self.do_consume(production)

    def producer(self, consumer):
        while True:
            product = self.do_produce()
            result = consumer.send(product)
            self.treat_consumer_result(result)


if __name__ == '__main__':
    a = ProducerConsumerModel()
    c = a.consumer()
    a.producer(c)
