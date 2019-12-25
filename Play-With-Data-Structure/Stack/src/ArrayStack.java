public class ArrayStack<Ele> implements Stack<Ele> {
    private Array<Ele> array;

    public ArrayStack(int Capacity) {
        array = new Array<>(Capacity);
    }

    public ArrayStack() {
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

    public int getCapacity() {
        return array.getCapacity();
    }

    @Override
    public void push(Ele e) {
        array.addLast(e);
    }

    @Override
    public Ele pop() {
        return array.removeLast();
    }

    @Override
    public Ele peek() {
        return array.get(array.getSize() - 1);
    }

    @Override
    public String toString() {
        StringBuilder res = new StringBuilder();
        res.append("Stack:");
        res.append('[');
        for (int i = 0; i < array.getSize(); i++) {
            res.append(array.get(i));
            if (i != array.getSize() - 1)
                res.append(',');
        }
        res.append("] top");
        return res.toString();
    }
}

