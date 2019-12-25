# 有关自定义异常：http://blog.csdn.net/kwsy2008/article/details/48468345
# __setattr__()：拦截所有属性的赋值语句
# 属性私有化的首选方式

class PrivateExc(Exception):
    def __init__(self, err="private attribute is forbidden to modify"):
        Exception.__init__(self, err)


class Privacy:
    def __setattr__(self, key, value):
        if key in self.privates:
            raise PrivateExc()
        else:
            self.__dict__[key] = value


class Test1(Privacy):
    privates = ["age"]


class Test2(Privacy):
    privates = ["name", "age"]

    def __init__(self):
        self.__dict__["name"] = "sun"


if __name__ == '__main__':
    x = Test1()
    y = Test2()
    x.name = "Tom"
    x.age = 20
    # y.name = "Mike"
    # y.age = 20
