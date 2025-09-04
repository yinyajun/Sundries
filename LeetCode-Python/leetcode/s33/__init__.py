from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        pass

        # [lo, hi)
        # a[mid] > tgt
        #     1. a[mid] > a[0], [lo, mid] inc -> left & right
        #           a. tgt > a[0] --> left
        #           b. tgt < a[0] --> right
        #     2. a[mid] < a[0], [mid, hi] inc -> left
        # a[mid] < tgt
        #     1. a[mid] > a[0], [lo, mid] inc -> right
        #     2. a[mid] < a[0], [mid, hi] inc -> left & right
        #           a. tgt < a[n-1] --> right
        #           b. tgt > a[n-1] --> left
        # -----------------------------------------------------------

        # 1. a[mid] >= a[lo], left inc
        #      a. a[lo] <= tgt < a[mid] -> left
        #      b. else: -> right
        # 2. a[mid] <= a[hi-1], right inc
        #      a. a[mid] < tgt <= a[hi-1] -> right
        #      b. else -> left
        # -----------------------------------------------------------------

        lo = 0
        hi = len(nums)
        # [0, lo) | [lo, hi) | [hi, n)

        while lo < hi:
            mid = (lo + hi) // 2
            if nums[mid] == target:
                return mid

            elif nums[mid] >= nums[lo]:  # left 单调
                if nums[lo] <= target < nums[mid]:
                    hi = mid  # [lo, mid)
                else:
                    lo = mid + 1  # [mid+1, hi)

            else:  # right 单调
                if nums[mid] < target <= nums[hi-1]:
                    lo = mid + 1 # [mid+1, hi)
                else:
                    hi = mid

        # lo == hi
        return -1
