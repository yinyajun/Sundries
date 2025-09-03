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
        #
        # 1. left:
        #   a[mid] > tgt > a[0]
        #   a[mid] > tgt && tgt < a[0]
