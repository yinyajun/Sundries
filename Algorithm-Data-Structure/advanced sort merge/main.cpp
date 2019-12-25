#include <iostream>
#include <string>
#include <stdlib.h>
#include "SortTestHelper.h"
#include "MergeSort_Opt.h"
#include "MergeSortBU.h"

using namespace std;

main(int argc, char const *argv[])
{
    int n = 200000;

    cout << "Test for Random Array, size = " << n << ", random range [0, " << n << "]" << endl;
    int *arr1 = SortTestHelper::generateRandomArray(n, 0, n);
    int *arr2 = SortTestHelper::copyIntArray(arr1, n);

    SortTestHelper::testSort("Merge Sort_Opt", MergeSort_Opt, arr1, n);
    SortTestHelper::testSort("Merge Sort_BU", MergeSort_BU, arr2, n);

    delete[] arr1; // new的arr数组需要手动释放内存
    delete[] arr2;

    cout<<endl;
    
    int swapTimes = 1000;
    cout<<"Test for Random Nearly Ordered Array, size = "<<n<<", swap time = "<<swapTimes<<endl;
    arr1 = SortTestHelper::generateNearlyOrderedArray(n,swapTimes);
    arr2 = SortTestHelper::copyIntArray(arr1, n);

    SortTestHelper::testSort("Merge Sort_Opt", MergeSort_Opt, arr1, n);
    SortTestHelper::testSort("Merge Sort_BU", MergeSort_BU, arr2, n);

    delete(arr1);
    delete(arr2);


    system("pause");
    return 0;
}
