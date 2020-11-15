#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/12/22 9:25
# @Author  : Yajun Yin
# @Note    :

"""
python2: type `str` save as [bytes]
python3: type `str` save as [unicode]

encode func: [unicode] ---> [bytes]
decode func: [bytes] ---> [unicode]

python2:
        s = 'hello'  # s is [bytes]
        s.encode('utf-8')
    # This may be very strange, how encode func works?
    # Because Python2 convert bytes to unicode secretly.
    # s.encode('utf-8')  <-------> s.decode(defaultencoding).encode("utf-8")
    # No wonder you will see code like this to set defaultencoding in Python2:
        reload(sys)
        sys.setdefaultencoding('utf-8')

python3:
    it is ok


conclusion:
    python2 `str` -> python3 `bytes`
    python2 `unicode` -> python3 `str`
"""

import os
import logging
import urllib.request
from logging import handlers

from web.constant import *


def to_byte(s):
    """
    to byte
    Convert to str.
    >>> to_str('s123') == b's123'
    True
    >>> to_str(u'\u4e2d\u6587') == '\xe4\xb8\xad\xe6\x96\x87'
    True
    >>> to_str(-123) == b'-123'
    True
    """
    if isinstance(s, bytes):
        return s
    if isinstance(s, str):
        return s.encode('utf-8')
    return str(s).encode('utf-8')


def to_unicode(s, encoding='utf-8'):
    """
    bytes to str
    Convert to unicode.
    >>> to_unicode(b'\xe4\xb8\xad\xe6\x96\x87') == u'\u4e2d\u6587'
    True
    """
    if isinstance(s, str):
        return s
    return s.decode(encoding)


def quote(s, encoding='utf-8'):
    """
    input bytes or unicode
    return unicode

    Url quote as bytes.
    >>> quote('http://example/test?a=1+')
    'http%3A//example/test%3Fa%3D1%2B'
    >>> quote(u'hello world!')
    'hello%20world%21'
    """
    if isinstance(s, bytes):
        s = s.decode(encoding)
    return urllib.request.quote(s)


def unquote(s, encoding="utf-8"):
    """
    s: unicode or bytes
    return : unicode

    Url unquote as unicode.
    >>> unquote('http%3A//example/test%3Fa%3D1+')
    u'http://example/test?a=1+'
    """
    if isinstance(s, bytes):
        s = s.decode(encoding)
    return urllib.request.unquote(s)


class UTC(datetime.tzinfo):
    """
    A UTC tzinfo object.
    >>> tz0 = UTC('+00:00')
    >>> tz0.tzname(None)
    'UTC+00:00'
    >>> tz8 = UTC('+8:00')
    >>> tz8.tzname(None)
    'UTC+8:00'
    >>> tz7 = UTC('+7:30')
    >>> tz7.tzname(None)
    'UTC+7:30'
    >>> tz5 = UTC('-05:30')
    >>> tz5.tzname(None)
    'UTC-05:30'
    >>> from datetime import datetime
    >>> u = datetime.utcnow().replace(tzinfo=tz0)
    >>> l1 = u.astimezone(tz8)
    >>> l2 = u.replace(tzinfo=tz8)
    >>> d1 = u - l1
    >>> d2 = u - l2
    >>> d1.seconds
    0
    >>> d2.seconds
    28800
    """

    def __init__(self, utc):
        utc = str(utc.strip().upper())
        mt = RE_TZ.match(utc)
        if mt:
            minus = mt.group(1) == '-'
            h = int(mt.group(2))
            m = int(mt.group(3))
            if minus:
                h, m = (-h), (-m)
            self._utcoffset = datetime.timedelta(hours=h, minutes=m)
            self._tzname = 'UTC%s' % utc
        else:
            raise ValueError('bad utc time zone')

    def utcoffset(self, dt):
        return self._utcoffset

    def dst(self, dt):
        return TIMEDELTA_ZERO

    def tzname(self, dt):
        return self._tzname

    def __str__(self):
        return 'UTC tzinfo object (%s)' % self._tzname

    __repr__ = __str__


def dir_path(path):
    d = os.path.dirname(os.path.abspath(path))
    return d


class Logger(object):
    def __init__(self, name, log_lvl=None, filename=None, split=None):
        self._formatter = logging.Formatter('%(asctime)s\t%(name)s\t%(levelname)s\t%(message)s', "%Y-%m-%d %H:%M:%S")
        self._logger = logging.getLogger(name)
        self._log_name = filename
        self._split = "\t" if split is None else split
        level = logging.INFO if log_lvl is None else log_lvl
        self._logger.setLevel(level)
        self.add_stream_handler()

    def add_stream_handler(self):
        sh = logging.StreamHandler()
        sh.setFormatter(self._formatter)
        self._logger.addHandler(sh)

    def add_file_handler(self):
        if self._log_name is None:
            return
        fh = handlers.TimedRotatingFileHandler(self._log_name, when='D', backupCount=5)
        fh.setFormatter(self._formatter)
        self._logger.addHandler(fh)

    def _parse(self, *args):
        return self._split.join(map(lambda p: str(p), args))

    def info(self, *args):
        return self._logger.info(self._parse(*args))

    def debug(self, *args):
        return self._logger.debug(self._parse(*args))

    def warning(self, *args):
        return self._logger.warning(self._parse(*args))

    def error(self, *args):
        return self._logger.error(self._parse(*args))
