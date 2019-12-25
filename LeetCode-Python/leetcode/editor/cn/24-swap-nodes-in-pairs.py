#给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。 
#
# 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。 
#
# 
#
# 示例: 
#
# 给定 1->2->3->4, 你应该返回 2->1->4->3.
# 
# Related Topics 链表



#leetcode submit region begin(Prohibit modification and deletion)
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution:
    def swapPairs(self, head: ListNode) -> ListNode:
        dummy_head = ListNode(None)
        dummy_head.next = head

        if not head:
            return head

        pre = dummy_head
        first = pre.next
        second = first.next

        while first and second:
            pre.next = second
            tmp = second.next
            second.next = first
            first.next = tmp

            pre = first
            first = first.next  # first已经可能是None了
            second = first.next if first else None

        return dummy_head.next
        
#leetcode submit region end(Prohibit modification and deletion)
