from collections import defaultdict


class TimeMap:
    def __init__(self):
        self.map = defaultdict(list)

    def set(self, key: str, value: str, timestamp: int) -> None:
        self.map[key].append((timestamp, value))

    def get(self, key: str, timestamp: int) -> str:
        values = self.map[key]
        if len(values) == 0:
            return ""

        # 最后一个小于等于tgt的索引
        # 相当于 第一个大于tgt的索引 - 1

        left = 0
        right = len(values)

        # [left, right)
        while left < right:
            mid = (left + right) // 2

            if values[mid][0] > timestamp:
                right = mid
            else:  # values[mid][0] <= timestamp
                left = mid + 1

        # left == right
        idx = left - 1
        if idx >= 0:
            return values[idx][1]
        return ""

# Your TimeMap object will be instantiated and called as such:
# obj = TimeMap()
# obj.set(key,value,timestamp)
# param_2 = obj.get(key,timestamp)


# v[t] <= tgt
# 最后一个值小于等于tgt

# key: [t1: val , t2: val]
