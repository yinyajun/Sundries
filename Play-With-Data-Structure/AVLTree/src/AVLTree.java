import java.util.ArrayList;

public class AVLTree<K extends Comparable<K>, V> implements Map<K, V> {

    private class Node {

        public K key;
        public V value;
        public Node left, right;
        public int height;


        public Node(K key, V value) {
            this.key = key;
            this.value = value;
            left = right = null;
            height = 1;
        }
    }

    private Node root;
    private int size;

    public AVLTree() {
        root = null;
        size = 0;
    }

    @Override
    public int getSize() {
        return size;
    }

    @Override
    public boolean isEmpty() {
        return size == 0;
    }

    public boolean isBST() {
        ArrayList<K> keys = new ArrayList<>();
        inOrder(root, keys);
        for (int i = 1; i < keys.size(); i++) {
            if (keys.get(i - 1).compareTo(keys.get(i)) > 0)
                return false;
        }
        return true;
    }


    public boolean isBalanced() {
        return isBalanced(root);
    }

    private boolean isBalanced(Node node) {
        if (node == null)
            return true;

        int balanceFactor = getBalaceFactor(node);
        if (Math.abs(balanceFactor) > 1)
            return false;
        return isBalanced(node.left) && isBalanced(node.right);
    }

    private void inOrder(Node node, ArrayList<K> keys) {
        if (node == null)
            return;
        inOrder(node.left, keys);
        keys.add(node.key);
        inOrder(node.right, keys);
    }

    private int getHeight(Node node) {
        if (node == null)
            return 0;
        return node.height;
    }

    private int getBalaceFactor(Node node) {
        if (node == null)
            return 0;
        return getHeight(node.left) - getHeight(node.right);

    }

    @Override
    public void add(K key, V value) {
        root = add(root, key, value);
    }

    private Node add(Node node, K key, V value) {
        if (node == null) {
            size++;
            return new Node(key, value);
        }
        if (key.compareTo(node.key) < 0) {
            node.left = add(node.left, key, value);
        } else if (key.compareTo(node.key) > 0)
            node.right = add(node.right, key, value);
        else {
            node.value = value;
        }

        // 更新height值
        node.height = 1 + Math.max(getHeight(node.left), getHeight(node.right));

        // 计算平衡因子
        int balanceFactor = getBalaceFactor(node);
        if (Math.abs(balanceFactor) > 1)
            System.out.println("unbalanced");

        // 维护平衡

        return node;
    }

    // 辅助函数，返回以node为根节点的bst中，key所在的节点
    private Node getNode(Node node, K key) {
        if (node == null)
            return null;

        if (key.compareTo(node.key) < 0) {
            return getNode(node.left, key);
        } else if (key.compareTo(node.key) > 0) {
            return getNode(node.right, key);
        } else
            return node;
    }

    @Override
    public boolean contains(K key) {
        return getNode(root, key) != null;
    }

    @Override
    public V get(K key) {
        Node node = getNode(root, key);
        return node == null ? null : node.value;
    }

    @Override
    public void set(K key, V newValue) {
        Node node = getNode(root, key);
        if (node == null)
            throw new IllegalArgumentException(key + "doesn't exist!");
        else
            node.value = newValue;
    }

    @Override
    public V remove(K key) {
        Node node = getNode(root, key);
        if (node != null) {
            root = remove(root, key);
            return node.value;
        }
        return null;
    }

    // 寻找后继的bst删除节点方法依赖两个私有方法：minimum，removeMin
    private Node minimun(Node node) {
        if (node.left == null)
            return node;
        return minimun(node.left);
    }

    private Node removeMin(Node node) {
        if (node.left == null) {
            Node rightNode = node.right;
            node.right = null;
            size--;
            return rightNode;
        }
        node.left = removeMin(node.left);
        return node;
    }

    private Node remove(Node node, K key) {
        if (node == null) {
            return null;
        }

        if (key.compareTo(node.key) < 0) {
            node.left = remove(node.left, key);
            return node;
        } else if (key.compareTo(node.key) > 0) {
            node.right = remove(node.right, key);
            return node;
        } else { // key == node.key
            if (node.left == null) {
                Node rightNode = node.right;
                rightNode.right = null;
                size--;
                return rightNode;
            } else if (node.right == null) {
                Node leftNode = node.left;
                leftNode.left = null;
                size--;
                return leftNode;
            } else {
                Node successor = minimun(node.right);
                successor.right = removeMin(node.right);
                successor.left = node.left;
                node.left = node.right = null;
                return successor;
            }
        }


    }
}
