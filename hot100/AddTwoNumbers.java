package hot100;

public class AddTwoNumbers {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode result = new ListNode(0);

        ListNode c1 = l1, c2 = l2, cr = result;
        int val1, val2, sum;

        while (c1 != null || c2 != null) {
            val1 = c1 == null ? 0 : c1.val;
            val2 = c2 == null ? 0 : c2.val;

            sum = val1 + val2 + cr.val;

            cr.val = sum / 10;
            cr.next = new ListNode(sum % 10);

            cr = cr.next;
            if (c1 != null) {
                c1 = c1.next;
            }
            if (c2 != null) {
                c2 = c2.next;
            }
        }

        if (cr.val == 0 && result.next != null) {
            ListNode cursor = result;
            while(cursor.next.next != null) {
                cursor = cursor.next;
            }
            cursor.next = null;
        }

        return result;
    }
}
