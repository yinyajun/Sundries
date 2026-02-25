# 将数组分为两部分，所有划分策略中，abs(左边最大值 - 右边最大值)的最大值是多少？
import sys
from collections import deque


class Solution:

    # 暴力解法： 遍历所有的划分方式，在每一种划分方式中，找到左右区间的最值

    def max_abs(self, nums: list[int]) -> int:
        # left: [0, i], right: [i+1, n-1]    i+1 <= n-1 -----> i <= n-2
        # left: [0, i], right: (i, n-1]      i < n-1
        n = len(nums)
        ans = 0
        for i in range(0, n - 1):
            l_max, r_max = -sys.maxsize, -sys.maxsize

            for j in range(0, i + 1):
                l_max = max(l_max, nums[j])

            for j in range(i + 1, n):
                r_max = max(r_max, nums[j])

            ans = max(ans, abs(r_max - l_max))

        return ans


class Solution2:
    # 暴力解法中，寻找区间最值花费了太多的时间，如果能O(1)时间内找到区间最值，那么整个算法也会降低为O(N)算法
    # 单调队列是维护区间上最值的一个好的数据结构，其中维护了很多备选答案

    def max_abs2(self, nums: list[int]) -> int:
        # left: [0, i], right: [i + 1, n - 1]
        # 左区间是不断扩张的，lmax求起来很简单
        # 右区间是动态变化的，所以用mono queue来存储备选答案

        n = len(nums)
        ans = 0
        l_max, r_max = -sys.maxsize, -sys.maxsize
        q = deque()

        # i=0， right [i+1, n-1]区间上的mono queue
        # 之后在不断收缩right区间
        for i in range(1, n):
            while q and nums[q[-1]] <= nums[i]:
                q.pop()
            q.append(i)

        for i in range(0, n - 1):
            # left: [0, i], right: [i + 1, n - 1]
            l_max = max(l_max, nums[i])

            if q and q[0] <= i:
                q.popleft()

            r_max = nums[q[0]]
            ans = max(ans, abs(r_max - l_max))

        return ans



    def max_abs(self, nums: list[int]) -> int:
        # left: [0, i], right: [i + 1, n - 1]
        # 左区间是不断扩张的，lmax求起来很简单
        # 右区间是动态变化的，所以用mono queue来存储备选答案

        n = len(nums)
        ans = 0
        l_max, r_max = -sys.maxsize, -sys.maxsize
        q = deque()

        # i=0， right [i+1, n-1]区间上的mono queue
        # 之后在不断收缩right区间
        for i in range(1, n):
            while q and nums[q[-1]] <= nums[i]:
                q.pop()
            q.append(i)

        for i in range(0, n - 1):
            # left: [0, i], right: [i + 1, n - 1]
            l_max = max(l_max, nums[i])
            r_max = nums[q[0]]
            ans = max(ans, abs(r_max - l_max))

            # for next right [i+2, n-1]
            if q and q[0] <= i+1:
                q.popleft()

        return ans
