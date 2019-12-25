"""vars([object]) -> dictionary
Without arguments, equivalent to locals().
With an argument, equivalent to object.__dict__.

vars()返回命名空间或者作用域，字典里会有许多当前模块的魔术方法的键值对。
"""

"""
{'__name__': '__main__', 
'__doc__': 'vars([object]) -> dictionary\nWithout arguments, equivalent to locals().\nWith an argument, equivalent to object.__dict__.\n\nvars()返回命名空间或者作用域，字典里会有许多当前模块的魔术方法的键值对。\n', 
'__package__': None, 
'__loader__': <_frozen_importlib_external.SourceFileLoader object at 0x00000000025FB080>, 
'__spec__': None, '__annotations__': {}, '__builtins__': <module 'builtins' (built-in)>, 
'__file__': 'D:/work/git_repo/Details-In-Python/built_in func/bt_vars().py', 
'__cached__': None}
"""

"""几个相似的函数： local"""
from cmd import Cmd
import os
import sys

# class Cli(Cmd):
#     prompt = 'spark>'
#     intro = 'Welcome to spark!'
#
#     def __init(self):
#         Cmd.__init__(self)
#
#     def do_hello(self, line):
#         print("hello", line)
#
#
# if __name__ == '__main__':
#     cli = Cli()
#     cli.cmdloop()
import timeit

foooo = """
sum = []
for i in range(1000):
    sum.append(i)
"""

print(timeit.timeit(stmt="[i for i in range(1000)]", number=100000))
print(timeit.timeit(stmt=foooo, number=100000))
