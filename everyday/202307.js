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

/**
 * @param {ListNode} head
 * @return {return}
 */
var detectCycle = function(head) {
    if (head === null) {
        return null;
    }
    let slow = head, fast = head;
    while (fast !== null) {
        slow = slow.next;
        if (fast.next !== null) {
            fast = fast.next.next;
        } else {
            return null;
        }
        if (fast === slow) {
            let ptr = head;
            while (ptr !== slow) {
                ptr = ptr.next;
                slow = slow.next;
            }
            return ptr;
        }
    }
    return null;
};