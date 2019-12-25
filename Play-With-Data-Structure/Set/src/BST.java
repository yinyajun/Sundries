import java.util.LinkedList;
import java.util.Queue;
import java.util.Stack;

public class BST<E extends Comparable<E>> {
    private class Node {
        public E e;
        public Node left, right;

        public Node(E e) {
            this.e = e;
            left = right = null;
        }
    }

    private Node root;
    private int size;

    public BST() {
        root = null;
        size = 0;
    }

    public int getSize() {
        return size;
    }

    public boolean isEmpty() {
        return size == 0;
    }

    public void add_1(E e) {
        if (root == null) {
            root = new Node(e);
            size++;
        } else {
            //递归需要表示更小规模的树，用root节点表示不同的树
            add_1(root, e);
        }
    }

    // 向以node为根的二分搜索树中插入元素e，递归算法
    private void add_1(Node node, E e) {
        // 基础情况，想象是一个节点的树，需要添加节点，怎么添加
        if (node.e == e) {
            return;
        } else if (e.compareTo(node.e) < 0 && node.left == null) {
            node.left = new Node(e);
            size++;
            return;
        } else if (e.compareTo(node.e) > 0 && node.right == null) {
            node.right = new Node(e);
            size++;
            return;
        }

        if (e.compareTo(node.e) < 0)
            add_1(node.left, e);
        else
            add_1(node.right, e);
    }


    // add_1的方法，并没有完全递归，因为首先判断根的node非空，也就是为非空的树添加元素
    // 反映在递归方法上，就是在基础情况中，为叶子节点创造新的孩子（已有偏序关系），
    // 那么原先的叶子节点和新的孩子节点的偏序关系已经确定。

    // 但是null也是一棵树，基础情况就是遇到null的子树，创建新节点
    // 通过归的过程去完成偏序关系的建立
    public void add_2(E e) {
        root = add_2(root, e);
    }

    // 返回新的bst的根节点
    private Node add_2(Node node, E e) {
        //  基础情况
        if (node == null) {
            // 创建新节点，要维护size
            node = new Node(e);
            size++;
            return node;
        }

        if (e.compareTo(node.e) < 0) {
            node.left = add_2(node.left, e);
        } else if (e.compareTo(node.e) > 0) {
            node.right = add_2(node.right, e);
        }
        return node;
    }

    // non-recursive
    public void add_3(E e) {
        if (root == null) {
            root = new Node(e);
            size++;
        }
        Node cur = root;
        Node prev = root;
        boolean isleft = true;
        while (cur != null) {
            if (e.compareTo(cur.e) == 0)
                return;
            else if (e.compareTo(cur.e) < 0) {
                prev = cur;
                isleft = true;
                cur = cur.left;
            } else {
                prev = cur;
                isleft = false;
                cur = cur.right;
            }
        }
        if (isleft) {
            prev.left = new Node(e);
            size++;
        } else {
            prev.right = new Node(e);
            size++;
        }
    }

    public boolean contains(E e) {
        return contains(root, e);
    }

    // 以node为根的bst中是否包含e
    private boolean contains(Node node, E e) {
        // 终止情况
        if (node == null) {
            return false;
        }

        if (e.compareTo(node.e) == 0)
            return true;
        else if (e.compareTo(node.e) < 0)
            return contains(node.left, e);
        else
            return contains(node.right, e);
    }

    public void preOrder() {
        preOrder(root);
    }

    private void preOrder(Node node) {
        if (node == null)
            return;
        System.out.println(node.e);
        preOrder(node.left);
        preOrder(node.right);
    }


    public void preOrderNR() {
        Stack<Node> stack = new Stack<>();
        stack.push(root);
        while (!stack.isEmpty()) {
            Node cur = stack.pop();
            if (cur == null)
                continue;
            System.out.println(cur.e);
            if (cur.right != null)
                stack.push(cur.right);
            if (cur.left != null)
                stack.push(cur.left);
        }

    }

