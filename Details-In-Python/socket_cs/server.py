#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/8/1 19:19
# @Author  : Yajun Yin
# @Note    :

import socket
import threading
import time

# from threading import Thread

service_ip = '10.18.242.215'
service_port = 9999


s = socket.socket()
# 监听端口:
s.bind((service_ip, service_port))
s.listen(5)
print('Waiting for connection...')


def tcplink(sock, addr):
    print('Accept new connection from %s:%s...' % addr)
    sock.send(b'Welcome!')
    while True:
        data = sock.recv(1024)
        time.sleep(1)
        if not data or data.decode('utf-8') == 'exit':
            break
        sock.send(('Hello, %s!' % data.decode('utf-8')).encode('utf-8'))
    sock.close()
    print('Connection from %s:%s closed.' % addr)


while True:
    # 接受一个新连接:
    sock, addr = s.accept()
    # 创建新线程来处理TCP连接:
    t = threading.Thread(target=tcplink, args=(sock, addr))
    t.start()

