public class MaxHeap<E extends Comparable<E>> {

    private Array<E> data;

    public MaxHeap(int capacity) {
        data = new Array<>(capacity);
    }

    public MaxHeap() {
        data = new Array<>();
    }

    public MaxHeap(E[] arr){
        data = new Array<>(arr);
        for (int i = parent(arr.length-1); i >=0 ; i--) {
            siftDown(i);
        }
    }

    public int size() {
        return data.getSize();
    }

    public boolean isEmpty() {
        return data.isEmpty();
    }

    private int parent(int index) {
        if (index == 0)
            throw new IllegalArgumentException("index-0 have no parent.");
        return (index - 1) / 2;
    }

    private int leftChild(int index) {
        return index * 2 + 1;
    }

    private int rightChild(int index) {
        return index * 2 + 2;
    }

    public void add(E e) {
        data.addLast(e);
        // the index of new element
        int index = data.getSize() - 1;
        siftUp(index);
    }

    // maintain the property of maxHeap
    private void siftUp(int k) {
        while (k > 0 && data.get(parent(k)).compareTo(data.get(k)) < 0) {
            data.swap(k, parent(k));
            k = parent(k);
        }
    }

    public E findMax() {
        if (data.getSize() == 0) {
            throw new IllegalArgumentException("cannot findMax when heap is empty");
        }
        return data.get(0);
    }

    public E extractMax() {
        E ret = findMax();
        data.swap(0, data.getSize() - 1);
        data.removeLast();
        siftDown(0);
        return ret;
    }

    private void siftDown(int k) {
        // util k has no child
        while (leftChild(k) < data.getSize()) {
            int maxChildIndex = leftChild(k);
            int j = leftChild(k);
            // j+1 is rightChild, but it is possible not exist
            if (j + 1 < data.getSize() && data.get(j + 1).compareTo(data.get(j)) > 0)
                maxChildIndex = rightChild(k);

            if (data.get(k).compareTo(data.get(maxChildIndex)) > 0) {
                break;
            } else {
                data.swap(k, maxChildIndex);
                k = maxChildIndex;
            }
        }
    }

    // 取出堆中元素，替换成e
    public E replace(E e){
        E ret = findMax();
        data.set(0,e);
        siftDown(0);
        return ret;
    }

}
