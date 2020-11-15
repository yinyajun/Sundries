#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/12/22 9:29
# @Author  : Yajun Yin
# @Note    :


from web.util import *


class Response(object):
    def __init__(self):
        self._status = '%d %s' % (200, RESPONSE_STATUSES[200])
        self._headers = {'CONTENT-TYPE': 'text/html; charset=utf-8'}

    @property
    def headers(self):
        """
        Return response headers as [(key1, value1), (key2, value2)...] including cookies.
        """
        L = [(RESPONSE_HEADER_DICT.get(k, k), v) for k, v in self._headers.items()]
        if hasattr(self, '_cookies'):
            for v in self._cookies.itervalues():
                L.append(('Set-Cookie', v))
        L.append(HEADER_X_POWERED_BY)
        return L

    def header(self, name):
        """
        Get header by name, case-insensitive.
        """
        key = name.upper()
        if key not in RESPONSE_HEADER_DICT:
            key = name
        return self._headers.get(key)

    def unset_header(self, name):
        """
        Unset header by name and value.
        """
        key = name.upper()
        if key not in RESPONSE_HEADER_DICT:
            key = name
        if key in self._headers:
            del self._headers[key]

    def set_header(self, name, value):
        """
        Set header by name and value.
        """
        key = name.upper()
        if key not in RESPONSE_HEADER_DICT:
            key = name
        self._headers[key] = to_unicode(value)

    @property
    def content_type(self):
        """
        Get content type from response. This is a shortcut for header('Content-Type').
        """
        return self.header('CONTENT-TYPE')

    @content_type.setter
    def content_type(self, value):
        """
        Set content type for response. This is a shortcut for set_header('Content-Type', value).
        """
        if value:
            self.set_header('CONTENT-TYPE', value)
        else:
            self.unset_header('CONTENT-TYPE')

    @property
    def content_length(self):
        """
        Get content length. Return None if not set.
        """
        return self.header('CONTENT-LENGTH')

    @content_length.setter
    def content_length(self, value):
        """
        Set content length, the value can be int or str.
        """
        self.set_header('CONTENT-LENGTH', str(value))

    def delete_cookie(self, name):
        """
        Delete a cookie immediately.
        """
        self.set_cookie(name, '__deleted__', expires=0)

    def set_cookie(self, name, value, max_age=None, expires=None, path='/', domain=None, secure=False, http_only=True):
        """
        Set a cookie.
        Args:
          name: the cookie name.
          value: the cookie value.
          max_age: optional, seconds of cookie's max age.
          expires: optional, unix timestamp, datetime or date object that indicate an absolute time of the
                   expiration time of cookie. Note that if expires specified, the max_age will be ignored.
          path: the cookie path, default to '/'.
          domain: the cookie domain, default to None.
          secure: if the cookie secure, default to False.
          http_only: if the cookie is for http only, default to True for better safty
                     (client-side script cannot access cookies with HttpOnly flag).
        """
        if not hasattr(self, '_cookies'):
            self._cookies = {}
        L = ['%s=%s' % (quote(name), quote(value))]
        if expires is not None:
            if isinstance(expires, (float, int)):
                L.append('Expires=%s' % datetime.datetime.fromtimestamp(expires, UTC('+00:00'))
                         .strftime('%a, %d-%b-%Y %H:%M:%S GMT'))
            if isinstance(expires, (datetime.date, datetime.datetime)):
                L.append('Expires=%s' % expires.astimezone(UTC('+00:00'))
                         .strftime('%a, %d-%b-%Y %H:%M:%S GMT'))
        elif isinstance(max_age, int):
            L.append('Max-Age=%d' % max_age)
        L.append('Path=%s' % path)
        if domain:
            L.append('Domain=%s' % domain)
        if secure:
            L.append('Secure')
        if http_only:
            L.append('HttpOnly')
        self._cookies[name] = '; '.join(L)

    def unset_cookie(self, name):
        """
        Unset a cookie.
        """
        if hasattr(self, '_cookies'):
            if name in self._cookies:
                del self._cookies[name]

    @property
    def status_code(self):
        """
        Get response status code as int.
        """
        return int(self._status[:3])

    @property
    def status(self):
        """
        Get response status. Default to '200 OK'.
        """
        return self._status

    @status.setter
    def status(self, value):
        """
        Set response status as int or str.
        """
        if isinstance(value, (int,)):
            if 100 <= value <= 999:
                st = RESPONSE_STATUSES.get(value, '')
                if st:
                    self._status = '%d %s' % (value, st)
                else:
                    self._status = str(value)
            else:
                raise ValueError('Bad response code: %d' % value)
        elif isinstance(value, (str, bytes)):
            if isinstance(value, bytes):
                value = value.decode('utf-8')
            if RE_RESPONSE_STATUS.match(value):
                self._status = value
            else:
                raise ValueError('Bad response code: %s' % value)
        else:
            raise TypeError('Bad type of response code.')
