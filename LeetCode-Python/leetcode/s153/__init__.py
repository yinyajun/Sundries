from typing import List


# 有可能旋转，也有可能没旋转
# 判断单调区间，判断左侧，就容易将旋转区间推向右侧；反之，就容易将旋转区间推向左侧
# 对于未旋转情况，推向右侧可能越界，所以采取推向左侧

# 寻找旋转点

class Solution:

    def findMin(self, nums: List[int]) -> int:
        lo, hi = 0, len(nums)
        # [lo, hi)

        while lo < hi:
            mid = (lo + hi) // 2
            print(lo, hi, mid, nums[lo:hi])
            # if nums[mid] <= nums[hi- 1]:  # right mono
            # 不该用动态下标，应该用固定下标
            # 这里也不用认为是right mono，而是
            # 旋转点右边的<=a[-1]，旋转点左边的>=a[-1]
            # 相对于 nums[-1]，数组是分成两块的
            if nums[mid] <= nums[- 1]:  # right of pivot -> left region
                hi = mid
            else:
                lo = mid + 1

            # # if nums[mid] >= nums[lo]:  # left 单调，不该将lo作为基准，容易错位，与固定端点比
            # if nums[mid] >= nums[0]:  # left 单调
            #     lo = mid + 1  # [mid+1, hi)
            # else:  # right 单调
            #     hi = mid  # [lo, mid)

        # lo == hi
        return nums[lo]


print(Solution().findMin([3, 1, 2]))
