#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/2/9 17:39
# @Author  : Yajun Yin
# @Note    : modified based on Programming Collective Intelligence pp. 142

import xml.dom.minidom
from urllib import request
from treepredict import *

zwskey = "X1-ZWz1chwxis15aj_9skq6"


def get_address_data(address, city):
    escad = address.replace(' ', '+')

    url = 'http://www.zillow.com/webservice/GetDeepSearchResults.htm?'
    url += 'zws-id=%s&address=%s&citystatezip=%s' % (zwskey, escad, city)
    print(url)

    doc = xml.dom.minidom.parseString(request.urlopen(url).read())
    code = doc.getElementsByTagName('code')[0].firstChild.data

    if code != '0':
        return None

    try:
        zipcode = doc.getElementsByTagName('zipcode')[0].firstChild.data
        use = doc.getElementsByTagName('useCode')[0].firstChild.data
        year = doc.getElementsByTagName('yearBuilt')[0].firstChild.data
        bath = doc.getElementsByTagName('bathrooms')[0].firstChild.data
        bed = doc.getElementsByTagName('bedrooms')[0].firstChild.data
        rooms = doc.getElementsByTagName('totalRooms')[0].firstChild.data
        price = doc.getElementsByTagName('amount')[0].firstChild.data
    except:
        return None

    return zipcode, use, year, bath, bed, rooms, price


def get_price_list():
    ll = []
    for line in open('addresslist.txt'):
        data = get_address_data(line.strip(), city='Cambridge,MA')
        if data:
            ll.append(data)
    return ll


def house_tree():
    data = get_price_list()
    tree = build_tree(data, 'var')
    draw_tree(tree, 'house_dt.jpeg')
