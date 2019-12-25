public interface Queue<Ele> {
    int getSize();

    boolean isEmpty();

    void enqueue(Ele e);

    Ele dequeue();

    Ele getFront();
}
