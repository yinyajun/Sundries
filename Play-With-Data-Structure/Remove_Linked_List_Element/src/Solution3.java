public class Solution3 {
    public ListNode removeElements(ListNode head, int val, int depth) {

        String depthString = generateDepthString(depth);
        System.out.print(depthString);
        System.out.println("Call: remove"+ val + " in" + head);

        if (head == null){
            //递归函数，每层都有递归两个动作
            //基础情况内，不继续（调用自己）递了，同时也不归了（调用自己后面的程序）
            //要想在基础情况内归（肯定不是和正常归的程序完全相同，至少return不同），就得在基础情况里写上。
            System.out.print(depthString);
            System.out.println(("Return: " + head));
            return null;
        }

        // deal with shorter linkedlist without head
        ListNode res = removeElements(head.next, val, depth+1);
        System.out.print(depthString);
        System.out.println(("After remove" + val + ": " + res));

        ListNode ret;
        if (head.val == val)
            ret =res;
        else {
            head.next = res;
            ret =head;

        }
        System.out.print(depthString);
        System.out.println(("Return:" + ret));
        return ret;
    }

    private String generateDepthString(int depth) {
        StringBuilder res = new StringBuilder();
        for (int i = 0; i < depth; i++) {
            res.append("__");
        }
        return res.toString();
    }


    public static void main(String[] args) {
        int[] nums = {1, 2, 6, 3, 4, 5, 6};
        ListNode head = new ListNode(nums);
        System.out.println(head);

        ListNode res = (new Solution3()).removeElements(head, 6, 0);
        System.out.println(res);
    }
}
