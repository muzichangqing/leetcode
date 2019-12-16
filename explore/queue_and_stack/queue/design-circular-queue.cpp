#include <vector>

using namespace std;

class MyCircularQueue {
protected:
    vector<int>* queue;
    int head;
    int tail;
    int size;
    int len;
public:
    /** Initialize your data structure here. Set the size of the queue to be k. */
    MyCircularQueue(int k) {
        queue = new vector<int>(k);
        head = tail = -1;
        size = k;
        len = 0;
    }
    
    /** Insert an element into the circular queue. Return true if the operation is successful. */
    bool enQueue(int value) {
        if (isFull()) {
            return false;
        } 
        if (head == -1) {
            head = tail = 0;
        } else {
            if (tail == size - 1) {
                tail = 0;
            } else {
                tail += 1;
            }
        }
        (*queue)[tail] = value;
        len += 1;
        return true;
    }
    
    /** Delete an element from the circular queue. Return true if the operation is successful. */
    bool deQueue() {
        if (isEmpty()) {
            return false;
        }
        if (head == size - 1) {
            head = 0;
        } else {
            head += 1;
        }
        len -= 1;
        if (len == 0) {
            head = tail = -1;
        }
        return true;
    }
    
    /** Get the front item from the queue. */
    int Front() {
        if (isEmpty()) return -1;
        return (*queue)[head];
    }
    
    /** Get the last item from the queue. */
    int Rear() {
        if (isEmpty()) return -1;
        return (*queue)[tail]; 
    }
    
    /** Checks whether the circular queue is empty or not. */
    bool isEmpty() {
        return len == 0;
    }
    
    /** Checks whether the circular queue is full or not. */
    bool isFull() {
        return len == size;
    }

    ~MyCircularQueue() {
        delete queue;
    }
};

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * MyCircularQueue* obj = new MyCircularQueue(k);
 * bool param_1 = obj->enQueue(value);
 * bool param_2 = obj->deQueue();
 * int param_3 = obj->Front();
 * int param_4 = obj->Rear();
 * bool param_5 = obj->isEmpty();
 * bool param_6 = obj->isFull();
 */