#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/1/25 16:40
# @Author  : Yajun Yin
# @Note    :

from collections import Iterable

"""
Target: get a flattened list from nested list
Method: using flatten(nested) to get a generator contains the flattened list of `nested`
Note: flatten(nested): generator      # implies that result is type of generator

We can get recursion
flatten(node) = F(flatten(node.child_1), flatten(node.child_2), ...,  flatten(node.child_N))

Then, what is F?

We see the recursive case first:
LHS, flatten(node) should make a generator of the flattened `node`
RHS, `node` has many child(e.g. child_N). Assume that flatten(any_node) works and then,
we can get many generators(e.g. n-th generator is result of flatten(node.child_N)).
So, F should make a new generator that contains each element in these generators from child node.

The core of recursion F can be code as:
>>> for child_nested in nested:               # traverse all child
>>>     for element in flatten(child_nested): # traverse element in a child generator  
>>>         yield element

It's time to talk about base case(terminating case): 
When the child_nested is a single element which is not iterable, it refers to the base case(leaf node).
Now, for ... in ... clause will raise an exception because child_nested is not iterable.
flatten(leaf_node) should make a generator of only the single element.
so, it can be coded as：
>>> from collections import Iterable
>>> if not isinstance(nested, Iterable):
>>>     yield nested

In this way, we finished coding the recursive function flatten().
>>> def flatten(nested):
>>>     if not isinstance(nested, Iterable):
>>>         yield nested
>>>     else:
>>>         for sublist in nested:
>>>             for element in flatten(sublist):
>>>                 yield element

My mistakes:
1. omitted `else`: yield will continue to execute, so else can not be omitted.
2. base case should return a generator(instead of `return`), so the recursive works. A generator makes another generator.

"""


# def f(state_n):
#     if state_n == state_base:
#         ret = base_do(state_base)
#         return ret
#
#     else:
#         # current state may be related to several previous state
#         ret = recursive_do(f(state_{n-1}), ..., f(state_{n-k}))
#         return ret

# 1. Define the state, state is the input of recursive function.
# 2. Determine state_base, use base_do() to get the result of state_base.
#       Return type also should be determined, denoted as f(state_base): type
# 3. Define the recursive_do, ret: type = recursive_do(input1: type, ..., input_k: type),
#       input_k:= f(state_{previous_k}). Note in recursive_do, input type corresponds to output type

def flatten1(nested):
    """函数的作用：生成将nested完全展开的generator"""
    if not isinstance(nested, Iterable):
        yield nested
    else:
        for sublist in nested:
            for element in flatten1(sublist):
                yield element


def flatten2(nested):
    try:
        for sublist in nested:
            for element in flatten2(sublist):
                yield element
    except TypeError:
        yield nested


def flatten3(nested):
    try:
        for sublist in nested:
            flatten3(sublist)
    except TypeError:
        print(nested)


def flatten4(nested):
    """print element in nested"""
    if not isinstance(nested, Iterable):
        print(nested)
    else:
        for sublist in nested:
            flatten4(sublist)


# def flatten(nested):
#     if not isinstance(nested, Iterable):
#         return [nested]
#     else:
#         ret = []
#         for sublist in nested:
#             for element in flatten(sublist):
#                 ret.append(element)
#         return ret


def flatten(nested):
    try:
        ret = []
        for sublist in nested:
            ret.extend(flatten(sublist))
        return ret
    except TypeError:
        return [nested]


a = (4, 5, (7, 8))
ret = flatten(a)
print(ret)
