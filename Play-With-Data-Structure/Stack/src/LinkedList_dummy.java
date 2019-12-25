public class LinkedList_dummy<E> {
    private class Node {
        public E e;
        public Node next;

        public Node(E e, Node next) {
            this.e = e;
            this.next = next;
        }

        public Node(E e) {
            this(e, null);
        }

        public Node() {
            this(null, null);
        }

        @Override
        public String toString() {
            return e.toString();
        }
    }

    private Node dummyHead;
    private int size;

    public LinkedList_dummy() {
        //dummy head points to -1 node(meaningless)
        dummyHead = new Node(null, null);
        size = 0;
    }

    public int getSize() {
        return size;
    }

    public boolean isEmpty() {
        return size == 0;
    }

    public void addFirst(E e) {
        add(0, e);
    }

    public void add(int index, E e) {
        // index is valid? we can add an element in the tail(size)
        if (index < 0 || index > size)
            throw new IllegalArgumentException("illegal index");

        // find the previous node
        Node prev = dummyHead;
        for (int i = 0; i < index; i++) {
            prev = prev.next;
        }
//            Node node = new Node(e);
//            node.next = prev.next;
//            prev.next = node;
        prev.next = new Node(e, prev.next);
        size++;
    }

    public void addLast(E e) {
        add(size, e);
    }

    public E get(int index) {
        if (index < 0 || index >= size)
            throw new IllegalArgumentException("illegal index");
        Node cur = dummyHead.next;
        for (int i = 0; i < index; i++) {
            cur = cur.next;
        }
        return cur.e;
    }

    public E getFirst() {
        return get(0);
    }

    public E getLast() {
        return get(size);
    }

    public void set(int index, E e) {
        if (index < 0 || index >= size)
            throw new IllegalArgumentException("illegal index");
        Node cur = dummyHead.next;
        for (int i = 0; i < index; i++) {
            cur = cur.next;
        }
        cur.e = e;
    }

    public boolean contains(E e) {
        Node cur = dummyHead.next;
        while (cur != null) {
            if (cur.e == e)
                return true;
            cur = cur.next;
        }
        return false;
    }

    public E remove(int index) {
        if (index < 0 || index >= size)
            throw new IllegalArgumentException("illegal index");

        Node prev = dummyHead;
        for (int i = 0; i < index; i++) {
            prev = prev.next;
        }
        Node delNode = prev.next;
        prev.next = delNode.next;

        delNode.next = null;
        size--;
        return delNode.e;
    }

    public E removeLast() {
        return remove(size - 1);
    }

    public E removeFirst() {
        return remove(0);
    }


    @Override
    public String toString() {
        StringBuilder res = new StringBuilder();

        Node cur = dummyHead.next;
        while (cur != null) {
            res.append(cur + "->");
            cur = cur.next;
        }
        res.append("NULL");
        return res.toString();
    }

}
