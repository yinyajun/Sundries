#include <iostream>
#include <string>
#include <stdlib.h>
#include "SortTestHelper.h"
#include "QuickSort.h"
#include "MergeSort.h"
#include "QuickSortAdvanced.h"
#include "HeapSort.h"

using namespace std;

// template<typename T>
// void __shiftDown2(T arr[], int n, int k){

//     T e = arr[k];
//     while( 2*k+1 < n ){
//         int j = 2*k+1;
//         if( j+1 < n && arr[j+1] > arr[j] )
//             j += 1;

//         if( e >= arr[j] ) break;


//         arr[k] = arr[j];
//         k = j;
//     }

//     // arr[k] = e;
// }

// template<typename T>
// void heapSort(T arr[], int n){

//     for( int i = (n-1)/2 ; i >= 0 ; i -- )
//         __shiftDown2(arr, n, i);

//     for( int i = n-1; i > 0 ; i-- ){
//         swap( arr[0] , arr[i] );
//         __shiftDown2(arr, i, 0);
//     }
// }



main(int argc, char const *argv[])
{
    int n = 500000;

    cout << "Test for Random Array, size = " << n << ", random range [0, " << n << "]" << endl;
    int *arr1 = SortTestHelper::generateRandomArray(n, 0, n);
    int *arr2 = SortTestHelper::copyIntArray(arr1, n);
    int *arr3 = SortTestHelper::copyIntArray(arr1, n);
    int *arr4 = SortTestHelper::copyIntArray(arr1, n);
    int *arr5 = SortTestHelper::copyIntArray(arr1, n);
    int *arr6 = SortTestHelper::copyIntArray(arr1, n);

    SortTestHelper::testSort("Merge Sort", MergeSort, arr1, n);
    SortTestHelper::testSort("Quick Sort", quickSort, arr2, n);
    SortTestHelper::testSort("Quick Sort Advanced", quickSort1, arr3, n);
    SortTestHelper::testSort("Heap Sort1", heapSort1, arr4, n);
    SortTestHelper::testSort("Heap Sort2", heapSort2, arr5, n);
    SortTestHelper::testSort("Heap Sort", heapSort, arr6, n);
    
    

    delete[] arr1; // new的arr数组需要手动释放内存
    delete[] arr2;
    delete[] arr3;
    delete[] arr4;
    delete[] arr5;
    delete[] arr6;

    cout<<endl;
    
    // int swapTimes = 100;
    // cout<<"Test for Random Nearly Ordered Array, size = "<<n<<", swap time = "<<swapTimes<<endl;
    // arr1 = SortTestHelper::generateNearlyOrderedArray(n,swapTimes);
    // arr2 = SortTestHelper::copyIntArray(arr1, n);
    // arr3 = SortTestHelper::copyIntArray(arr1, n);

    // SortTestHelper::testSort("Merge Sort", MergeSort, arr1, n);
    // SortTestHelper::testSort("Quick Sort", quickSort, arr2, n);
    // SortTestHelper::testSort("Quick Sort Advanced", quickSort1, arr3, n);

    // delete[] arr1; // new的arr数组需要手动释放内存
    // delete[] arr2;
    // delete[] arr3;

    // cout<<endl;

//     // cout << "Test for Random Array, size = " << n << ", random range [0, " << n << "]" << endl;
//     // arr1 = SortTestHelper::generateRandomArray(n, 0, 10);
//     // arr2 = SortTestHelper::copyIntArray(arr1, n);
//     // arr3 = SortTestHelper::copyIntArray(arr1, n);
//     // arr4 = SortTestHelper::copyIntArray(arr1, n);
//     // arr5 = SortTestHelper::copyIntArray(arr1, n);

//     // SortTestHelper::testSort("Merge Sort", MergeSort, arr1, n);
//     // SortTestHelper::testSort("Quick Sort", quickSort, arr2, n);
//     // SortTestHelper::testSort("Quick Sort Advanced", quickSort1, arr3, n);
//     // SortTestHelper::testSort("Quick Sort Advanced2", quickSort2, arr4, n);
//     // SortTestHelper::testSort("Quick Sort 3ways", quickSort3ways, arr5, n);

//     // delete[] arr1; // new的arr数组需要手动释放内存
//     // delete[] arr2;
//     // delete[] arr3;
//     // delete[] arr4;
//     // delete[] arr5;

    system("pause");
    return 0;
}
// using namespace std;






// int main() {

//     int n = 100000;

//     // 测试1 一般性测试
//     cout<<"Test for Random Array, size = "<<n<<", random range [0, "<<n<<"]"<<endl;
//     int* arr1 = SortTestHelper::generateRandomArray(n,0,n);
//     int* arr2 = SortTestHelper::copyIntArray(arr1, n);
//     int* arr3 = SortTestHelper::copyIntArray(arr1, n);
//     int* arr4 = SortTestHelper::copyIntArray(arr1, n);
//     int* arr5 = SortTestHelper::copyIntArray(arr1, n);
//     int* arr6 = SortTestHelper::copyIntArray(arr1, n);

//     SortTestHelper::testSort("Merge Sort", MergeSort, arr1, n);
//     SortTestHelper::testSort("Quick Sort", quickSort, arr2, n);
//     SortTestHelper::testSort("Quick Sort Advanced", quickSort1, arr3, n);
//     SortTestHelper::testSort("Heap Sort 1", heapSort1, arr4, n);
//     SortTestHelper::testSort("Heap Sort 2", heapSort2, arr5, n);
//     SortTestHelper::testSort("Heap Sort 3", heapSort, arr6, n);

//     delete[] arr1;
//     delete[] arr2;
//     delete[] arr3;
//     delete[] arr4;
//     delete[] arr5;
//     delete[] arr6;

//     cout<<endl;

//     system("pause");
//     return 0;
// }