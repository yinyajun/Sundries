#include <iostream>
#include "SortTestHelper.h"
#include"InsertionSort.h"

using namespace std;

//将arr[l..mid]和arr[mid+1..r]两部分合并
template<typename T>
void __merge1(T arr[], int l, int mid, int r){
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
void Mergesort3(T arr[], int l, int r){
    if(r-l<=15){ // 优化2：小数组用insertion sort
        InsertionSort(arr, l, r);
        return;
    }

    int mid = l + (r - l) / 2;
    Mergesort3(arr, l, mid);
    Mergesort3(arr, mid + 1, r);
    if (arr[mid]>arr[mid +1]) // 优化1：已经有序数组不用merge
        __merge1(arr, l, mid, r);
}

template<typename T>
static void MergeSort_Opt(T arr[], int n){
    Mergesort3(arr, 0, n - 1);
}