from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        return f1(nums, target)


def f1(nums, tgt):
    left = 0
    right = len(nums)

    # [0, left) | [left, right) | [right, n)
    while left < right:
        mid = (left + right) // 2
        if nums[mid] == tgt:
            return mid

        elif nums[mid] > tgt:
            # right = mid - 1  # 缩到左半区 [lo, mid)
            right = mid  # 缩到左半区 [lo, mid)

        else:  # nums[mid] < tgt
            left = mid + 1  # 缩到右半区 [mid+1, hi)

    # left == right
    return -1
