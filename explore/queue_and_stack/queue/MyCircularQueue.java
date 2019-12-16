class MyCircularQueue {

    private int[] data;
    private int head;
    private int tail;
    private int size;
    private int length;

    /** Initialize your data structure here. Set the size of the queue to be k. */
    public MyCircularQueue(int k) {
        size = k;
        length = 0;
        data = new int[size];
        head = tail = -1;
    }
    
    /** Insert an element into the circular queue. Return true if the operation is successful. */
    public boolean enQueue(int value) {
        if (length == size) {
            return false;
        }
        if (head == -1) {
            data[0] = value;
            tail = head = 0;
        } else {
            tail = (tail + 1) % size;
            data[tail] = value;
        }
        length++;
        return true;
    }
    
    /** Delete an element from the circular queue. Return true if the operation is successful. */
    public boolean deQueue() {
        if (length == 0) {
            return false;
        }
        head = head + 1 % size;
        length--;
        if (length == 0) {
            head = tail = -1;
        }
        return true;
    }
    
    /** Get the front item from the queue. */
    public int Front() {
        if (length == 0) {
            return -1;
        }
        return data[head];
    }
    
    /** Get the last item from the queue. */
    public int Rear() {
        if (length == 0) {
            return -1;
        }
        return data[tail];
    }
    
    /** Checks whether the circular queue is empty or not. */
    public boolean isEmpty() {
        return length == 0;
    }
    
    /** Checks whether the circular queue is full or not. */
    public boolean isFull() {
        return length == size;
    }
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * MyCircularQueue obj = new MyCircularQueue(k);
 * boolean param_1 = obj.enQueue(value);
 * boolean param_2 = obj.deQueue();
 * int param_3 = obj.Front();
 * int param_4 = obj.Rear();
 * boolean param_5 = obj.isEmpty();
 * boolean param_6 = obj.isFull();
 */