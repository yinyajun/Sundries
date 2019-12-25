#include <iostream>
#include <string>
#include <stdlib.h>
#include "SortTestHelper.h"

using namespace std;

template <typename T>
void BubbleSort(T arr[], int n)
{
    for (int i = 0; i < n - 1; i++)
    {
        // large element move to the right
        for (int j = 0; j < n - 1 - i; j++)
        {
            if (arr[j] > arr[j + 1])
                swap(arr[j], arr[j + 1]);
        }
    }
}

template <typename T>
void BubbleSort2(T arr[], int n)
{
    for (int i = 0; i < n - 1; i++)
    {
        // small element move to the left
        for (int j = n - 1; j > i; j--)
        {
            if (arr[j] < arr[j - 1])
                swap(arr[j], arr[j - 1]);
        }
    }
}

template <typename T>
void BubbleSortAdvance(T arr[], int n)
{
    bool swapped = true;
    do
    {
        swapped = false;
        for (int j = 0; j < n - 1; j++)
        {
            if (arr[j] > arr[j + 1])
            {
                swap(arr[j], arr[j + 1]);
                swapped = true;
            }
        }
        n--;
    } while (swapped);
}

template <typename T>
void BubbleSortAdvance2(T arr[], int n)
{
    bool swapped = true;
    do
    {
        int lastSwap = 0;
        swapped = false;
        for (int j = 0; j < n - 1; j++)
        {
            if (arr[j] > arr[j + 1])
            {
                swap(arr[j], arr[j + 1]);
                swapped = true;
                lastSwap = j + 1;
            }
        }
        n = lastSwap;
    } while(swapped);
}

template <typename T>
void CockTailSort(T arr[], int n)
{
    int left = 0;
    int right = n - 1;
    while (left < right)
    {
        // large element move to the right
        for (int j = left; j < right; j++)
        {
            if (arr[j] > arr[j + 1])
                swap(arr[j], arr[j + 1]);
        }
        right--;

        // small element move to the left
        for (int j = right; j > left; j--)
        {
            if (arr[j] < arr[j - 1])
                swap(arr[j], arr[j - 1]);
        }
        left++;
    }
}
