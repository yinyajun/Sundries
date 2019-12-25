#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/5/21 15:45
# @Author  : Yajun Yin
# @Note    :


import pymysql

MYSQL_CONF = {
    'host': '0.0.0.0',
    'user': 'me',
    'password': '321',
    'db': 'user_score',
    'port': 566}

# insert sql
user_f = [('0', '100'), ('1', '99')]
sql = "insert into table (user, score) values(%s, %s) on duplicate key update feature= %s"
params = [(us[0], us[1], us[1]) for us in user_f]

db = pymysql.connect(**MYSQL_CONF)
try:
    cursor = db.cursor()
    cursor.executemany(sql, params)
    db.commit()
    cursor.close()
finally:
    db.close()

# query sql
sql = "select * from table"
conn = pymysql.connect(**MYSQL_CONF)
retry_times = 0
while retry_times < 3:
    try:
        with conn.cursor() as cursor:
            cursor.execute(sql)
            ret = cursor.fetchall()
        conn.commit()
    except pymysql.OperationalError:
        pass
    finally:
        conn.close()
