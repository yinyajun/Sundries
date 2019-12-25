import java.util.TreeMap;

public class WordDictionary {

    private class Node {
        public boolean isWord;
        public TreeMap<Character, Node> children;

        public Node(boolean isWord) {
            this.isWord = isWord;
            children = new TreeMap<>();
        }

        public Node() {
            this(false);
        }
    }

    private Node root;


    public WordDictionary() {
        root = new Node();
    }

    public void addWord(String word) {
        Node cur = root;
        for (int i = 0; i < word.length(); i++) {
            char c = word.charAt(i);
            if (cur.children.get(c) == null) {
                cur.children.put(c, new Node());
            }
            cur = cur.children.get(c);
        }
        cur.isWord = true;
    }

    public boolean search(String word) {

        return search(root, word, 0);
    }

    private boolean search(Node node, String word, int index) {
        if (index == word.length() - 1) {
            return node.isWord;
        }

        char c = word.charAt(index);
        if (c != '.') {
            if (node.children.get(c) == null)
                return false;
            return search(node.children.get(c), word, index + 1);
        } else {
            for (char nextCHar : node.children.keySet()) {
                if (search(node.children.get(nextCHar), word, index + 1))
                    return true;
            }
            return false;
        }
    }
}
