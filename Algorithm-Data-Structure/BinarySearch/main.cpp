#include <iostream>
#include <cassert>
#include <ctime>
#include <math.h>

using namespace std;

template<typename T>
int binarySearch(T arr[], int n, T target) {

    // 在arr[l...r]中查找target
    int l = 0, r = n - 1;
    while (l <= r) {

        int mid = l + (r - l) / 2;

        if (arr[mid] == target)
            return mid;

        if (arr[mid] < target)
            // 在arr[mid+1...r]中查找target,mid已经不可能是要查的索引了
            l = mid + 1;
        else // arr[mid]>target
            //在arr[l...mid-1]中查找target,mid已经不可能是要查的索引了
            r = mid - 1;
    }
    return -1; //循环结束都没有找到target对应的索引
}

template<typename T>
int binarySearchRecursive(T arr[], int l, int r, T target) {
    // 在arr[l...r]中查找target
    if (l > r)
        return -1;

    int mid = l + (r - l) / 2;
    if (arr[mid] > target)
        // arr[l...mid-1]中查找target
        return binarySearchRecursive(arr, l, mid - 1, target);
    else if (arr[mid] < target)
        //在arr[mid+1...r]中查找target
        return binarySearchRecursive(arr, mid + 1, r, target);
    else
        return mid;
}


template<typename T>
int binarySearch2(T arr[], int n, T target) {
    return binarySearchRecursive(arr, 0, n - 1, target);
}


template<typename T>
int Rank(T arr[], int n, T target) {
    int l = 0, r = n - 1;
    // 小于指定键的数目
    // 指定键的值不重复，第一个大于等于该键的键的index
    // 指定键重复，等于该键的某个index
    while (l <= r) {
        int mid = l + (r - l) / 2;
        if (arr[mid] == target)
            return mid;
        if (arr[mid] < target)
            l = mid + 1;
        else // arr[mid]> target
            r = mid - 1;
    }
    return l;
}

template<typename T>
int Rank2(T arr[], int n, T target) {
    int l = 0, r = n - 1;
    // 小于指定键的数目
    // 指定键的值不重复，第一个大于等于该键的键的index
    // 指定键重复，等于该键的某个index
    while (l <= r) {
        int mid = l + (r - l) / 2;
        if (arr[mid] == target)
            return mid;
        if (arr[mid] < target)
            l = mid + 1;
        else // arr[mid]> target
            r = mid - 1;
    }
    return r;
}


template<typename T>
int Ceil(T arr[], int n, T target) {
    // 二分查找法, 在有序数组arr中, 查找target
    // 如果找到target, 返回最后一个target相应的索引index
    // 如果没有找到target, 返回比target大的最小值相应的索引, 如果这个最小值有多个, 返回最小的索引
    // 如果这个target比整个数组的最大元素值还要大, 则不存在这个target的ceil值, 返回整个数组元素个数n

    // 寻找比target大的最小索引值
    int l = 0, r = n;
    while (l < r) {
        // 使用普通的向下取整即可避免死循环
        int mid = l + (r - l) / 2;
        if (arr[mid] <= target)
            l = mid + 1;
        else // arr[mid] > target
            r = mid;
    }
    assert(l == r);
    // 如果该索引-1就是target本身, 该索引+1即为返回值
    if (r - 1 >= 0 && arr[r - 1] == target)
        return r - 1;
    // 否则, 该索引即为返回值
    return r;
}

template<typename T>
int floor2(T arr[], int n, T target) {
    int lo = 0, hi = n - 1;
    while (lo < hi) {
        int mid = (lo + hi) / 2;
        if (arr[mid] >= target) {
            hi = mid;
        } else {
            lo = mid + 1;
        }
    }
    return hi;
}

template<typename T>
int floor3(T arr[], int n, T target) {
    int lo = 0, hi = n - 1;
    while (lo < hi) {
        int mid = (lo + hi) / 2;
        if (arr[mid] >= target) {
            hi = mid - 1;
        } else {
            lo = mid + 1;
        }
    }
    return hi;
}

template<typename T>
int ceil2(T arr[], int n, T target) {
    int lo = 0, hi = n - 1;
    while (lo < hi) {
        int mid = (int) ceil((lo + hi) / 2.0);
        if (arr[mid] > target)
            hi = mid - 1;
        else
            lo = mid;
    }
    return lo;
}

//int main() {
//    int n = 1000000;
//    int *a = new int[n];
//    for (int i = 0; i < n; i++)
//        a[i] = i;
//
//    // 测试非递归二分查找法
//    clock_t startTime = clock();
//    for (int i = 0; i < 2 * n; i++) {
//        int v = binarySearch(a, n, i);
//        if (i < n)
//            assert(v == i);
//        else
//            assert(v == -1);
//    }
//    clock_t endTime = clock();
//    cout << "Binary Search (Without Recursion): " << double(endTime - startTime) / CLOCKS_PER_SEC << " s" << endl;
//
//    // 测试递归的二分查找法
//    startTime = clock();
//    for (int i = 0; i < 2 * n; i++) {
//        int v = binarySearch2(a, n, i);
//        if (i < n)
//            assert(v == i);
//        else
//            assert(v == -1);
//    }
//    endTime = clock();
//    cout << "Binary Search (Recursion): " << double(endTime - startTime) / CLOCKS_PER_SEC << " s" << endl;
//
//    delete[] a;
//    return 0;
//}

int main() {
    int array[10] = {1, 3, 6, 6, 6, 9, 9, 9, 12, 15};
    int index1 = binarySearch(array, 10, 8);
    int index2 = Rank(array, 10, 8);
    int index3 = floor2(array, 10, 5);
    int index4 = floor3(array, 10, 5);
//    cout << index1 << ":" << array[index1] << endl;
//    cout << index2 << ":" << array[index2] << endl;
    cout << index3 << ":" << array[index3] << endl;
    cout << index4 << ":" << array[index4] << endl;

    return 0;

}