#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/2/9 23:39
# @Author  : Yajun Yin
# @Note    : modified based on Programming Collective Intelligence pp. 142
import re
from urllib import request
from urllib.parse import urljoin
import sqlite3
import bs4
from bs4 import BeautifulSoup

stop_words = {'the', 'of', 'to', 'and', 'a', 'in', 'is', 'it'}


class crawler:
    def __init__(self, dbname):
        self.conn = sqlite3.connect(dbname)

    def __del__(self):
        self.conn.close()

    def dbcommit(self):
        self.conn.commit()

    def get_entry_id(self, table, field, value, create_new=True):
        """
        返回某一条目的ID，如果条目不存在，则程序会在数据库中新建一条记录，并返回ID
        :param table:
        :param field:
        :param value:
        :param create_new:
        :return:
        """
        cur = self.conn.execute("select rowid from %s where %s='%s'" % (table, field, value))
        res = cur.fetchone()
        if res is None:
            cur = self.conn.execute("insert into %s (%s) values ('%s')" % (table, field, value))
            return cur.lastrowid
        else:
            return res[0]  # res[0] is rowid

    def add_to_index(self, url, soup):
        """
        得到一个出现于网页中的单词列表，索引网页；索引所有单词；在网页和单词之间建立关联，保存单词在文档中出现的位置（列表中的索引号）
        :param url:
        :param soup:
        :return:
        """
        if self.is_indexed(url):
            return
        print('Indexing %s' % url)
        # 获取每个单词
        text = self.get_text_only(soup)
        words = self.separate_words(text)
        # 得到URL的id
        url_id = self.get_entry_id('urllist', 'url', url)

        for i in range(len(words)):
            word = words[i]
            if word in stop_words:
                continue
            # 获得word的id
            word_id = self.get_entry_id('wordlist', 'word', word)
            # 对一个网页而言，将urlid，wordid和word_location(就是这个word在words中的位置,+1为了保持和rowid一致)
            self.conn.execute(
                "insert into wordlocation(urlid,wordid,location) values(%d,%d,%d)" % (url_id, word_id, i + 1))

    def get_text_only(self, soup):
        """
        递归向下对HTML的DOM树遍历，保留了各章节文本的前后顺序;由于还有分词的步骤，不需要担心没有处理'\n','\t'。
        :param soup:
        :return:
        """
        v = soup.string
        if v is None:  # 不含有字符，就可能嵌套有tag
            c = soup.contents  # tag或是navigableString的list
            # 如果就仅仅是没有string，那么这里的c就是一个[],直接返回一个''。
            result_text = ''
            for t in c:
                # doctype标签的内容也不算
                if t == '\n' or isinstance(t, bs4.element.Doctype):
                    continue
                subtext = self.get_text_only(t)
                result_text += subtext + '\n'
            return result_text
        else:
            return v.strip()

    def separate_words(self, text):
        """
        分词
        :param text:
        :return:
        """
        splitter = re.compile('\\W')
        return [s.lower() for s in splitter.split(text) if s != '']

    def is_indexed(self, url):
        """
        判断网页是否已经存入数据库，如果存在，判断是否有任何单词与之关联。
        :param url:
        :return:
        """
        u = self.conn.execute("select rowid from urllist where url ='%s'" % url).fetchone()
        if u is not None:
            # url已经存在urllist中
            v = self.conn.execute("select * from wordlocation where urlid = %d" % u[0]).fetchone()
            if v is not None:
                return True
        return False

    def add_link_ref(self, urlFrom, urlTo, linkText):
        """
        添加一个关联两个网页的链接
        :param urlFrom:
        :param urlTo:
        :param linkText:
        :return:
        """
        words = self.separate_words(linkText)
        from_id = self.get_entry_id('urllist', 'url', urlFrom)
        to_id = self.get_entry_id('urllist', 'url', urlTo)
        if from_id == to_id:
            return
        cur = self.conn.execute("insert into link(fromid, toid) values (%d, %d)" % (from_id, to_id))
        link_id = cur.lastrowid
        # linkwords表利用字段wordid和linkid记录哪些单词和链接实际相关
        for word in words:
            if word in stop_words:
                continue
            word_id = self.get_entry_id('wordlist', 'word', word)
            self.conn.execute('insert into linkwords(linkid, wordid) values (%d, %d)' % (link_id, word_id))

    def crawl(self, pages, depth=2):
        """
        广度优先的搜索方式
        :param pages:每一层深度的网页list
        :param depth:
        :return:
        """
        for i in range(depth):
            newpages = set()
            for page in pages:
                try:
                    c = request.urlopen(page)
                except:
                    print("open %s failed" % page)
                    continue
                soup = BeautifulSoup(c.read(), 'html.parser')
                # 去除soup中的script标签
                [s.extract() for s in soup('script')]
                self.add_to_index(page, soup)

                links = soup('a')
                for link in links:
                    # link是a标签的bs4的Tag，Tag的attrs属性得到所有属性的字典
                    if 'href' in link.attrs:
                        url = urljoin(page, link['href'])
                        if url.find("'") != -1:  # 找到了‘'’，这个url拼接的不对，不要了
                            continue
                        # 去除url位置标识符
                        url = url.split('#')[0]
                        # 判断url是否正确，并且没有被索引
                        if url[:4] == 'http' and not self.is_indexed(url):
                            newpages.add(url)
                        link_text = self.get_text_only(link)
                        self.add_link_ref(page, url, link_text)
                self.dbcommit()
            pages = newpages

    def create_index_tables(self):
        """
        为全文索引建立数据库，索引对应与一个列表，其中包含了所有不同的单词，单词所在的文档，以及单词在文档中出现位置
        :return:
        """
        self.conn.execute('create table urllist(url)')
        self.conn.execute('create table wordlist(word)')
        self.conn.execute('create table wordlocation(urlid,wordid,location)')
        self.conn.execute('create table link(fromid integer,toid integer)')
        self.conn.execute('create table linkwords(wordid,linkid)')
        self.conn.execute('create index wordidx on wordlist(word)')
        self.conn.execute('create index urlidx on urllist(url)')
        self.conn.execute('create index wordurlidx on wordlocation(wordid)')
        self.conn.execute('create index urltoidx on link(toid)')
        self.conn.execute('create index urlfromidx on link(fromid)')
        self.dbcommit()


