from typing import List, Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def mergeTwoLists(self,
                      list1: Optional[ListNode],
                      list2: Optional[ListNode]) -> Optional[ListNode]:
        if list1 is None:
            return list2
        if list2 is None:
            return list1
        head = ListNode()
        cursor = head
        while list1 is not None and list2 is not None:
            if list1.val < list2.val:
                cursor.next = list1
                list1 = list1.next
            else:
                cursor.next = list2
                list2 = list2.next
            cursor = cursor.next
        if list1 is not None:
            cursor.next = list1
        else:
            cursor.next = list2
        return head.next

    def removeDuplicates(self, nums: List[int]) -> int:
        slow, quick = 0, 1
        while quick < len(nums):
            if nums[quick] != nums[quick-1]:
                slow += 1
                nums[slow] = nums[quick]
            quick += 1
        return slow+1
        