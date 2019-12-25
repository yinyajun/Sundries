public class Array<Element> {
    private Element[] data;
    private int size;

    Array(int capacity) {
        this.data = (Element[]) new Object[capacity];
    }

    public Array() {
        this(10);
    }

    public int getSize() {
        return this.size;
    }

    public int getCapacity() {
        return this.data.length;
    }

    public boolean isEmpty() {
        return this.size == 0;
    }

    void addLast(Element e) {
        add(size, e);
    }

    public void addFirst(Element e) {
        add(0, e);
    }

    Element get(int index) {
        if (index < 0 || index >= size)
            throw new IllegalArgumentException("INVALID INDEX.");
        return data[index];
    }

    void set(int index, Element e) {
        if (index < 0 || index >= size)
            throw new IllegalArgumentException("INVALID INDEX.");
        data[index] = e;
    }

    public void add(int index, Element e) {
        if (index < 0 || index > this.size) {
            throw new IllegalArgumentException("INVALID INDEX.");
        }
        if (this.size == this.data.length) {
            resize(2 * data.length);
        }
        for (int i = size - 1; i >= index; i--) {
            this.data[i + 1] = this.data[i];
        }
        this.data[index] = e;
        this.size++;
    }

    private void resize(int newCapacity) {
        Element[] newData = (Element[]) new Object[newCapacity];
        for (int i = 0; i < size; i++) {
            newData[i] = data[i];
        }
        data = newData;
    }

    public boolean contains(Element e) {
        for (int i = 0; i < size; i++) {
            if (data[i].equals(e))
                return true;
        }
        return false;
    }

    public int find(Element e) {
        for (int i = 0; i < size; i++) {
            if (data[i].equals(e))
                return i;
        }
        return -1;
    }

    public Element remove(int index) {
        if (index < 0 || index >= size)
            throw new IllegalArgumentException("INVALID INDEX.");
        Element e = data[index];
        for (int i = index; i < size - 1; i++) {
            data[i] = data[i + 1];
        }
        size--;
        data[size] = null;
        if (size == data.length / 4 && data.length/2 != 0) {
            resize(data.length/2);
        }
        return e;
    }

    public Element removeFirst() {
        return remove(0);
    }

    public Element removeLast() {
        return remove(size - 1);
    }

    public void removeElement(Element e) {
        int index = find(e);
        if (index != -1) {
            remove(index);
        }
    }


    @Override
    public String toString() {
        StringBuilder res = new StringBuilder();
        res.append(String.format("Array:size =%d, capacity=%d\n", size, data.length));
        res.append('[');
        for (int i = 0; i < size; i++) {
            res.append(data[i]);
            if (i != size)
                res.append(',');
        }
        res.append(']');
        return res.toString();
    }
}
