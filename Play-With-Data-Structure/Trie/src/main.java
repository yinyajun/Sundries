import java.util.TreeMap;

public class main {


    public static void main(String[] args) {
        Trie trie = new Trie();
        trie.add1("1");
        trie.add1("3654");

        System.out.println(trie.contains1("3"));
        System.out.println(trie.getSize());

    }
}
