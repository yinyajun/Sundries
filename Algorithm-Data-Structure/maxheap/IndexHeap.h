#include <iostream>
#include <string>
#include <stdlib.h>
#include <algorithm>
#include <ctime>
#include <cmath>
#include <cassert>

using namespace std;


template<typename Item>
class IndexMaxHeap{
private:
    Item *data;
    int *indexes;
    int* reverse; 
    
    int count;
    int capacity;

    // 这里的k是索引数组的位置，要拿到data，得data[indexes[k]]
    void shiftUp(int k){
        while( k > 1 && data[indexes[k/2] < data[indexes[k]] ){
            swap( indexes[k/2], indexes[k] );//真正交换的是索引
            reverse[indexes[k/2]] = k/2; //indexes[k] = indexes[k/2]
            reverse[index[k]]=k;
            k /= 2;
        }
    }

    void shiftDown(int k){
        while( 2*k <= count ){
            int j = 2*k;
            if( j+1 <= count && data[indexes[j+1] > data[indexes[j]] ) j ++;
            if( data[indexes[k] >= data[indexes[j] ) break;
            swap( indexes[k] , indexes[j] );
            reverse[indexes[k]] = k;
            reverse[indexes[k/2]] = k/2;
            k = j;
        }
    }

public:

    MaxHeap(int capacity){
        data = new Item[capacity+1];
        indexes = new int[capacity + 1];
        reverse = new int[capacity + 1];
        // reverse记录index i在indexes堆数组中的位置，全都初始化为0,index 0 没有使用
        for (int i=0; i<=capacity;i++)
            reverse[i]=0;

        count = 0;
        this->capacity = capacity;
    }

    MaxHeap(Item arr[], int n){
        data = new Item[n+1];
        capacity = n;

        for( int i = 0 ; i < n ; i ++ )
            data[i+1] = arr[i];
        count = n;

        for( int i = count/2 ; i >= 1 ; i -- )
            shiftDown(i);
    }

    ~MaxHeap(){
        delete[] data;
        delete[] indexes;
        delete[] reverse;
    }

    int size(){
        return count;
    }

    bool isEmpty(){
        return count == 0;
    }

    // 在heap内部索引是从1开始的，但是传入的i对用户而言是从0索引开始的
    void insert(int i, Item item){
        assert( count + 1 <= capacity );
        assert( i+1>=1 && i+1<=capacity);

        i+=1; // i变为从1开始的索引

        data[i] = item;
        indexes[count + 1] = i; // 在末尾添加索引，indexHeap添加索引
        reverse[i] = count+1;

        shiftUp(count+1);
        count ++;
    }

    Item extractMax(){
        assert( count > 0 );

        Item ret = data[indexes[1]];
        swap( indexes[1] , indexes[count] );

        reverse[indexes[1]] = 1;
        reverse[indexes[count]] = 0;//注意indexes[count]的值（索引）删除了


        count --;
        shiftDown(1);
        return ret;
    }

    bool contain(int i){
        assert(i+1>=1 && i+1 <=capacity);
        return reverse[i+1] !=0; // i+1是为了从1开始

    }

    void change( int i, Item newItem){

        assert( contain(i)); //确保索引i一定在索引堆中

        i+=1; //从1开始
        data[i]=newItem;

        // // 得到在堆中的位置后，shiftUp和shiftDown就好了
        // // 找到data[i]在堆中的位置，即find indexes[j]=i, j就是data[i]在索引堆中的位置
        // // 顺序查找的时间复杂度O(N)
        // for (int j =1 ; j<=count; j++){
        //     if (indexes[j]==i){
        //         shiftUp(j);
        //         shiftDown(j);
        //         return ;
        //     }
        // }

        int j = reverse[i]; // 通过反向查找，使得时间复杂度变为O(lgN)
        shiftUp(j);
        shiftDown(j);
    }
};
