#include <iostream>
#include <string>
#include <stdlib.h>
#include "SortTestHelper.h"
#include "SelectionSort.h"
#include "InsertionSort.h"
#include "BubbleSort.h"
#include "ShellSort.h"

using namespace std;

main(int argc, char const *argv[])
{
    int n = 40000;
    int *arr = SortTestHelper::generateNearlyOrderedArray(n, 1000);
    int *arr2 = SortTestHelper::copyIntArray(arr, n);
    int *arr3 = SortTestHelper::copyIntArray(arr, n);
    int *arr4 = SortTestHelper::copyIntArray(arr, n);
    // SortTestHelper::testSort("Insertion Sort", InsertionSort, arr, n);
    // SortTestHelper::testSort("Insertion Sort Advanced", InsertionSortAdvanced, arr2, n);
    // SortTestHelper::testSort("Selection Sort", selectionSort, arr3, n);
    SortTestHelper::testSort("Insert Sort", InsertionSort, arr, n);
    SortTestHelper::testSort("Shell Sort", ShellSort, arr2, n);
    

    // SortTestHelper::printArray(arr,n);
    // BubbleSortAdvance2(arr,n);
    // SortTestHelper::printArray(arr,n);

    

    delete[] arr; // new的arr数组需要手动释放内存
    delete[] arr2;
    delete[] arr3;
    delete[] arr4;

    system("pause");
    return 0;
}
