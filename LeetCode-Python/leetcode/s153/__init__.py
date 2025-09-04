from typing import List


class Solution:
    def findMin(self, nums: List[int]) -> int:
        # 寻找旋转点
        # 1. left 单调： a[mid] >= a[lo] -> right
        # 2. right 单调： a[mid] < a[lo] -> left

        lo, hi = 0, len(nums)

        # [lo, hi)

        while lo < hi:
            mid = (lo + hi) // 2

            # if nums[mid] >= nums[lo]:  # left 单调，不该将lo作为基准，容易错位，与固定端点比
            if nums[mid] >= nums[0]:  # left 单调
                lo = mid + 1  # [mid+1, hi)
            else: # right 单调
                hi = mid  # [lo, mid)

        # lo == hi
        return nums[lo]


Solution().findMin([11,13,15,17])
