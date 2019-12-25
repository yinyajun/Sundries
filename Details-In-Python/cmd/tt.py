from __future__ import print_function

from cmd import Cmd
import datetime
from pyspark.sql import functions as F


class SparkData(object):
    def __init__(self, day, interval=1):
        self.day = str(day)
        self.interval = interval
        days, self.period = self.get_period(self.day, self.interval)
        self.source_df = self.source_data(days)
        # flag
        self.is_filter = False
        self.filtered_source_df = None
        self.report = None
        # cache
        self.source_df.cache()
        print("query on %s" % self.day)

    @staticmethod
    def get_period(day, interval):
        target_day = datetime.datetime.strptime(day, "%Y%m%d")
        period = []
        for i in range(interval):
            d = (target_day - datetime.timedelta(i)).strftime("%Y%m%d")
            period.append(d)
        return ','.join(period), period

    @staticmethod
    def source_data(days):
        sql = '''
            select 
                cast(user_id as int) as uid,
                cast(item_id as int) as pid,
                cast(score as int) as score , 
                cast(user_id % 10 as int) as tail, 
                day
            from 
                user_item_score_tbl
            where 
                day in ({days}) 
        '''.format(days=days)
        return sqlContext.sql(sql)

    def filter_source_data(self):
        self.filtered_source_df = self.source_df.filter('score>=3 and score<=50')
        self.is_filter = True
        self.filtered_source_df.cache()
        print("After filter, count is:", self.filtered_source_df.count())

    def de_filter_data(self):
        print("source data, count is:", self.source_df.count())
        self.is_filter = False

    def get_data(self):
        return self.filtered_source_df if self.is_filter else self.source_df

    def show_fitler(self):
        if self.is_filter:
            msg = "数据已经过滤"
        else:
            msg = "数据没有过滤"
        print(msg)

    def show_max_score(self, truncate=50):
        self.show_fitler()
        df = self.get_data()
        df.orderBy(F.desc('score')).show(truncate)

    def show_user_watchtime_by_uid(self, uid):
        self.show_fitler()
        df = self.get_data()
        df.filter("uid=%s" % str(uid)).show()

    def show_user_watchtime_by_tail(self, tail):
        self.show_fitler()
        df = self.get_data()
        df.filter("tail=%s" % str(tail)) \
            .groupBy('uid') \
            .agg(F.sum('score').alias('score')) \
            .orderBy(F.desc('score')).show()

    def report_form(self):
        self.show_fitler()
        df = self.get_data()
        self.report = df.groupBy(['day', 'tail']) \
            .agg(F.sum('score').alias('total_score'),
                 F.count('uid').alias("cnt"),
                 F.countDistinct("uid").alias("uid_cnt"),
                 F.countDistinct("pid").alias("pid_cnt")) \
            .orderBy([F.asc('day'), F.asc('tail')])
        self.report.show(self.interval * 10)

    def analyse_report(self):
        if not self.report:
            print("请先生成报表，才能分析")
            raise ValueError
        self.show_fitler()
        df = self.report
        df = df.withColumn('avg_user_score', df['total_score'] / df['uid_cnt'])
        df = df.withColumn('avg_user_cnt', df['cnt'] / df['uid_cnt'])
        df = df.withColumn('avg_user_peo_num', df.pid_cnt / df['uid_cnt'])
        self.report = df
        self.report.show()


class ReportCmd(Cmd):
    def __init__(self, date, interval=1):
        Cmd.prompt = "ReportCmd>"
        Cmd.intro = "welcome to ReportCmd"
        Cmd.__init__(self)
        self.data = SparkData(date, interval)
        self.func_names = []
        self.funcs = []

        names = ['过滤', '显示报表', '报表分析', '按尾号查score', '按uid查score', '显示TopN的score']
        funcs = ['filter_source_data', 'report_form', 'analyse_report', 'show_user_watchtime_by_tail',
                 'show_user_watchtime_by_uid', 'show_max_score']
        self.index = 0
        for f1, f2 in zip(names, funcs):
            self.register_func(f1, f2)
            self.index += 1

    def register_func(self, func_name, func):
        self.func_names.append(func_name)
        self.funcs.append(func)
        print(self.index, func_name, func)

    def do_show(self, arg):
        print("")
        index = 0
        for i, j in zip(self.func_names, self.funcs):
            print(index, i, j)
            index += 1

    def emptyline(self):
        pass

    def help_show(self):
        print("输入:show,显示功能")

    def help_exit(self):
        print("输入:exit,退出程序")

    def help_func(self):
        print("输入:func+数字,执行功能")

    def do_exit(self, line):
        print("再见")
        raise KeyboardInterrupt

    def exec_func(self, index):

        index = int(index)
        print(index, self.func_names[index], self.funcs[index])
        try:
            getattr(self.data, self.funcs[index])()
        except AttributeError:
            print("%s不支持!" % self.func_names[index])
        except ValueError:
            print("请重新选择功能")

    def do_func(self, line):
        try:
            index = int(line)
            self.exec_func(index)
        except ValueError:
            print("请输入数字,输入show查看支持功能")
        except IndexError:
            print("不存在对应功能,输入show查看支持功能")

    def default(self, line):
        print("无效的输入，请输入：func+数字 执行功能，如：func 0")


cmd = ReportCmd(20190321, 1)
cmd.cmdloop()
