public class SegmentTree<E> {
    private E[] data;
    private E[] tree;
    private Merger<E> merger;

    public SegmentTree(E[] arr, Merger<E> merger) {
        this.merger = merger;
        data = (E[]) new Object[arr.length];
        for (int i = 0; i < arr.length; i++) {
            data[i] = arr[i];
        }
        tree = (E[]) new Object[4 * arr.length];
        buildSegmentTree(0, 0, data.length - 1);
    }

    //在treeIndex的位置创建表示区间[l,r]的线段树
    private void buildSegmentTree(int treeIndex, int l, int r) {
        if (l == r) {
            tree[treeIndex] = data[l];
            return;
        }
        // 当这个节点不是叶子节点的时候，先建立左右子树
        int leftTreeIndex = leftChild(treeIndex);
        int rightTreeIndex = rightChild(treeIndex);
        int mid = l + (r - l) / 2;
        buildSegmentTree(leftTreeIndex, l, mid);
        buildSegmentTree(rightTreeIndex, mid+1, r);

        tree[treeIndex] = merger.merge(tree[leftTreeIndex], tree[rightTreeIndex]);
    }

    public void set(int index, E e){
        if (index<0||index>=data.length)
            throw new IllegalArgumentException("index is illegal");
        data[index] = e;
        set(0,0,data.length,index, e);
    }

    // 更新以treeIndex为根的线段树中index的值为e
    private void set(int treeIndex, int l, int r, int index, E e){
        if (l == r) {
            tree[treeIndex] = e;
            return;
        }

        int leftTreeIndex = leftChild(treeIndex);
        int rightTreeIndex = rightChild(treeIndex);
        int mid = l + (r - l) / 2;
        if(index>=mid+1)
            set(rightTreeIndex,mid+1,r,index,e);
        else //index<=mid
            set(leftTreeIndex,0,mid,index,e);
        tree[treeIndex] = merger.merge(tree[leftTreeIndex], tree[rightTreeIndex]);
    }


    public E query(int queryL, int queryR) {
        if (queryL < 0 || queryL >= data.length || queryR < 0 || queryR >= data.length || queryL > queryR)
            throw new IllegalArgumentException("Index is illegal.");
        return query(0,0,data.length-1,queryL,queryR);
    }

    private E query(int treeIndex, int l, int r, int queryL, int queryR) {
        if (l == queryL && r == queryR)
            return tree[treeIndex];

        int mid = l + (r - l) / 2;
        int leftTreeIndex = leftChild(treeIndex);
        int rightTreeIndex = rightChild(treeIndex);
        if (queryL > mid) {
            return query(rightTreeIndex, mid + 1, r, queryL, queryR);
        } else if (queryR <= mid)
            return query(leftTreeIndex, l, mid, queryL, queryR);
        else {
            E leftResult = query(leftTreeIndex, l, mid, queryL, mid);
            E rightResult = query(rightTreeIndex, mid + 1, r, mid + 1, queryR);
            return merger.merge(leftResult, rightResult);
        }


    }

    public int getSize() {
        return data.length;
    }

    public E get(int index) {
        if (index < 0 || index >= data.length) {
            throw new IllegalArgumentException("index is illegal");
        }
        return data[index];
    }

    private int leftChild(int index) {
        return 2 * index + 1;
    }

    private int rightChild(int index) {
        return 2 * index + 2;
    }


}