class searcher:
    def __init__(self, dbname):
        self.conn = sqlite3.connect(dbname)

    def __del__(self):
        self.conn.close()

    def commit(self):
        self.conn.commit()

    def get_match_rows(self, q):
        """
        最朴素的搜索结果，通过sql完成。sql是多个wordlocation表的自身连接（cross join）
        join条件是每张表的wordid=分别的id，同时所有表的urlid相等（共同的url）
        :param q:
        :return:
        """
        fieldlist = 'w0.urlid'
        tablelist = ''
        clauselist = ''
        word_ids = []

        words = q.split(' ')
        table_number = 0
        for word in words:
            word_row = self.conn.execute("select rowid from wordlist where word='%s'" % word).fetchone()
            if word_row is not None:
                word_id = word_row[0]
                word_ids.append(word_id)
                if table_number > 0:
                    tablelist += ','
                    clauselist += ' and '
                    clauselist += ' w%d.urlid=w%d.urlid and ' % (table_number - 1, table_number)
                fieldlist += ',w%d.location' % table_number
                tablelist += 'wordlocation w%d' % table_number
                clauselist += 'w%d.wordid=%d' % (table_number, word_id)
                table_number += 1
        full_query = "select %s from %s where %s" % (fieldlist, tablelist, clauselist)
        # print(full_query)
        cur = self.conn.execute(full_query)
        rows = [row for row in cur]
        return rows, word_ids

    def get_scored_list(self, rows, word_ids):
        """
        检索顺序呈现的排序已经不能满足我们的需要，我们需要一种能够针对给定查询条件为网页进行评价的方法，
        并且能在返回结果中将评价最高者排在最前面。
        评价度量：单词频度，文档位置，单词距离
        :param rows:
        :param word_ids:
        :return:
        """
        total_scores = dict([(row[0], 0) for row in rows])  # 相同url去重
        # 不同的评价方法
        weights = [(1.0, self.frequency_score(rows)), (0.5, self.location_score(rows)),
                   (0.5, self.distance_score(rows)), (1.0, self.pagerank_score(rows)),
                   (1.0, self.link_text_score(rows, word_ids))]
        for weight, score in weights:
            for url in total_scores:
                total_scores[url] += weight * score[url]

        return total_scores

    def get_url_name(self, url_id):
        return self.conn.execute("select url from urllist where rowid=%d" % url_id).fetchone()[0]

    def query(self, q):
        rows, wordids = self.get_match_rows(q)
        scores = self.get_scored_list(rows, wordids)
        rankedscores = sorted([(score, url) for url, score in scores.items()], reverse=1)
        print(rankedscores)
        for score, urlid in rankedscores[0:10]:
            print('%f\t%s' % (score, self.get_url_name(urlid)))

    def normalize_scores(self, scores, small_is_better=0):
        """
        不同方法的返回结果进行比较，对结果归一化处理（0,1之间），1是最佳score
        :param scores:
        :param small_is_better:
        :return:
        """
        vsmall = 0.00001  # 避免被0整除
        if small_is_better:
            min_score = min(scores.values())
            return dict([(u, float(min_score) / max(vsmall, l)) for (u, l)
                         in scores.items()])
        else:
            max_score = max(scores.values())
            if max_score == 0:
                max_score = vsmall
            return dict([(u, float(c) / max_score) for u, c in scores.items()])

    def frequency_score(self, rows):
        """
        单词在网页中出现的次数对网页进行评价
        :param rows:
        :return:
        """
        counts = dict([(row[0], 0) for row in rows])
        for row in rows:
            counts[row[0]] += 1
        return self.normalize_scores(counts)

    def location_score(self, rows):
        """
        待查单词在文档中出现越早，评分越高
        :param rows:
        :return:
        """
        locations = dict([(row[0], 1000000) for row in rows])
        for row in rows:
            loc = sum(row[1:])
            if loc < locations[row[0]]:
                locations[row[0]] = loc
        return self.normalize_scores(locations, small_is_better=1)

    def distance_score(self, rows):
        """
        当查询多个单词时，单词彼此间距离很近的网页往往很有意义
        :param rows:
        :return:
        """
        # 如果只有一个单词，就是rows中每行就（urlid，word）
        if len(rows[0]) <= 2:
            return dict([(row[0], 1.0) for row in rows])

        min_distance = dict([(row[0], 1000000) for row in rows])
        for row in rows:
            dist = sum([abs(row[i] - row[i - 1]) for i in range(2, len(row))])
            if dist < min_distance[row[0]]:
                min_distance[row[0]] = dist
        return self.normalize_scores(min_distance, small_is_better=1)

    def in_bound_link_score(self, rows):
        """
        类似于论文引用，引用次数越多的越重要；（基于内容的评价方法-网页相关度；基于外部回指链接评价方法-网页重要度）
        但是每个外部回指链接拥有相同权重，没有突出热门网页的链接的贡献度。
        :param rows:
        :return:
        """
        unique_urls = set([row[0] for row in rows])
        # 注意这里sql执行完fetchone()后是tuple，要能通过取元素来得到对应项
        in_bound_count = dict(
            [(u, self.conn.execute("select count(*) from link where toid=%d" % u).fetchone()[0])
             for u in unique_urls])
        return self.normalize_scores(in_bound_count)

    def calculate_pagerank(self, iterations=20):
        # 清楚当前的pagerank表
        self.conn.execute("drop table if exists pagerank")
        self.conn.execute("create table pagerank(urlid primary key, score)")

        # 初始化每个url的pagerank为1.0
        self.conn.execute("insert into pagerank select rowid, 1.0 from urllist")
        self.commit()

        for i in range(iterations):
            print("Iteration %d" % i)
            # 因为执行sql返回的每一行是一个tuple，这样取出就是一个int
            for (urlid,) in self.conn.execute("select rowid from urllist"):
                pr = 0.15

                # 遍历指向当前网页的所有其他网页
                for (linker,) in self.conn.execute("select distinct fromid from link where toid=%d" % urlid):
                    # 链接源对应网页的pagerank值
                    linking_pr = self.conn.execute("select score from pagerank where urlid=%d" % linker).fetchone()[0]
                    # 链接源对应网页的向外链接数目
                    linking_count = self.conn.execute("select count(*) from link where fromid=%d" % linker).fetchone()[
                        0]
                    # 一步步计算这个网页的pagerank值
                    pr += 0.85 * (linking_pr / linking_count)
                self.conn.execute("update pagerank set score=%f where urlid=%d" % (pr, urlid))
                self.commit()

    def pagerank_score(self, rows):
        """
        离线计算好pagerank值，记录到表中。
        :param rows:
        :return:
        """
        unique_urls = set([row[0] for row in rows])
        pageranks = dict(
            [(url, self.conn.execute("select score from pagerank where urlid=%d" % url).fetchone()[0])
             for url in unique_urls])
        return self.normalize_scores(pageranks)

    def link_text_score(self, rows, word_ids):
        """
        链接文字通常更有意义，会是指向网页的解释。那么包含这些单词的链接，目标链接在搜索结果，则将原链接的pr值加入到目标网页的评价中。
        :param rows:
        :param word_ids: 待查的单词的ids
        :return:
        """
        link_scores = dict([(row[0], 0) for row in rows])
        for word_id in word_ids:
            # 和某单词有关的所有链接关系（fromid，toid）
            cur = self.conn.execute(
                "select link.fromid,link.toid from link,linkwords where wordid=%d and linkwords.linkid=link.rowid" % word_id)
            for fromid, toid in cur:
                # 在这些链接关系落入搜索结果
                if toid in link_scores:
                    pr = self.conn.execute("select score from pagerank where urlid=%d" % fromid).fetchone()[0]
                    link_scores[toid] += pr
        return self.normalize_scores(link_scores)




if __name__ == '__main__':
    # crawler = crawler('test.db')
    # crawler.create_index_tables()
    # pages = ["http://segaran.com/"]
    # crawler.crawl(pages)
    # print([row for row in crawler.conn.execute("select * from wordlist")])
    # print([row for row in crawler.conn.execute("select * from urllist")])
    # print([row for row in crawler.conn.execute("select * from wordlocation")])
    # print("******************")
    e = searcher('test.db')
    a, b = e.get_match_rows('index ico')
    cur = e.conn.execute("select * from pagerank order by score desc")
    print(a)
    print(b)
    print("******************")
    e.query('index ico')
    print("******************")
