#include <iostream>
#include <queue>
#include <cassert>

using namespace std;

template<typename Key, typename Value>
class BST {
private:
    struct Node {
        Key key;
        Value value;
        Node *left;
        Node *right;

        Node(Key key, Value value) {
            this->key = key;
            this->value = value;
            this->left = this->right = NULL;
        }

        Node(Node* node){
            this->key = node->key;
            this->value=node->value;
            this->left = node->left;
            this->right = node->right;
        }
    };

    Node *root;
    int count;
public:
    BST() {
        root = NULL;
        count = 0;
    }

    ~BST() {
        destroy(root);
    }

    bool isEmpty() {
        return count == 0;
    }

    int size() {
        return count;
    }


    void insert(Key key, Value value) {
        // 树的递归以根节点区分子结构；数组的递归以区间区分子结构
        root = insert(this->root, key, value);
    }

    void insert1(Key key, Value value) {
        Node **node = &root;
        Node *cur = *node;
        Node *cur_parent = cur;
        bool is_left = true;

        while (cur != NULL) {
            if (cur->key == key) {
                cur->value = value;
                return;
            }

            cur_parent = cur;

            if (cur->key < key) {
                is_left = false;
                cur = cur->right;
            } else //cur->key>key
                cur = cur->left;
        }

        if (cur_parent) {
            if (is_left)
                cur_parent->left = new Node(key, value);
            else
                cur_parent->right = new Node(key, value);
        } else
            *node = new Node(key, value);
        count++;
    }

    //非递归插入
    void insert2(Key key, Value value) {
        Node **cur = &root;
        while ((*cur) != NULL) {
            if ((*cur)->key == key) {
                (*cur)->value = value;
                return;
            }
            if ((*cur)->key < key)
                (*cur) = (*cur)->right;
            else //(*cur)->key>key
                (*cur) = (*cur)->left;
        }
        //没有找到这个节点，那么就插入新节点
        //法二：直接将cur替换为newNode,需要cur
        *cur = new Node(key, value);
        count++;
    }


    bool contain(Key key) {
        return contain(root, key);
    }

    Value *search(Key key) {
        return search(root, key);
    }

    void preOrder() {
        preOrder(root);
    }

    void levelOrder() {
        queue<Node *> q;
        q.push(root);
        while (!q.empty()) {
            Node *node = q.front();
            q.pop();

            cout << node->key << endl;
            if (node->left)
                q.push(node->left);
            if (node->right)
                q.push(node->right);
        }
    }


    // 寻找最小的键值
    Key minimum() {
        assert(count != 0);
        Node *minNode = minimum(root);
        return minNode->key;
    }

    // 寻找最大的键值
    Key maximum() {
        assert(count != 0);
        Node *maxNode = maximum(root);
        return maxNode;
    }

    // 从二叉树中删除最小值所在节点
    void removeMin() {
        if (root)
            root = removeMin(root);
    }

    // 从二叉树中删除最大值所在节点
    void removeMax() {
        if (root)
            root = removeMax(root);
    }


    void remove(Key key) {
        root = remove(root, key);
    }


private:
    // 删除掉以node为根的二分搜索树中的键值为key的节点
    // 返回删除节点后的新的二分搜索树的根
    Node *remove(Node *node, Key key) {

        if (node == NULL)
            return NULL;

        // 先要找到这个节点
        if (node->key == key) {
            // 左孩子为空
            if (node->left == NULL) {
                Node *rightNode = node->right;
                delete node;
                count--;
                return rightNode;
            }
            // 右孩子为空
            if (node->right == NULL) {
                Node *leftNode = node->left;
                delete node;
                count--;
                return leftNode;
            }
//            // 左右孩子均不为空
//            Node *s = minimum(node->right);
//            s->right = removeMin(node->right);
//            s->left = node->left;
//            delete node;
//            count--;
//            return s;
            Node *s = new Node(minimum(node->right));
            count++; // removeMin里面count--
            s->right = removeMin(node->right);
            s->left = node->left;
            delete node;
            count--;
            return s;
        } else if (node->key < key)
            node->right = remove(node->right, key);
        else// node->key > key
            node->left = remove(node->left, key);
        return node;
    }


    // 删除掉以node为根的二分搜索树中的最小节点
    // 返回删除节点后的新的二分搜索树的根
    Node *removeMin(Node *node) {

        if (node->left == NULL) {
            Node *rightNode = node->right;
            delete node;
            count--;
            return rightNode;
        }

        node->left = removeMin(node->left);
        return node;
    }


    // 删除掉以node为根的二分搜索树中的最大节点
    // 返回删除节点后的新的二分搜索树的根
    Node *removeMax(Node *node) {

        if (node->right == NULL) {
            Node *leftNode = node->left;
            delete node;
            count--;
            return leftNode;
        }

        node->right = removeMax(node->right);
        return node;
    }


    // 以node为根的二叉搜索树中，返回最小键值的节点
    Node *minimum(Node *node) {
        if (node->left == NULL)
            return node;

        Node *minNode = minimum(node->left);
        return minNode;
    }

    // 以node为根的二叉搜索树中，返回最大键值的节点
    Node *maximum(Node *node) {
        if (node->right == NULL)
            return node;

        Node *maxNode = minimum(node->right);
        return maxNode;
    }

    // 以node为根的二叉树进行前序遍历
    void preOrder(Node *node) {
        if (node == NULL)
            return;

        cout << node->key << endl;
        preOrder(node->left);
        preOrder(node->right);
    }


    // 查看以node为根的二叉搜索树中是否包含键值为key的节点
    bool contain(Node *node, Key key) {
        if (node == NULL)
            return false;

        if (key == node->key)
            return true;
        else if (key < node->key)
            return contain(node->left, key);
        else // key > node->key
            return contain(node->right, key);
    }


    // 在以node为根的二叉搜索树中查找key所对应的value
    Value *search(Node *node, Key key) {

        if (node == NULL)
            return NULL;

        if (key == node->key)
            return &(node->value);
        else if (key < node->key)
            return search(node->left, key);
        else // key > node->key
            return search(node->right, key);
    }


    // 向以node为根的二叉搜索树中，插入节点(key, value)
    // 返回插入新节点后的二叉树的根
    Node *insert(Node *node, Key key, Value value) {
        if (node == NULL) {
            count++;
            return new Node(key, value);
        }

        if (key == node->key)
            node->value = value;
        else if (key < node->key)
            node->left = insert(node->left, key, value);
        else //key > node->key
            node->right = insert(node->right, key, value);
        return node;
    }


    // 使用后序遍历的方式递归的删除节点
    void destroy(Node *node) {
        if (node != NULL) {
            destroy(node->left);
            destroy(node->right);

            delete node;
            count--;
        }
    }
};


int main() {
    BST<string, int> bst = BST<string, int>();
    bst.insert1("a", 1);
    bst.insert("b", 1);
    bst.insert("c", 1);
    bst.insert("d", 1);
    int *res = bst.search("a");
    (*res)++;
    (*res)++;

    cout << *bst.search("a") << endl;


    return 0;
}