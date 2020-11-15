#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/12/23 18:44
# @Author  : Yajun Yin
# @Note    :

"""
A simple, lightweight, WSGI-compatible web framework.
"""

import functools
import mimetypes
import sys
import threading
import traceback
import types
from io import BytesIO

import web.error as error
from web.request import Request
from web.util import *

from web.response import Response

logger = Logger("web.py", logging.DEBUG)

ctx = threading.local()  # thread local object for storing request and response:


def get(path):
    """
    A @get decorator.
    @get('/:id')
    def index(id):
        pass
    >>> @get('/test/:id')
    ... def test():
    ...     return 'ok'
    """

    def _decorator(func):
        func.__web_route__ = path
        func.__web_method__ = 'GET'
        return func

    return _decorator


def post(path):
    """
    A @post decorator.
    >>> @post('/post/:id')
    ... def testpost():
    ...     return '200'
    """

    def _decorator(func):
        func.__web_route__ = path
        func.__web_method__ = 'POST'
        return func

    return _decorator


_re_route = re.compile(r'(\:[a-zA-Z_]\w*)')


def _build_regex(path):
    """
    Convert route path to regex.
    """
    re_list = ['^']
    var_list = []
    is_var = False
    for v in _re_route.split(path):
        if is_var:
            var_name = v[1:]
            var_list.append(var_name)
            re_list.append(r'(?P<%s>[^\/]+)' % var_name)
        else:
            s = ''
            for ch in v:
                if '0' <= ch <= '9':
                    s = s + ch
                elif 'A' <= ch <= 'Z':
                    s = s + ch
                elif 'a' <= ch <= 'z':
                    s = s + ch
                else:
                    s = s + '\\' + ch
            re_list.append(s)
        is_var = not is_var
    re_list.append('$')
    return ''.join(re_list)


class Route(object):
    """
    A Route object is a callable object.
    """

    def __init__(self, func):
        self.path = func.__web_route__
        self.method = func.__web_method__
        self.is_static = _re_route.search(self.path) is None
        if not self.is_static:
            self.route = re.compile(_build_regex(self.path))
        self.func = func

    def match(self, url):
        m = self.route.match(url)
        if m:
            return m.groups()
        return None

    def __call__(self, *args):
        return self.func(*args)

    def __str__(self):
        if self.is_static:
            return 'Route(static,%s,path=%s)' % (self.method, self.path)
        return 'Route(dynamic,%s,path=%s)' % (self.method, self.path)

    __repr__ = __str__


class StaticFileRoute(object):
    def __init__(self):
        self.method = 'GET'
        self.is_static = False
        self.route = re.compile('^/static/(.+)$')

    @staticmethod
    def match(url):
        if url.startswith('/static/'):
            return (url[1:],)
        return None

    def __call__(self, *args):

        def _static_file_generator(fpath):
            BLOCK_SIZE = 8192
            with open(fpath, 'rb') as f:
                block = f.read(BLOCK_SIZE)
                while block:
                    yield block
                    block = f.read(BLOCK_SIZE)

        fpath = os.path.join(ctx.application['document_root'], args[0])
        if not os.path.isfile(fpath):
            raise error.notfound()
        fext = os.path.splitext(fpath)[1]
        ctx.response.content_type = mimetypes.types_map.get(fext.lower(), 'application/octet-stream')

        return _static_file_generator(fpath)


class Template(object):
    def __init__(self, template_name, **kw):
        """
        Init a template object with template name, model as dict, and additional kw that will append to model.
        """
        self.template_name = template_name
        self.model = dict(**kw)


class TemplateEngine(object):
    """
    Base template engine.
    """

    def __call__(self, path, model):
        return '<!-- override this method to render template -->'


class Jinja2TemplateEngine(TemplateEngine):
    """
    Render using jinja2 template engine.
    """

    def __init__(self, templ_dir, **kw):
        from jinja2 import Environment, FileSystemLoader

        kw['autoescape'] = True if 'autoescape' not in kw else False
        self._env = Environment(loader=FileSystemLoader(templ_dir), **kw)

    def add_filter(self, name, fn_filter):
        self._env.filters[name] = fn_filter

    def __call__(self, path, model):
        template = self._env.get_template(path)
        return template.render(**model)


def view(path):
    """
    A view decorator that render a view by dict.
    >>> @view('test/view.html')
    ... def hello():
    ...     return dict(name='Bob')
    """

    def _decorator(func):
        @functools.wraps(func)
        def _wrapper(*args, **kw):
            r = func(*args, **kw)
            if isinstance(r, dict):
                # logger.info('return Template')
                return Template(path, **r)
            raise ValueError('Expect return a dict when using @view() decorator.')

        return _wrapper

    return _decorator


def interceptor(pattern='/'):
    """
    An interceptor decorator.
    >>> @interceptor('/admin/')
    ... def check_admin(req, resp):
    ...     pass
    """

    def _build_pattern_fn(pattern):
        _RE_INTERCEPTROR_STARTS_WITH = re.compile(r'^([^\*\?]+)\*?$')
        _RE_INTERCEPTROR_ENDS_WITH = re.compile(r'^\*([^\*\?]+)$')

        m = _RE_INTERCEPTROR_STARTS_WITH.match(pattern)
        if m:
            return lambda p: p.startswith(m.group(1))
        m = _RE_INTERCEPTROR_ENDS_WITH.match(pattern)
        if m:
            return lambda p: p.endswith(m.group(1))
        raise ValueError('Invalid pattern definition in interceptor.')

    def _decorator(func):
        func.__interceptor__ = _build_pattern_fn(pattern)
        return func

    return _decorator


def _build_interceptor_fn(func, next):
    def _wrapper():
        if func.__interceptor__(ctx.request.path_info):
            return func(next)
        else:
            return next()

    return _wrapper


