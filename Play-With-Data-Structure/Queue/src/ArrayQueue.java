public class ArrayQueue<Ele> implements Queue<Ele> {
    private Array<Ele> array;

    public ArrayQueue(int Capacity){
        array = new Array<>(Capacity);
    }

    public ArrayQueue(){
        array = new Array<>();
    }

    @Override
    public int getSize() {
        return array.getSize();
    }

    @Override
    public boolean isEmpty() {
        return array.isEmpty();
    }

    public int getCapacity(){
        return array.getCapacity();
    }

    @Override
    public void enqueue(Ele e) {
        array.addLast(e);
    }

    @Override
    public Ele dequeue() {
        return array.removeFirst();
    }

    @Override
    public Ele getFront() {
        return array.get(0);
    }

    @Override
    public String toString() {
        StringBuilder res = new StringBuilder();
        res.append("Queue: ");
        res.append("front [");
        for (int i = 0; i < array.getSize(); i++) {
            res.append(array.get(i));
            if (i != array.getSize() - 1)
                res.append(',');
        }
        res.append(']');
        return res.toString();
    }
}
