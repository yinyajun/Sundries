import java.util.TreeMap;

public class Trie {
    private class Node {

        public boolean isWord;
        public TreeMap<Character, Node> next;

        public Node(boolean isWord) {
            this.isWord = isWord;
            next = new TreeMap<>();
        }

        public Node() {
            this(false);
        }
    }

    private Node root;
    private int size;

    public Trie() {
        root = new Node();
        size = 0;
    }


    public int getSize() {
        return size;
    }

    //像trie中添加一个新的单词
    public void add(String word) {
        Node cur = root;
        for (int i = 0; i < word.length(); i++) {
            char c = word.charAt(i);
            if (cur.next.get(c) == null)
                cur.next.put(c, new Node());
            cur = cur.next.get(c);
        }
        if (!cur.isWord) {
            cur.isWord = true;
            size++;
        }
    }

    public void add1(String word) {
        root = add11(root, word);
    }

    // 为node为节点的trie树添加string
    private Node add1(Node node, String word) {
        // 当前节点是叶子节点的孩子
        if (node == null)
            node = new Node();
        if (word.length() == 0) {
            node.isWord = true;
            size++;
            return node;
        }

        char c = word.charAt(0);
        String sub_string = word.substring(1, word.length());
        // 这里我犯了一个错误，把这个递归的基本情况搞错了。只有在基本情况下才需要return，表明递归到底了。
        // bst中，添加的节点是叶子节点。当找到一个node==null时候，一定是递归到底了。所以要node==null是基本情况。
        // 而这里，trie添加的节点代表string，当string遍历到最后一个char时，才算递归到底。也就是说node==null的情况下，新建node后
        // 仍然要继续递归。
        // 换句话说，基本情况是建立好string消耗完的trie子树。

        node.next.put(c, add1(node.next.get(c), sub_string));
        return node;
    }

    private Node add11(Node node, String word) {
        if (word.length() == 0) {
            node.isWord = true;
            size++;
            return node;
        }

        char c = word.charAt(0);
        String sub_string = word.substring(1, word.length());

        // 在通往word的这条路径上，当前节点是叶子节点
        if (node.next.get(c) == null)
            node.next.put(c, add1(new Node(), sub_string));
        else
            node.next.put(c, add1(node.next.get(c), sub_string));
        return node;
    }

    public void add2(String word) {
        add2(root, word);
    }

    private void add2(Node node, String word) {
        if (word.length() == 0) {
            node.isWord = true;
            size++;
            return;
        }
        char c = word.charAt(0);
        String sub_string = word.substring(1, word.length());
        if (node.next.get(c) == null) {
            node.next.put(c, new Node());
        }
        add2(node.next.get(c), sub_string);
    }


    public boolean contains(String word) {
        Node cur = root;
        for (int i = 0; i < word.length(); i++) {
            char c = word.charAt(i);
            if (cur.next.get(c) == null)
                return false;
            cur = cur.next.get(c);
        }
        return cur.isWord;
    }

    public boolean contains1(String word) {
        return contains2(root, word);
    }

    private boolean contains1(Node node, String word) {
        // 以cur为节点，保证cur非空
        if (node == null)
            return false;
        if (word.length() == 0) {
            return node.isWord;
        }
        char c = word.charAt(0);
        String sub_word = word.substring(1, word.length());
        return contains1(node.next.get(c), sub_word);
    }

    private boolean contains2(Node node, String word) {
        if (word.length() == 0) {
            return node .isWord;
        }
        char c = word.charAt(0);
        String sub_word = word.substring(1, word.length());

        // 在通往word的这条路径上，当前节点是叶子节点
        if (node.next.get(c) == null)
            return false;
        return contains1(node.next.get(c), sub_word);
    }

    public boolean isPrefix(String prefix) {
        Node cur = root;
        for (int i = 0; i < prefix.length(); i++) {
            char c = prefix.charAt(i);
            if (cur.next.get(c) == null)
                return false;
            cur = cur.next.get(c);
        }
        return true;
    }
}
