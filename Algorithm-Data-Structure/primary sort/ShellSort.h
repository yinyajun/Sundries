#include <iostream>
#include <string>
#include <stdlib.h>
#include "SortTestHelper.h"

using namespace std;

template <typename T>
void ShellSort(T arr[], int n)
{
    int h = 1;
    while (h < n / 3)
        h = 3 * h + 1;
    while (h >= 1)
    {
        // 将数组变成h有序
        for (int i = h; i < n; i++)
        {
            // insert arr[i] into arr[i-h], arr[i-2h],...
            for (int j = i; j >= h && arr[j - h] > arr[j]; j -= h)
            {
                swap(arr[j - h], arr[j]);
            }
        }
        h = h / 3;
    }
}