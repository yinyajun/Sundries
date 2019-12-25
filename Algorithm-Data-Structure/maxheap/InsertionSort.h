#include <iostream>

using namespace std;

//原先是[0,n-1],现在是[l,r]
//[1,n)--> [l+1,r+1)
template <typename T>
void InsertionSort(T arr[], int l, int r)
{
    for (int i = l+ 1; i < r+1; i++)
    {
        T e = arr[i];
        int j; //变量会在循环语句结束后销毁，必须将变量的作用域扩大。
        // [0,i]--->[l,i]
        // a in (0,i], compare(a-1,a)
        // a in (l,i]
        for (j = i ; j > l && arr[j-1] > e; j--)
            arr[j] = arr[j-1];
        arr[j] = e;
    }
}
