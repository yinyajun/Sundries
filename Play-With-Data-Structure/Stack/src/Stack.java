public interface Stack<Ele> {

    int getSize();
    boolean isEmpty();
    void push(Ele e);
    Ele pop();
    Ele peek();
}