    public void levelOrder() {
        Queue<Node> q = new LinkedList<>();
        q.add(root);
        while (!q.isEmpty()) {
            Node cur = q.remove();
            if (cur == null)
                continue;
            System.out.println(cur.e);
            if (cur.left != null)
                q.add(cur.left);
            if (cur.right != null)
                q.add(cur.right);
        }

    }

    public E minimun() {
        if (size == 0)
            throw new IllegalArgumentException("BST is empty");
        return minimum(root).e;
    }

    private Node minimum(Node node) {
        if (node.left == null)
            return node;
        return minimum(node.left);
    }

    public E maximum() {
        if (size == 0)
            throw new IllegalArgumentException("BST is empty");
        return maximum(root).e;
    }

    private Node maximum(Node node) {
        if (node.right == null)
            return node;
        return maximum(node.right);
    }


    public E removeMin() {
        E ret = minimun();
        root = removeMin(root);
        return ret;
    }

    // 删除掉以node为根的bst的最小节点
    // 返回删除节点后的新的bst的根
    private Node removeMin(Node node) {
        // 一般情况下，树的节点都是操作当前节点cur，除非像这种一直找左子树的，类似于链表的操作，才能用prev
        if (node.left == null) {
            Node rightNode = node.right;
            // 将最小值的节点的右子树 与 原bst脱离关系
            node.right = null;
            size --;
            return rightNode;
        }
        node.left = removeMin(node.left);
        return node;
    }

    public E removeMax(){
        E ret = maximum();
        root = removeMax(root);
        return ret;
    }


    private Node removeMax(Node node){
        // 终止情况仅仅是新的子树的根，偏序关系要在归的过程中
        if(node.right==null){
            Node leftNode = node.left;
            node.left = null;
            size--;
            return leftNode;
        }
        node.right = removeMax(node.right);
        return node;
    }


    public void remove(E e){
        root = remove(root, e);
    }

    // 删除以node为根的bst中值为e的节点
    // 返回删除节点后新的bst的根
    private Node remove(Node node, E e){
        if (node == null)
            return null;

        if (e.compareTo(node.e)< 0){
            node.left = remove(node.left,e);
            return node;
        }
        else if(e.compareTo(node.e)>0){
            node.right = remove(node.right, e);
            return node;
        }
        else{// e == node.e
            // 待删除节点左子树为空
            if (node.left ==  null){
                Node rightNode = node.right;
                node.right = null;
                size --;
                return rightNode;
            }
            // 待删除节点右子树为空
            if(node.right==null){
                Node leftNode = node.left;
                node.left = null;
                size--;
                return leftNode;
            }
            // 左右子树都存在的情况
            // 将后继节点代替原节点
            Node successor = minimum(node.right);
            successor.right = removeMin(node.right);
            successor.left = node.left;
            // size不用维护了，在removeMin中减过一次了

            node.left = node.right=null;
            return successor;


        }
    }


    public void print(Node root) {
        if (root == null)
            return;
        print(root.left);
        System.out.print(root.e);
        System.out.print(' ');
        print(root.right);
    }

    @Override
    public String toString() {
        StringBuilder res = new StringBuilder();
        generateBSTString(root, 0, res);
        return res.toString();
    }

    private void generateBSTString(Node node, int depth, StringBuilder res) {
        if (node == null) {
            res.append(generateDepthString(depth) + "null\n");
            return;
        }
        res.append(generateDepthString(depth) + node.e + '\n');
        generateBSTString(node.left, depth + 1, res);
        generateBSTString(node.right, depth + 1, res);
    }

    private String generateDepthString(int depth) {
        StringBuilder res = new StringBuilder();
        for (int i = 0; i < depth; i++) {
            res.append("--");
        }
        return res.toString();
    }

    public static void main(String[] args) {
        BST<Integer> bst = new BST<>();
        bst.add_3(5);
        bst.add_3(7);
        bst.add_3(3);
        bst.add_3(2);
        bst.add_3(4);
        bst.print(bst.root);
        System.out.println();
        bst.removeMax();
        bst.print(bst.root);

    }

}






