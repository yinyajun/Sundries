#include <iostream>
#include "SortTestHelper.h"
#include "InsertionSort.h"
using namespace std;

template <typename T>
int __partition2(T arr[], int l, int r)
{
    swap(arr[l], arr[rand() % (r - l + 1) + l]);
    T v = arr[l]; // 取第一个元素作为基准值

    // arr[l+1...i) <=v; arr(j...r]>=v
    int i = l, j = r + 1;
    while (true)
    {
        //找到左边>=v,右边<=v的,swap
        while (arr[++i] < v && i <= r)
            ;
        while (arr[--j] > v && j >= l + 1)
            ;
        if (i >= j)
            break;
        swap(arr[i], arr[j]);
    }
    swap(arr[l], arr[j]);
    return j;
}

// 对arr[l..r]部分进行partition操作
// 返回p， 使得arr[l..p-1]< arr[p]; arr[p+1...r]>arr[p]
template <typename T>
int __partition1(T arr[], int l, int r)
{
    swap(arr[l], arr[rand() % (r - l + 1) + l]);

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
void __quickSort1(T arr[], int l, int r)
{
    // if (l >= r)
    //     return;
    if (r <= l + 15)
    {
        InsertionSort(arr, l, r);
        return;
    }

    int p = __partition1(arr, l, r);
    __quickSort1(arr, l, p - 1);
    __quickSort1(arr, p + 1, r);
}

template <typename T>
void __quickSort2(T arr[], int l, int r)
{
    // if (l >= r)
    //     return;
    if (r <= l + 15)
    {
        InsertionSort(arr, l, r);
        return;
    }

    int p = __partition2(arr, l, r);
    __quickSort2(arr, l, p - 1);
    __quickSort2(arr, p + 1, r);
}

// 三路快速排序处理 arr[l...r]
// 将arr[l..r]分为<v; ==v; >v三部分
// 之后递归对<v ; >v 两部分进行三路快速排序
template <typename T>
void __quickSort3ways(T arr[], int l, int r)
{
    if (r <= l + 15)
    {
        InsertionSort(arr, l, r);
        return;
    }

    // partition
    swap(arr[l], arr[rand() % (r - l + 1) + l]);
    T v = arr[l]; // 取第一个元素作为基准值

    int lt = l;     // arr[l+1...lt] < v
    int gt = r + 1; // arr[gt...r] > v
    int i = l + 1;  // arr[lt+1...i) ==v

    while (i < gt)
    {
        if (arr[i] < v)
        {
            swap(arr[lt + 1], arr[i]);
            lt++;
            i++;
        }
        else if (arr[i] > v)
        {
            swap(arr[gt - 1], arr[i]);
            gt--; //这里i位置有新的元素，所以i不要动
        }
        else
            i++;
    }
    swap(arr[l], arr[lt]); //此时 arr[l...lt-1]<v, arr[lt] == v
    __quickSort3ways(arr, l, lt - 1);
    __quickSort3ways(arr, gt, r);
}

template <typename T>
void quickSort1(T arr[], int n)
{
    srand(time(NULL));
    __quickSort1(arr, 0, n - 1);
}

template <typename T>
void quickSort2(T arr[], int n)
{
    srand(time(NULL));
    __quickSort2(arr, 0, n - 1);
}

template <typename T>
void quickSort3ways(T arr[], int n)
{
    srand(time(NULL));
    __quickSort3ways(arr, 0, n - 1);
}
