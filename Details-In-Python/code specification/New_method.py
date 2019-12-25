 class parent(object):
    def __new__(cls):
        print('parent new')
        return super(parent, cls).__new__(cls)

    def __init__(self):
        print('parent')


class child(parent):
    def __new__(cls):
        print('child new')
        return super(child, cls).__new__(cls)

    def __init__(self):
        super(child, self).__init__()
        print('child')


bb = child()

class Child1(collections.namedtuple('parent', ['value'])):
    def __new__(cls, value):
        print('child', value)
        return super(Child1, cls).__new__(cls, value)


a = Child1(5)
print(a)
print(type(a))


# results:
# child new
# parent new
# parent
# child

# new()方法先于init() 方法，是类的静态方法，返回类的实例

# new()是在新式类中新出现的方法，它作用在构造方法init()建造实例之前，可以这么理解，在Python 中存在于类里面的构造方法init()负责将类的实例化，而在init()调用之前，new()决定是否要使用该init()方法，因为new()可以调用其他类的构造方法或者直接返回别的对象来作为本类 的实例。 
# 如果将类比喻为工厂，那么init()方法则是该工厂的生产工人，init()方法接受的初始化参 数则是生产所需原料，init()方法会按照方法中的语句负责将原料加工成实例以供工厂出货。而 new()则是生产部经理，new()方法可以决定是否将原料提供给该生产部工人，同时它还决定着出 货产品是否为该生产部的产品，因为这名经理可以借该工厂的名义向客户出售完全不是该工厂的产品。 
# new()方法的特性： 
# new()方法是在类准备将自身实例化时调用。 
# new()方法始终都是类的静态方法，即使没有被加上静态方法装饰器。

# 如果要得到当前类的实例，应当在当前类中的new()方法语句中调用当前类的父类 的new()方法。
# 所以通常这么写：
# class PositiveInteger(int):

  # def __new__(cls, value):

    # return super(PositiveInteger, cls).__new__(cls, abs(value))


# another example, like factory
# class Shape(object):
    # def __new__(cls, desc):
        # if cls is Shape:
            # if desc == 'big':   return super(Shape, cls).__new__(Rectangle)
            # if desc == 'small': return super(Shape, cls).__new__(Triangle)
        # else:
            # return super(Shape, cls).__new__(cls, desc)
			
	# def __init__(self, desc):
        # print "init called"
        # self.desc = desc
			
			
# class Triangle(Shape):
    # @property
    # def number_of_edges(self): return 3

# class Rectangle(Shape):
    # @property
    # def number_of_edges(self): return 4
