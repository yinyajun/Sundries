public class BSTSET<E extends Comparable<E>> implements Set<E> {

    private BST<E> bst;

    public BSTSET(){
        bst = new BST<>();
    }

    @Override
    public int getSize() {
        return bst.getSize();
    }

    @Override
    public boolean isEmpty() {
        return bst.isEmpty();
    }

    @Override
    public void add(E e) {
        bst.add_1(e);
    }

    @Override
    public boolean contains(E e) {
        return bst.contains(e);
    }

    @Override
    public void remove(E e) {
        bst.remove(e);
    }


    public static void main(String[] args) {
        BSTSET<Integer> set = new BSTSET<>();
        set.add(1);
        set.add(2);
        set.add(1);
        set.add(1);
        System.out.println(set.bst);
    }
}
