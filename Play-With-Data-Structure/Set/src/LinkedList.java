public class LinkedList<E> {
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

    private Node head;
    private int size;

    public LinkedList() {
        head = null;
        size = 0;
    }

    public int getSize() {
        return size;
    }

    public boolean isEmpty() {
        return size == 0;
    }

    public boolean contains(E e) {
        Node cur = head;
        while (cur != null) {
            if (cur.e == e)
                return true;
            cur = cur.next;
        }
        return false;
    }

    public void addFirst(E e) {
//        Node node = new Node(e);
//        node.next = head;
//        head = node;
        head = new Node(e, head);
        size++;
    }

    public void add(int index, E e) {
        // index is valid? we can add an element in the tail(size)
        if (index < 0 || index > size)
            throw new IllegalArgumentException("illegal index");

        // find the previous node
        // head has no previous node
        if (index == 0)
            addFirst(e);
        else {
            Node prev = head;
            for (int i = 0; i < index - 1; i++) {
                prev = prev.next;
            }
//            Node node = new Node(e);
//            node.next = prev.next;
//            prev.next = node;
            prev.next = new Node(e, prev.next);
        }
        size++;
    }

    public void addLast(E e) {
        add(size, e);
    }

    public void removeElement(E e) {
        Node cur = head;
        Node prev = new Node(null, head);
        while (cur != null) {
            if (cur.e.equals(e)) {
                break;
            }
            prev = cur;
            cur = cur.next;
        }

        // 提前终止
        if (cur != null) {
            Node delNode = prev.next;
            prev.next = delNode.next;
            delNode.next = null;
            size--;
        }
    }

    @Override
    public String toString() {
        StringBuilder res = new StringBuilder();

        Node cur = head;
        while (cur != null) {
            res.append(cur).append("->");
            cur = cur.next;
        }
        res.append("NULL");
        return res.toString();
    }
}
