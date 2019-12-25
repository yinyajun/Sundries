class Singleton(object):
    _instance = None

    def __new__(cls, *args, **kwargs):
        if not cls._instance:
            cls._instance = super(Singleton, cls).__new__(cls)
        return cls._instance


class Dog(Singleton):
    def __init__(self, name):
        self.name = name


if __name__ == '__main__':
    a = Dog("Husky")
    print(id(a))
    print(a.name)

    b = Dog("p")
    print(id(b))
    print(b.name)

    print(a==b)
