# 柱形图中最大矩形
# 矩形=长*宽
import sys


# 暴力解法：枚举长宽

class Solution:
    # 暴力解法，遍历所有可能的宽度，在每种宽度中寻找最小高度
    # 所有可能的宽度：[left, right]  width = right - left + 1
    # 这样的区间个数就有 O(N^2)个

    def max_rect_hist(self, heights: list[int]) -> int:
        n = len(heights)
        ans = 0
        for left in range(0, n):
            min_height = sys.maxsize
            for right in range(left, n):
                # [left, right]
                width = right - left + 1
                if heights[right] < min_height:
                    min_height = heights[right]

                area = width * min_height
                if area > ans:
                    ans = area

        return ans


class Solution2:
    # 暴力解法： 遍历所有的高度，然后从这个高度分别向左右延伸（直到遇到比当前高度小的，就确定了左右边界）

    def max_rect_hist(self, heights: list[int]) -> int:
        ans = 0

        for i in range(0, len(heights)):

            # find left boundry (find first `left` that heights[left] < heights[i]， [left+1, i] >= heights[i])
            # ==== for left = i; left >0 && heights[left] >= heights[i]; left -- {}
            left = i
            while left >= 0 and heights[i] <= heights[left]:
                left -= 1
            # ---> heights[i] > heights[left]

            # find right boundry
            # === for right = i; right < n && heights[right]>= heights[i]; right ++ {}
            right = i
            while right < len(heights) and heights[i] <= heights[right]:
                right += 1

            # region: [left + 1, right -1]
            width = right - left - 1
            area = heights[i] * width

            if area > ans:
                ans = area

        return ans


class Solution3:

    # 在暴力2的方法上，如果能用O(1)的时间来确定左右边界，那么整体的时间可以就是O(N)
    # 核心逻辑肯定是空间换时间，怎么用空间换时间
    # 问题：寻找最近的更小元素
    # 思路：遍历每个元素的时候，只维护一个最近的更小元素是不够的。
    # 对于一个新的元素，从原有的最近更小元素，可能无法得到当前元素的最近更小元素，从而需要扫一遍
    # 什么时候会有这种现象？比如当前元素<之前维护的最近最小元素，那么之前维护的最近最小元素就作废了，需要扫一遍才知道当前元素的最近最小元素
    # 如果存有这些备选答案（次近的更小元素，次次近的更更小元素等等，这些备选答案可以在最近最小元素作废的时候，顺利转正）
    # 怎么维护备胎？（移除哪些不可能的答案，遍历过程中，i<j && height[i] >= height[j], heights[i]不可能是备胎）
    # 那么此时发现备选答案满足单调性质：height[j0] < height[j1] < ...
    # 而最近加入数据结构中的备选答案就是当前答案，那么这样LIFO性质，使得需要栈来存储

    # 有个细节问题：这里的单调栈中需要严格单调吗？需要！如果没有严格单调，就会出现相同元素，
    # 那么栈顶可能就是和当前元素相等的元素，不是正确的左右边界 -> 就不是最近的更小元素了

    def max_rect_hist(self, heights: list[int]) -> int:
        ans = 0



        stack = []

        # 需要遍历一次获取
        for i in range(0, len(heights)):

            # find left boundry
            while len(stack) > 0 and heights[stack[-1]] >= heights[i]:
                stack.pop(-1)

            # len(left_stack) ==0  or heights[left_stack[-1]] < heights[i]
            if len(stack) == 0 :
                left = -1
            else:
                left = stack[-1]

            stack.append(i)





if __name__ == '__main__':
    heights = [2, 1, 5, 6, 2, 3]
    print(Solution().max_rect_hist(heights))
    print(Solution2().max_rect_hist(heights))
