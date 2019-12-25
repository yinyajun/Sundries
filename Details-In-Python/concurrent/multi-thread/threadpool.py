#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/5/21 19:34
# @Author  : Yajun Yin
# @Note    :

import io
import requests
import threadpool
import multiprocessing
from PIL import Image

PATH = 'sfdafgsgarg'


def download_img(liveid, url):
    # liveid, url = pair
    response = requests.get(url, timeout=(1, 3))  # timeout for connect and read
    if response.status_code != 200:
        return
    content = response.content
    if content is None or len(content) == 0:
        return
    try:
        im = Image.open(io.BytesIO(content))  # BytesIO(二进制)作为文件来读写
        im.save(PATH % liveid)
    except Exception as e:
        print("download_imgs()", liveid, url, e)


cpu_num = multiprocessing.cpu_count()
cores = cpu_num // 2
if cores == 0:
    cores = 1

total_urls = [(['123', 'www.fdsaoif.dsfafsd'], None), (['123', 'www.fdsaoif.dsfafsd'], None),
              (['123', 'www.fdsaoif.dsfafsd'], None)]
pool = threadpool.ThreadPool(cores)
reqs = threadpool.makeRequests(download_img, total_urls)
map(pool.putRequest, reqs)
pool.wait()
