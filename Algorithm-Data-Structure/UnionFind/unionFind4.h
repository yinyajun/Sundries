//
// Created by Administrator on 2019/3/3.
//

#ifndef UNIONFIND_UNIONFIND4_H
#define UNIONFIND_UNIONFIND4_H

#include <iostream>
#include <cassert>

using namespace std;

namespace UF4 {
    class UnionFind{
    private:
        int* parent;
        int* rank; // rank[i]表示以i为根的集合中所表示树的层数
        int count;
    public:
        UnionFind(int count){
            parent = new int[count];
            rank = new int[count];
            this->count = count;
            for( int i=0; i<count;i++){
                parent[i]= i;
                rank[i] = 1;
            }
        }

        void __union(int p, int q){
            int pRoot = find(p);
            int qRoot = find(q);

            if(pRoot == qRoot)
                return;

            if( rank[pRoot] < rank[qRoot] ) {
                parent[pRoot] = qRoot;
                count--;

            }
            else if( rank[qRoot]< rank[pRoot] )
            {// sz[qRoot]< sz[pRoot]
                parent[qRoot] = pRoot;
                count--;
            }
            else{
                parent[pRoot] = qRoot;
                rank[qRoot]+=1;
            }
        }
    };
}


#endif //UNIONFIND_UNIONFIND4_H
