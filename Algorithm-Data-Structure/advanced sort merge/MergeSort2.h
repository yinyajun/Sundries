#include <iostream>
#include "SortTestHelper.h"

using namespace std;
//习题2.2.9，将aux声明为MergeSort的局部变量。

//将arr[l..mid]和arr[mid+1..r]两部分合并
template<typename T>
void __merge(T arr[], T aux[], int l, int mid, int r){
    for (int i = l; i <= r; i++)
        aux[i - l] = arr[i]; // 有l的位置偏移

    int i = l, j = mid + 1; 
    for (int k = l; k <= r; k++){
        if (i > mid){ arr[k] = aux[j - l];j++;}
        else if (j > r) {arr[k] = aux[i - l];i++;}
        else if (aux[i - l] < aux[j - l]){arr[k] = aux[i - l];i++;}
        else {arr[k] = aux[j - l];j++;}
    }
}

// 递归使用归并，对[l..r]的范围进行排序
template<typename T>
void Mergesort2(T arr[], int l, int r){
    if (l >= r)
        return;

    T aux[r - l +1];
    int mid = l + (r - l) / 2;
    Mergesort2(arr, l, mid);
    Mergesort2(arr, mid + 1, r);
    __merge(arr, aux, l, mid, r);
}

template<typename T>
static void MergeSort2(T arr[], int n){
    Mergesort2(arr, 0, n - 1);
}