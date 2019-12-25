public class LoopQueue<Ele> implements Queue<Ele> {
    private Ele[] data;
    private int front;
    private int tail;
    private int size;

    public LoopQueue(int capacity) {
        data = (Ele[]) new Object[capacity + 1];
        front = 0;
        tail = 0;
        size = 0;
    }

    public LoopQueue() {
        this(10);
    }

    public int getCapacity() {
        return data.length - 1;
    }

    @Override
    public int getSize() {
        return size;
    }

    @Override
    public boolean isEmpty() {
        return front == tail;
    }

    @Override
    public void enqueue(Ele e) {
        // queue is full?
        if ((tail + 1) % data.length == front)
            resize(getCapacity() * 2);
        data[tail] = e;
        tail = (tail + 1) % data.length;
        size++;
    }

    @Override
    public Ele dequeue() {
        if (isEmpty())
            throw new IllegalArgumentException("queue is empty, cannot dequeue.");
        Ele res = data[front];
        front = (front + 1) % data.length;
        size--;
        if (size == getCapacity() / 4 && getCapacity() / 2 != 0)
            resize(getCapacity() / 2);
        return res;
    }

    @Override
    public Ele getFront() {
        if (isEmpty())
            throw new IllegalArgumentException("queue is empty.");
        return data[front];
    }

    private void resize(int newCapacity) {
        Ele[] newData = (Ele[]) new Object[newCapacity + 1];
        for (int i = 0; i < size; i++) {
            newData[i] = data[(i + front) % data.length];
        }
        data = newData;
        front = 0;
        tail = size;
    }

    @Override
    public String toString() {
        StringBuilder res = new StringBuilder();
        res.append("Queue: size = %d, capacity = %d\n", size, getCapacity());
        res.append("front [");
        for (int i = front; i != tail; i = (i + 1) % data.length) {
            res.append(data[i]);
            if ((i + 1) % data.length != tail)
                res.append(',');
        }
        res.append(']');
        return res.toString();
    }
}
