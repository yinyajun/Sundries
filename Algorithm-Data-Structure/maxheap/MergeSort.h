#include <iostream>
#include "SortTestHelper.h"

using namespace std;

//将arr[l..mid]和arr[mid+1..r]两部分合并
template<typename T>
void __merge(T arr[], int l, int mid, int r){
    T aux[r - l + 1]; // len([l,r]) = r - l + 1
    for (int i = l; i <= r; i++)
        aux[i - l] = arr[i]; // 有l的位置偏移

    int i = l, j = mid + 1; // 设置索引i,j,k
    for (int k = l; k <= r; k++){
        // 要用数组的索引，必须先保证索引不越界
        if (i > mid){ arr[k] = aux[j - l];j++;}
        else if (j > r) {arr[k] = aux[i - l];i++;}
        else if (aux[i - l] < aux[j - l]){arr[k] = aux[i - l];i++;}
        else {arr[k] = aux[j - l];j++;}
    }
}

// 递归使用归并，对[l..r]的范围进行排序
template<typename T>
void Mergesort(T arr[], int l, int r){
    if (l >= r)
        return;

    int mid = l + (r - l) / 2;
    Mergesort(arr, l, mid);
    Mergesort(arr, mid + 1, r);
    __merge(arr, l, mid, r);
}

template<typename T>
static void MergeSort(T arr[], int n){
    Mergesort(arr, 0, n - 1);
}