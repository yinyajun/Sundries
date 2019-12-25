#include <iostream>
#include "SortTestHelper.h"

using namespace std;

// 对arr[l..r]部分进行partition操作
// 返回p， 使得arr[l..p-1]< arr[p]; arr[p+1...r]>arr[p]
template <typename T>
int __partition(T arr[], int l, int r)
{

    T v = arr[l]; // 取第一个元素作为基准值
    // arr[l+1...j]<v ; arr[j+1...i)> v
    // i是当前考察的位置
    int j = l; //相当于左边的区间是空的
    //右边区间一开始也是空的
    for (int i = l + 1; i <= r; i++)
    {
        if (arr[i] < v)
        {
            swap(arr[j + 1], arr[i]);
            j++;
        }
    }
    swap(arr[l], arr[j]);
    return j;
}

// swap(arr[ ++j], arr[i]); 更加优雅

// 对arr[l..r]部分进行快速排序
template <typename T>
void __quickSort(T arr[], int l, int r)
{
    if (l >= r)
        return;

    int p = __partition(arr, l, r);
    __quickSort(arr, l, p - 1);
    __quickSort(arr, p + 1, r);
}

template <typename T>
void quickSort(T arr[], int n)
{
    __quickSort(arr, 0, n - 1);
}
