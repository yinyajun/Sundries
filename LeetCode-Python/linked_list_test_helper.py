class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class LinkedList:
    def __init__(self):
        self.length = 0
        self.dummy_head = ListNode(None)

    def insert(self, index: int, num):
        assert self.length >= index >= 0
        pre = self.dummy_head

        # find i-idx pre ele
        while index:
            pre = pre.next
            index -= 1
        # new node
        node = ListNode(num)
        node.next = pre.next

        pre.next = node
        self.length += 1

    def append(self, num):
        self.insert(self.length, num)

    def add_first(self, num):
        self.insert(0, num)

    def get(self, index):
        assert self.length - 1 >= index >= 0
        cur = self.dummy_head.next

        while index:
            cur = cur.next
            index -= 1
        return cur.val

    def remove(self, index):
        assert self.length - 1 >= index >= 0

        pre = self.dummy_head
        while index:
            pre = pre.next
            index -= 1

        del_node = pre.next
        pre.next = del_node.next
        del_node.next = None
        self.length -= 1
        return del_node.val

    def head(self):
        return self.dummy_head.next


def create_linked_list(array):
    ll = LinkedList()
    for i in array:
        ll.append(i)
    return ll


def traverse(head):
    res = []
    cur = head
    while cur:
        res.append(cur.val)
        cur = cur.next
    print(res)


class Solution:
    def insertionSortList(self, head: ListNode) -> ListNode:
        if not head:
            return head

        dummy_head = ListNode(None)
        dummy_head.next = head

        p = dummy_head
        pre = dummy_head.next
        cur = pre.next  # 当前待排序的节点
        while cur:
            if cur.val >= pre.val:
                pre = pre.next
            else:
                if p.val and p.val > cur.val:
                    p = dummy_head
                while p.next != cur and p.next.val <= cur.val:
                    p = p.next
                if p.next != cur:  # p.next.val > cur.val
                    pre.next = cur.next  # cur will be removed
                    cur.next = p.next
                    p.next = cur
                else:  # pre.next == cur
                    pre = pre.next
            cur = pre.next
        return dummy_head.next


# class Solution:
#     def insertionSortList(self, head: ListNode) -> ListNode:
#         if not head or not head.next:
#             return head
#         _head = ListNode(0)
#         p = _head
#         while head:
#             cur = head.next
#             if p.val > head.val:
#                 p = _head
#
#             while p.next and p.next.val < head.val:
#                 p = p.next
#
#             head.next = p.next
#             p.next = head
#             head = cur
#
#         return _head.next


if __name__ == '__main__':
    s = Solution()

    ll = create_linked_list([-1, 5, 3, 4, 0])
    traverse(Solution().insertionSortList(ll.head()))
