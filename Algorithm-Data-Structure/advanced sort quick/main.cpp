#include <iostream>
#include <string>
#include <stdlib.h>
#include "SortTestHelper.h"
#include "QuickSort.h"
#include "MergeSort.h"
#include "QuickSortAdvanced.h"

using namespace std;

main(int argc, char const *argv[])
{
    int n = 60000;

    cout << "Test for Random Array, size = " << n << ", random range [0, " << n << "]" << endl;
    int *arr1 = SortTestHelper::generateRandomArray(n, 0, n);
    int *arr2 = SortTestHelper::copyIntArray(arr1, n);
    int *arr3 = SortTestHelper::copyIntArray(arr1, n);

    SortTestHelper::testSort("Merge Sort", MergeSort, arr1, n);
    SortTestHelper::testSort("Quick Sort", quickSort, arr2, n);
    SortTestHelper::testSort("Quick Sort Advanced", quickSort1, arr3, n);

    delete[] arr1; // new的arr数组需要手动释放内存
    delete[] arr2;
    delete[] arr3;

    cout<<endl;
    
    int swapTimes = 100;
    cout<<"Test for Random Nearly Ordered Array, size = "<<n<<", swap time = "<<swapTimes<<endl;
    arr1 = SortTestHelper::generateNearlyOrderedArray(n,swapTimes);
    arr2 = SortTestHelper::copyIntArray(arr1, n);
    arr3 = SortTestHelper::copyIntArray(arr1, n);

    SortTestHelper::testSort("Merge Sort", MergeSort, arr1, n);
    SortTestHelper::testSort("Quick Sort", quickSort, arr2, n);
    SortTestHelper::testSort("Quick Sort Advanced", quickSort1, arr3, n);

    delete[] arr1; // new的arr数组需要手动释放内存
    delete[] arr2;
    delete[] arr3;

    cout<<endl;

    cout << "Test for Random Array, size = " << n << ", random range [0, " << n << "]" << endl;
    arr1 = SortTestHelper::generateRandomArray(n, 0, 10);
    arr2 = SortTestHelper::copyIntArray(arr1, n);
    arr3 = SortTestHelper::copyIntArray(arr1, n);
    int *arr4 = SortTestHelper::copyIntArray(arr1, n);
    int *arr5 = SortTestHelper::copyIntArray(arr1, n);

    SortTestHelper::testSort("Merge Sort", MergeSort, arr1, n);
    SortTestHelper::testSort("Quick Sort", quickSort, arr2, n);
    SortTestHelper::testSort("Quick Sort Advanced", quickSort1, arr3, n);
    SortTestHelper::testSort("Quick Sort Advanced2", quickSort2, arr4, n);
    SortTestHelper::testSort("Quick Sort 3ways", quickSort3ways, arr5, n);

    delete[] arr1; // new的arr数组需要手动释放内存
    delete[] arr2;
    delete[] arr3;
    delete[] arr4;
    delete[] arr5;

    system("pause");
    return 0;
}
