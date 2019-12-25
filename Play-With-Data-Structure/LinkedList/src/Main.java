public class Main {
    public static void main(String[] args) {
        LinkedList_dummy<Integer> l = new LinkedList_dummy<Integer>();
        l.addFirst(1);
        l.addFirst(2);
        l.addFirst(3);
        System.out.println(l);
        l.add(2, 666);
        System.out.println(l);
    }
}
