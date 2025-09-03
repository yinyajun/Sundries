# 解法一（推荐）：半开区间不变式 `[lo, hi)`

**核心不变式**

- (A) `∀ k ∈ [0, lo): a[k] < target`  —— 左侧全都不满足
- (B) `∀ k ∈ [hi, n): a[k] >= target` —— 右侧全都满足
- 未决区间始终是 `[lo, hi)`；答案一定在这个区间里。

**初始化**

```
lo = 0
hi = n            // 注意 hi 取到 n（半开）
```

**循环与收缩**

```
while lo < hi:
    mid = lo + (hi - lo) // 2
    if a[mid] >= target:
        hi = mid         # 保留 mid，答案在左半含 mid
    else:
        lo = mid + 1     # 丢弃 mid，答案在右半不含 mid
return lo                # 或 hi，二者相等
```

**正确性速证**

- 进入循环前，(A)(B) 显然成立（`[0,0)`和`[n,n)`为空）。
- 若 `a[mid] >= target`：把 `hi` 收到 `mid`，(B) 继续成立（因为 `[mid, n)` 仍然满足），(A) 不变。
- 若 `a[mid] < target`：把 `lo` 推到 `mid+1`，(A) 继续成立（`[0, mid+1)` 都 `< target`），(B) 不变。
- 每次都严格缩小区间长度，`lo==hi` 时，未决区间为空，根据不变式返回的就是最小满足下标；若完全不存在，最终会得到 `lo==n`。

**边界行为**

- 空数组：`n=0`，直接返回 `0`（即 `n`）。
- `target <= a[0]` → 返回 `0`；`target > a[n-1]` → 返回 `n`。
- 重复元素：由于遇到 `>=` 就把 `hi=mid`，最终停在**第一**个符合的位置。

# 解法二：闭区间 + 备用答案

同样正确，适合喜欢 `lo<=hi` 模式的人：

```
lo, hi = 0, n-1
ans = n
while lo <= hi:
    mid = lo + (hi - lo) // 2
    if a[mid] >= target:
        ans = mid
        hi = mid - 1
    else:
        lo = mid + 1
return ans
```

`ans` 记录目前最左的可行位置；找不到即保持为 `n`。

------

# 复杂度

时间 `O(log n)`，空间 `O(1)`。

------

# 例子快速校验

```
a = [1,2,2,4,7], n=5
```

- `target=2` → 返回 `1`（第一 个 2）
- `target=5` → 返回 `4`（第一个 ≥5 是 7 的下标 4）
- `target=0` → 返回 `0`
- `target=10` → 返回 `5`（即 n）

------

# 常见坑

- `hi` 初值：半开模板一定要用 `hi = n`，否则会错过答案= n 的情况。
- 循环条件：半开用 `while lo < hi`；闭区间用 `while lo <= hi`。
- `mid` 计算：用 `lo + (hi - lo)//2` 以免在 C++/Java 溢出（Python 不会）。
- 空数组与全不满足：返回 `n` 的约定要在上层调用里处理好（例如判断是否存在）。

------

# 代码片段

**Python（半开模板）**

```
def lower_bound(a, target):
    n = len(a)
    lo, hi = 0, n
    while lo < hi:
        mid = lo + (hi - lo) // 2
        if a[mid] >= target:
            hi = mid
        else:
            lo = mid + 1
    return lo  # 若为 n 表示不存在
```

**C++（STL 同名算法是 `std::lower_bound`）**

```
int lower_bound_idx(const vector<int>& a, int target) {
    int n = (int)a.size();
    int lo = 0, hi = n;            // [lo, hi)
    while (lo < hi) {
        int mid = lo + (hi - lo) / 2;
        if (a[mid] >= target) hi = mid;
        else lo = mid + 1;
    }
    return lo; // == n 表示不存在
}
```

------

# 小变体

- “第一个 **> target** 的下标”（upper_bound）：把条件改成 `if a[mid] > target: hi = mid else: lo = mid + 1`，返回 `lo`。
- “最后一个 `< target` 的下标”：可用 `upper_bound(target) - 1` 得到（注意越界判断）。