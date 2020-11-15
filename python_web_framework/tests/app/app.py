#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/12/23 23:39
# @Author  : Yajun Yin
# @Note    :

import time
from web.util import *
import web.web as web


def datetime_filter(t):
    delta = int(time.time() - t)
    if delta < 60:
        return u'1分钟前'
    if delta < 3600:
        return u'%s分钟前' % (delta // 60)
    if delta < 86400:
        return u'%s小时前' % (delta // 3600)
    if delta < 604800:
        return u'%s天前' % (delta // 86400)
    dt = datetime.datetime.fromtimestamp(t)
    return u'%s年%s月%s日' % (dt.year, dt.month, dt.day)


# 创建一个WSGIApplication
app = web.WSGIApplication(dir_path(__file__))
# 初始化模板引擎
template_engine = web.Jinja2TemplateEngine(os.path.join(dir_path(__file__), 'templates'))
template_engine.add_filter('datetime', datetime_filter)
app.template_engine = template_engine


@web.get('/welcome:user')
def app_dynamic_route_test(user):
    return to_byte('<h1>Hello, %s</h1>' % user)


class Users:
    def __init__(self, name, email):
        self.name = name
        self.email = email


class Blog:
    def __init__(self, id, user_id, user_name, name, summary, content, created_at):
        self.id = id
        self.user_id = user_id
        self.user_name = user_name
        self.name = name
        self.summary = summary
        self.content = content
        self.created_at = created_at


@web.view('blogs.html')
@web.get('/')
def app_view_test():
    blogs = [Blog(3, 10, 'yyj', 'article', 'summary', 'cotent', 1577167235)] * 5
    user = Users('nike', 'fsafe')
    return dict(users=user, blogs=blogs)


if __name__ == '__main__':
    app.add_url(app_dynamic_route_test)
    app.add_url(app_view_test)
    # 在9000端口上启动本地测试服务器
    app.tornado_run(9000)
