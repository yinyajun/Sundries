#include <iostream>
// #include "InsertionSort.h"
#include "SortTestHelper.h"

using namespace std;

//将arr[l..mid]和arr[mid+1..r]两部分合并
template<typename T>
void __merge2(T arr[], int l, int mid, int r){
    T aux[r - l + 1]; 
    for (int i = l; i <= r; i++)
        aux[i - l] = arr[i]; 

    int i = l, j = mid + 1; 
    for (int k = l; k <= r; k++){
        if (i > mid){ arr[k] = aux[j - l];j++;}
        else if (j > r) {arr[k] = aux[i - l];i++;}
        else if (aux[i - l] < aux[j - l]){arr[k] = aux[i - l];i++;}
        else {arr[k] = aux[j - l];j++;}
    }
}


template<typename T>
static void MergeSort_BU(T arr[], int n){
    for(int sz=1; sz<n; sz+=sz){ //sz子数组大小
        for(int i=0; i+sz<n; i+=2*sz){ // 每次对[i..i+sz-1][i+sz..i+sz+sz-1]做merge
            if(sz<=15)
                InsertionSort(arr, i, min(i+sz+sz-1, n-1));
            else{
                if(arr[i+sz-1]> arr[i+sz]) 
                    __merge2(arr, i, i+sz-1, min(i+sz+sz-1, n-1));
            }
        }
    }
}