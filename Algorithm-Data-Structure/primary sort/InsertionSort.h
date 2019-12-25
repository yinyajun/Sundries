#include <iostream>
#include <string>
#include <stdlib.h>
#include "SortTestHelper.h"

using namespace std;

template <typename T>
void InsertionSort(T arr[], int n)
{

    for (int i = 1; i < n; i++)
    {
        // insert arr[i] into a[i-1],a[i-2],...
        for (int j = i; j > 0 && arr[j - 1] > arr[j]; j--)
        {
            swap(arr[j], arr[j - 1]);
        }
    }
}

template <typename T>
void InsertionSortAdvanced(T arr[], int n)
{
    for (int i = 1; i < n; i++)
    {
        T e = arr[i];
        int j; //变量会在循环语句结束后销毁，必须将变量的作用域扩大。
        for (j = i ; j > 0 && arr[j-1] > e; j--)
            arr[j] = arr[j-1];
        arr[j] = e;
    }
}