def _build_interceptor_chain(last_fn, *interceptors):
    """
    Build interceptor chain.
    _build_interceptor_chain(target, f1, f2, f3) => f1(f2(f3(target)))
    """
    L = list(interceptors)
    L.reverse()  # reverse to get proper interceptors order
    fn = last_fn
    for f in L:
        fn = _build_interceptor_fn(f, fn)
    return fn


class WSGIApplication(object):
    def __init__(self, document_root=None, **kw):
        """
        Init a WSGIApplication.
        Args:
          document_root: document root path.
        """
        self._running = False
        self._document_root = document_root

        self._interceptors = []
        self._template_engine = None

        self._get_static = {}
        self._post_static = {}

        self._get_dynamic = []
        self._post_dynamic = []

    def _check_not_running(self):
        if self._running:
            raise RuntimeError('Cannot modify WSGIApplication when running.')

    @property
    def template_engine(self):
        return self._template_engine

    @template_engine.setter
    def template_engine(self, engine):
        self._check_not_running()
        self._template_engine = engine

    def add_module(self, mod):
        def _load_module(module_name):
            """
            Load module from name as str.
            __import__: dynamic import module
            """
            last_dot = module_name.rfind('.')
            if last_dot == (-1):
                return __import__(module_name, globals(), locals())
            from_module = module_name[:last_dot]
            sub_module = module_name[last_dot + 1:]
            m = __import__(from_module, globals(), locals(), [sub_module])
            return getattr(m, sub_module)

        self._check_not_running()
        if isinstance(mod, types.ModuleType):
            m = mod
        else:
            m = _load_module(mod)
        logger.info('Add module: %s' % m.__name__)
        for name in dir(m):
            fn = getattr(m, name)
            if callable(fn) and hasattr(fn, '__web_route__') and hasattr(fn, '__web_method__'):
                self.add_url(fn)

    def add_url(self, func):
        self._check_not_running()
        route = Route(func)
        if route.is_static:
            if route.method == 'GET':
                self._get_static[route.path] = route
            if route.method == 'POST':
                self._post_static[route.path] = route
        else:
            if route.method == 'GET':
                self._get_dynamic.append(route)
            if route.method == 'POST':
                self._post_dynamic.append(route)
        logger.info('Add route: %s' % str(route))

    def add_interceptor(self, func):
        self._check_not_running()
        self._interceptors.append(func)
        logger.info('Add interceptor: %s' % str(func))

    def run(self, port=9000, host='127.0.0.1'):
        from wsgiref.simple_server import make_server
        logger.info('application (%s) will start at %s:%s...' % (self._document_root, host, port))
        server = make_server(host, port, self.get_wsgi_application(debug=True))
        server.serve_forever()

    def tornado_run(self, port=9000, host='127.0.0.1'):
        from tornado.httpserver import HTTPServer
        from tornado.wsgi import WSGIContainer
        from tornado.ioloop import IOLoop
        from tornado.options import options

        options.parse_command_line()
        server = HTTPServer(WSGIContainer(self.get_wsgi_application(debug=True)))
        server.listen(port, host)
        IOLoop.instance().start()

    def get_wsgi_application(self, debug=False):
        self._check_not_running()
        if debug:
            self._get_dynamic.append(StaticFileRoute())
        self._running = True

        _application = {'document_root': self._document_root}

        def fn_route():
            request_method = ctx.request.request_method
            path_info = ctx.request.path_info
            if request_method == 'GET':
                fn = self._get_static.get(path_info, None)
                if fn:
                    return fn()
                for fn in self._get_dynamic:
                    args = fn.match(path_info)
                    if args:
                        return fn(*args)
                raise error.notfound()
            if request_method == 'POST':
                fn = self._post_static.get(path_info, None)
                if fn:
                    return fn()
                for fn in self._post_dynamic:
                    args = fn.match(path_info)
                    if args:
                        return fn(*args)
                raise error.notfound()
            raise error.badrequest()

        fn_exec = _build_interceptor_chain(fn_route, *self._interceptors)

        def wsgi(env, start_response):
            """
            WSGI function
            return value should be iterable(bytes)
            """
            # init ctx
            ctx.application = _application
            ctx.request = Request(env)
            response = ctx.response = Response()
            try:
                r = fn_exec()
                if isinstance(r, Template):
                    r = self._template_engine(r.template_name, r.model)
                if isinstance(r, str):
                    r = [r.encode('utf-8')]
                if isinstance(r, bytes):
                    r = [r]
                if r is None:
                    r = []
                start_response(response.status, response.headers)
                return r
            except error.RedirectError as e:
                response.set_header('Location', e.location)
                start_response(e.status, response.headers)
                return []
            except error.HttpError as e:
                start_response(e.status, response.headers)
                return [b'<html><body><h1>', to_byte(e.status), b'</h1></body></html>']
            except Exception as e:
                logger.warning(e)
                if not debug:
                    start_response('500 Internal Server Error', [])
                    return [b'<html><body><h1>500 Internal Server Error</h1></body></html>']
                exc_type, exc_value, exc_traceback = sys.exc_info()
                fp = BytesIO()
                traceback.print_exception(exc_type, exc_value, exc_traceback, file=fp)
                stacks = fp.getvalue()
                fp.close()
                start_response('500 Internal Server Error', [])
                return [
                    b'''<html><body><h1>500 Internal Server Error</h1><div style="font-family:Monaco, Menlo, Consolas, 'Courier New', monospace;"><pre>''',
                    stacks.replace(b'<', b'&lt;').replace(b'>', b'&gt;'),
                    b'</pre></div></body></html>']
            finally:
                del ctx.application
                del ctx.request
                del ctx.response

        return wsgi
