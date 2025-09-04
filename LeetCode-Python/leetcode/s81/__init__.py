from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> bool:
        # 1. a[mid] >= a[lo], left inc
        #      a. a[lo] <= tgt < a[mid] -> left
        #      b. else: -> right
        # 2. a[mid] <= a[hi-1], right inc
        #      a. a[mid] < tgt <= a[hi-1] -> right
        #      b. else -> left
        # -----------------------------------------------------------------
        # 但是这里有重复元素，有可能 a[lo] == a[mid] == a[hi-1]，这样无法区分那个区间单调
        # 那么无法二分，只能通过[lo+1, hi-1)方式来缩小区间
        lo = 0
        hi = len(nums)

        # [0, lo) | [lo, hi) | [hi, n)
        while lo < hi:
            mid = (lo + hi) // 2

            if nums[mid] == target:
                return True

            if nums[lo] == nums[mid] == nums[hi - 1]:
                lo += 1
                hi -= 1

            elif nums[mid] >= nums[lo]:  # left 单调
                if nums[lo] <= target < nums[mid]:  # target在left单调区间
                    hi = mid  # [lo, mid)
                else:
                    lo = mid + 1

            else:  # right 单调
                if nums[mid] < target <= nums[hi - 1]:  # target在right单调区间
                    lo = mid +1
                else:
                    hi = mid

        # lo == hi
        return False
