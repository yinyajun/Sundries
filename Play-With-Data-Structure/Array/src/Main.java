public class Main {
    public static void main(String[] args) {
        Array<Integer> arr = new Array<>(20);
        for (int i = 0; i < 10; i++) {
            arr.addLast(i);
        }
        arr.addLast(3);
        arr.removeFirst();
        arr.removeFirst();
        System.out.println(arr);
    }
}
