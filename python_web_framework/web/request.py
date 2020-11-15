#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/12/22 9:29
# @Author  : Yajun Yin
# @Note    :

import cgi

from web.util import *


class MultipartFile(object):
    """
    Multipart file storage get from request input.
    f = ctx.request['file']
    f.filename # 'test.png'
    f.file # file-like object
    """

    def __init__(self, storage):
        self.filename = to_unicode(storage.filename)
        self.file = storage.file


class Request(object):
    """
    Request object for obtaining all http request information.
    """

    def __init__(self, environ):
        self._environ = environ

    def _parse_input(self):
        """
        use cgi.FieldStorage to parse input, be called only once
        :return dict
        """

        def _convert(item):
            if isinstance(item, list):
                return [to_unicode(i.value) for i in item]
            if item.filename:
                return MultipartFile(item)
            return to_unicode(item.value)

        fp = self._environ["wsgi.input"]
        self.fs = cgi.FieldStorage(
            fp,
            environ=self._environ,
            keep_blank_values=True)  # prevents fs object from gc
        inputs = {key: _convert(self.fs[key]) for key in self.fs}
        return inputs

    def _get_raw_input(self):
        """
        Get raw input as dict containing values as unicode, list or MultipartFile.
        """
        if not hasattr(self, "_raw_input"):
            self._raw_input = self._parse_input()

        return self._raw_input

    def __getitem__(self, key):
        """
        Get input parameter value. If the specified key has multiple value, the first one is returned.
        If the specified key is not exist, then raise KeyError.
        """
        r = self._get_raw_input()[key]
        if isinstance(r, list):
            return r[0]
        return r

    def get(self, key, default=None):
        """
        The same as request[key], but return default value if key is not found.
        """
        r = self._get_raw_input().get(key, default)
        if isinstance(r, list):
            return r[0]
        return r

    def gets(self, key):
        """
        Get multiple values for specified key.
        """
        r = self._get_raw_input()[key]
        if isinstance(r, list):
            return r[:]
        return [r]

    def input(self, **kw):
        """
        Get input as dict from request, fill dict using provided default value if key not exist.
        """
        copy = dict(**kw)
        raw = self._get_raw_input()
        for k, v in raw.items():
            copy[k] = v[0] if isinstance(v, list) else v
        return copy

    def get_body(self):
        """
        Get raw data from HTTP POST and return as bytes.
        """
        fp = self._environ['wsgi.input']
        return fp.read()

    @property
    def remote_addr(self):
        """
        Get remote addr. Return '0.0.0.0' if cannot get remote_addr.
        """
        return self._environ.get("REMOTE_ADDR", '0.0.0.0')

    @property
    def document_root(self):
        """
        Get raw document_root as unicode. Return '' if no document_root.
        """
        return self._environ.get("DOCUMENT_ROOT", '')

    @property
    def query_string(self):
        """
        Get raw query string as unicode. Return '' if no query string.
        """
        return self._environ.get("QUERY_STRING", '')

    @property
    def environ(self):
        """
        Get raw environ as dict, both key, value are unicode.
        """
        return self._environ

    @property
    def request_method(self):
        """
        Get request method. The valid returned values are 'GET', 'POST', 'HEAD'.
        """
        return self._environ["REQUEST_METHOD"]

    @property
    def path_info(self):
        """
        Get request path as unicode.
        """
        return unquote(self._environ.get("PATH_INFO", ''))

    @property
    def host(self):
        """
        Get request host as unicode. Default to '' if cannot get host..

        """
        return self._environ.get("HTTP_HOST", '')

    def _get_headers(self):
        if not hasattr(self, '_headers'):
            headers = {}
            for k, v in self._environ.items():
                if k.startswith('HTTP_'):  # convert 'HTTP_ACCEPT_ENCODING' to 'ACCEPT-ENCODING'
                    headers[k[5:].replace('_', '-').upper()] = to_unicode(v)
            self._headers = headers
        return self._headers

    @property
    def headers(self):
        """
        Get all HTTP headers with key as str and value as unicode. The header names are 'XXX-XXX' uppercase.
        """
        return dict(**self._get_headers())

    def header(self, header, default=None):
        """
        Get header from request as unicode, return None if not exist, or default if specified.
        The header name is case-insensitive such as 'USER-AGENT' or 'content-Type'.
        """
        return self._get_headers().get(header.upper(), default)

    def _get_cookies(self):
        if not hasattr(self, '_cookies'):
            cookies = {}
            cookie_str = self._environ.get("HTTP_COOKIE")
            if cookie_str:
                for c in cookie_str.split(';'):
                    pos = c.find('=')
                    if pos > 0:
                        cookies[c[:pos].strip()] = unquote(c[pos + 1:])
            self._cookies = cookies
        return self._cookies

    @property
    def cookies(self):
        """
        Return all cookies as dict. The cookie name is str and values is unicode.
        """
        return dict(**self._get_cookies())

    def cookie(self, name, default=None):
        """
        Return specified cookie value as unicode. Default to None if cookie not exists.
        """
        return self._get_cookies().get(name, default)

# if __name__ == '__main__':
#     from io import BytesIO

# boundary = '----WebKitFormBoundaryQQ3J8kPsjFpTmqNz'
#     post_data = """------WebKitFormBoundaryQQ3J8kPsjFpTmqNz
# Content-Disposition: form-data; name="name"
#
# Scofield
# ------WebKitFormBoundaryQQ3J8kPsjFpTmqNz
# Content-Disposition: form-data; name="name"
#
# Lincoln
# ------WebKitFormBoundaryQQ3J8kPsjFpTmqNz
# Content-Disposition: form-data; name="file"; filename="test.txt"
# Content-Type: text/plain
#
# just a test
# ------WebKitFormBoundaryQQ3J8kPsjFpTmqNz
# Content-Disposition: form-data; name="id"
#
# 4008009001
# ------WebKitFormBoundaryQQ3J8kPsjFpTmqNz--"""
#     env = {
#         'REQUEST_METHOD': 'POST',
#         'CONTENT_TYPE': 'multipart/form-data; boundary={}'.format(boundary),
#         'CONTENT_LENGTH': str(len(post_data)),
#         'wsgi.input': BytesIO(post_data.encode('utf-8'))
#     }
#     r = Request(env)
#     print(r.get("name"))
#     f = r.get("file")
#     print(f.filename)
#     print(f.file.read())

# r = Request({'REQUEST_METHOD': 'GET', 'wsgi.url_scheme': 'http'})
# print(r.environ)
