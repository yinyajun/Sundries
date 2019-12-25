import sun.awt.image.ImageWatched;

public class Linked_list_dummy_recursive<E> {
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

    public Linked_list_dummy_recursive(){
        dummyHead = new Node(null,null);
        size = 0;
    }

    public void add(int index, E e) {
        if (index < 0 || index > size)
            throw new IllegalArgumentException("illegal index");
        add(dummyHead, index, e);
    }

    private void add(Node prev, int index, E e) {
        if (index == 0) {
            size++;
            prev.next = new Node(e, prev.next);
            return;
        }
        add(prev.next, index - 1, e);
    }


    public E remove(int index){
        if (index < 0 || index >= size)
            throw new IllegalArgumentException("illegal index");
        return remove(dummyHead, index);
    }

    private E remove(Node prev, int index){
        if(index == 0){
            size --;
            Node delNode = prev.next;
            prev.next = delNode.next;
            delNode.next = null;
            return delNode.e;
        }
        return remove(prev.next,index-1);
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

    public static void main(String[] args) {
        Linked_list_dummy_recursive<Integer> l = new Linked_list_dummy_recursive<>();
        l.add(0, 5);
        System.out.println(l);
        l.add(1,2);
        System.out.println(l);
        l.add(2,3);
        System.out.println(l);
        l.remove(1);
        System.out.println(l);

    }

}
