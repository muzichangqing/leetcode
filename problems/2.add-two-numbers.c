/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */

#include <stdlib.h>

struct ListNode
{
    int val;
    struct ListNode *next;
};

#ifndef NULL
#define NULL ((void *)0)
#endif

struct ListNode *addTwoNumbers(struct ListNode *l1, struct ListNode *l2)
{
    struct ListNode *p, *lastp;
    p = l1;
    int carry = 0;
    int val2 = 0;
    while(p) {
        if (l2) {
            val2 = l2->val;
            l2 = l2->next;
        } else {
            val2 = 0;
        }
        p->val = p->val + val2 + carry;
        carry = p->val / 10;
        p->val = p->val % 10;
        lastp = p;
        p = p->next;
    }
    
    while(l2) {
        if (carry) {
            l2->val = l2->val + carry;
            carry = l2->val / 10;
            l2->val = l2->val % 10;
        }
        lastp->next = l2;
        lastp = l2;
        l2 = l2->next;
    }

    if (carry) {
        lastp->next = (struct ListNode *)malloc(sizeof(struct ListNode));
        lastp->next->next = NULL;
        lastp->next->val = carry;
        lastp = lastp->next;
    }
    return l1;
}