//
// Created by Administrator on 2019/3/3.
//

#ifndef UNIONFIND_UNIONFIND3_H
#define UNIONFIND_UNIONFIND3_H

#include <iostream>
#include <cassert>

using namespace std;

namespace UF3 {
    class UnionFind{
    private:
        int* parent;
        int* sz; // sz[i]表示以i为根的集合中元素个数
        int count;
    public:
        UnionFind(int count){
            parent = new int[count];
            sz = new int[count];
            this->count = count;
            for( int i=0; i<count;i++){
                parent[i]= i;
                sz[i] = 1;
            }
        }

        ~UnionFind(){
            delete[] parent;
            delete[] sz;
        }

        int find(int p){
            assert(p >= 0 && p < count);
            while(p!=parent[p])
                p=parent[p];
            return p;
        }

        bool isConnnected( int p, int q){
            return find(p) == find(q);
        }

        void __union(int p, int q){
            int pRoot = find(p);
            int qRoot = find(q);

            if(pRoot == qRoot)
                return;

            if( sz[pRoot] < sz[qRoot] ) {
                parent[pRoot] = qRoot;
                count--;
                sz[qRoot]+=sz[pRoot];
            }
            else{// sz[qRoot]< sz[pRoot]
                parent[qRoot] = pRoot;
                count--;
                sz[pRoot]+=sz[qRoot];
            }
        }
    };
}


#endif //UNIONFIND_UNIONFIND3_H
