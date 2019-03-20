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
    int n1 = 0;
    int n2 = 0;
    int i = 0;

    while (l1)
    {
        n1 += l1->val * pow10(i);
        l1 = l1->next;
        ++i;
    }

    i = 0;
    while (l2)
    {
        n2 += l2->val * pow10(i);
        l2 = l2->next;
        ++i;
    }

    struct ListNode *rs = (struct ListNode *)malloc(sizeof(struct ListNode));
    struct ListNode *p = rs;

    int sum = n1 + n2;
    int mod = sum % 10;
    sum = sum / 10;
    p->val = mod;
    p->next = NULL;

    while (sum)
    {
        mod = sum % 10;
        sum = sum / 10;
        p->next = (struct ListNode *)malloc(sizeof(struct ListNode));
        p->next->val = mod;
        p->next->next = NULL;
        p = p->next;
    }
    return rs;
}

int pow10(int i)
{
    int rs = 1;
    while (i)
    {
        rs *= 10;
        --i;
    }
    return rs;
}