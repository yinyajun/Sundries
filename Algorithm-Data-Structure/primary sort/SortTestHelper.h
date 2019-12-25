#ifndef _SORTTESTHELPER_H
#define _SORTTESTHELPER_H

#include <iostream>
#include <ctime>
#include <cassert>
using namespace std;

namespace SortTestHelper
{
// generate random array, array length is n, range from [left, right]
int *generateRandomArray(int n, int left, int right)
{

    assert(left <= right);

    int *arr = new int[n];
    srand(time(NULL));
    for (int i = 0; i < n; i++)
        arr[i] = rand() % (right - left + 1) + left;
    return arr;
}

int *generateNearlyOrderedArray(int n, int swapTimes)
{
    int *arr = new int[n];
    for (int i = 0; i < n; i++)
        arr[i] = i;
    srand(time(NULL));
    for (int i = 0; i < swapTimes; i++)
    {
        int x = rand() % n;
        int y = rand() % n;
        swap(arr[x], arr[y]);
    }
    return arr;
}

template <typename T>
void printArray(T arr[], int n)
{
    for (int i = 0; i < n; i++)
        cout << arr[i] << " ";
    cout << endl;
    return;
}

template <typename T>
bool isSorted(T arr[], int n)
{
    for (int i = 0; i < n - 1; i++)
        if (arr[i] > arr[i + 1])
            return false;
    return true;
}

template <typename T>
void testSort(string sortName, void (*sort)(T[], int n), T arr[], int n)
{
    clock_t start = clock();
    sort(arr, n);
    clock_t end = clock();
    assert(isSorted(arr, n));
    cout << sortName << " : " << double(end - start) / CLOCKS_PER_SEC << " s" << endl;
    return;
}

int *copyIntArray(int a[], int n)
{
    int *arr = new int[n];
    copy(a, a + n, arr);
    return arr;
}

} // namespace SortTestHelper
#endif