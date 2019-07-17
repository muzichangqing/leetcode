<?php
/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val) { $this->val = $val; }
 * }
 */
class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val) { $this->val = $val; }
}

class Solution {

    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function addTwoNumbers($l1, $l2) {
        $carry = 0;
        $header = new ListNode(0);
        $pointer = $header;

        while($l1 != null && $l2 != null) {
            $v1 = $l1->val;
            $v2 = $l2->val;
            $sum = $v1 + $v2 + $carry;
            if ($sum > 9) {
                $sum = $sum % 10;
                $carry = 1;
            } else {
                $carry = 0;
            }
            $pointer->next = new ListNode($sum);
            $pointer = $pointer->next;
            $l1 = $l1->next;
            $l2 = $l2->next;
        }

        while($l1 != null) {
            $v1 = $l1->val;
            $sum = $v1 + $carry;
            if ($sum > 9) {
                $sum = $sum % 10;
                $carry = 1;
            } else {
                $carry = 0;
            }
            $pointer->next = new ListNode($sum);
            $pointer = $pointer->next;
            $l1 = $l1->next;
        }

        while($l2 != null) {
            $v2 = $l2->val;
            $sum = $v2 + $carry;
            if ($sum > 9) {
                $sum = $sum % 10;
                $carry = 1;
            } else {
                $carry = 0;
            }
            $pointer->next = new ListNode($sum);
            $pointer = $pointer->next;
            $l2 = $l2->next;
        }

        if ($carry != 0) {
            $pointer->next = new ListNode($carry);
        }

        return $header->next;
    }
}