#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/8/1 19:19
# @Author  : Yajun Yin
# @Note    :

import socket

service_ip = '10.18.242.215'
service_port = 9999

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
# 建立连接:
s.connect((service_ip, service_port))
# 接收欢迎消息:
print(s.recv(1024).decode('utf-8'))
for data in [b'Michael', b'Tracy', b'Sarah']:
    # 发送数据:
    s.send(data)
    print(s.recv(1024).decode('utf-8'))
s.send(b'exit')
s.close()



