var ListNode = function(val) {
    this.val= val;
    this.next = null;
};

/**
 * @param {ListNode} head
 * @return {boolean}
 */
var hasCycle = function(head) {
    let slow = head;
    let fast = head;

    while (fast !== null) {
        fast = fast.next;
        if (fast === null) {
            break;
        }
        fast = fast.next;
        slow = slow.next;
        if (fast === slow) {
            return true;
        }
    }
    
    return false;
};