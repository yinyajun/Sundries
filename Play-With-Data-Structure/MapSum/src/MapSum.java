import java.util.TreeMap;

public class MapSum {

    private class Node {
        public int value;
        public TreeMap<Character, Node> children;

        public Node(int value) {
            this.value = value;
            children = new TreeMap<>();
        }

        public Node() {
            this(0);
        }
    }

    private Node root;

    public MapSum() {
        root = new Node();
    }

    public void insert(String key, int val) {
        Node cur = root;
        for (int i = 0; i < key.length(); i++) {
            char c = key.charAt(i);
            if (cur.children.get(c) == null)
                cur.children.put(c, new Node());
            cur = cur.children.get(c);
        }
        cur.value = val;
    }


    public int sum(String prefix) {
        Node cur = root;
        for (int i = 0; i < prefix.length(); i++) {
            char c = prefix.charAt(i);
            if (cur.children.get(c) == null)
                return 0;
            cur = cur.children.get(c);
        }
        // last node
        // calc sum of value in each subtree
        return sum(cur);

    }

    private int sum(Node node) {
        // leaf
        if (node.children.size() == 0)
            return node.value;

        int ret = node.value;
        for (char c : node.children.keySet()) {
            ret += sum(node.children.get(c));

        }
        return ret;

    }


}
