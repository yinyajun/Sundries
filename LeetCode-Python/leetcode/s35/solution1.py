from typing import List


class Solution:
    # 寻找target的插入位置
    # nums[pos-1] < target <= nums[pos]
    # 在一个有序数组中找第一个大于等于 target 的下标
    # 转化为这么一个lower bound问题
    def searchInsert(self, nums: List[int], target: int) -> int:
        # return self.search1(nums, target, 0, len(nums) - 1)
        # return self.search2(nums, target, 0, len(nums))
        return self.search3(nums, target)

    # 闭区间非递归
    def search4(self, nums, target):
        left = 0
        right = len(nums) - 1

        # [left, right]未决
        ans = len(nums)
        while left <= right:
            mid = left + (right - left) // 2
            if nums[mid] >= target:
                ans = mid  # mid 满足条件，可能是答案
                right = mid - 1  # 继续向左收缩，看看有没有更小的
            else:
                left = mid + 1

        # left > right
        # left 可能越界
        return ans

    # 开区间非递归
    def search3(self, nums, target):
        left = 0
        right = len(nums)
        # [left, right) 未决区域
        while left < right:
            mid = left + (right - left) // 2
            if nums[mid] >= target:
                # mid 可能是答案，继续缩到左半区 [left, mid)
                right = mid
            else:
                # mid 不行，答案只能在右半区 [mid+1, hi)
                left = mid + 1

        # left >= right
        return left

    # 开区间写法：递归写法
    def search2(self, nums, tgt, left, right):
        # [0, left) | 未决区域 [left, right) | [right, n)
        if left >= right:
            # [0, left) | ∅ | [left, n)
            return left  # 未决区域为空，直接返回分界点。

        mid = left + (right - left) // 2

        if nums[mid] >= tgt:
            # mid 可能是答案，继续缩到左半区 [left, mid)
            return self.search2(nums, tgt, left, mid)
        else:
            # nums[mid] < tgt
            # mid 不行，答案只能在右半区 [mid+1, hi)
            return self.search2(nums, tgt, mid + 1, right)

    # 闭区间写法：递归写法
    def search1(self, nums, target, left, right):
        # 未决区域 [left, right]
        if left > right:
            return len(nums)  # 未决区域为空，此时还没有找到

        mid = left + (right - left) // 2

        if nums[mid] >= target:
            # mid是候选，继续左半边找
            ans = self.search1(nums, target, left, mid - 1)
            # 如果左半边没有，ans是哨兵值
            return min(ans, mid)
        else:
            # nums[mid] < target, mid太小，去右边找
            return self.search1(nums, target, mid + 1, right)
