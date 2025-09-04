from typing import List


class Solution:
    def findMin(self, nums: List[int]) -> int:
        lo, hi = 0, len(nums)
        # [lo, hi) | [hi, n)

        while lo < hi:
            mid = (lo + hi) // 2
            print(lo, hi, mid, nums[lo:hi])

            if hi - lo == 1 and hi < len(nums):
                return min(nums[lo], nums[hi])

            # a[mid] > a[-1]  # right region
            # a[mid] < a[-1]  # left region
            # a[mid] == a[-1] # shrink

            if nums[mid] == nums[hi - 1]:
                hi -= 1
            elif nums[mid] > nums[hi - 1]:
                lo = mid + 1
            else:
                hi = mid + 1

        return nums[lo]


print(Solution().findMin([3, 3, 1, 3]))
