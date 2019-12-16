import java.util.LinkedList;
import java.util.Queue;

class MovingAverage {

    private Queue<Integer> window;
    private int windowSize;
    private double sum;

    /** Initialize your data structure here. */
    public MovingAverage(int size) {
        window = new LinkedList<Integer>();
        windowSize = size;
        sum = 0;
    }
    
    public double next(int val) {
        if (isFull()) {
            sum -= window.poll();
        }
        window.offer(val);
        sum += val;
        return sum / window.size();
    }

    private boolean isFull() {
        return window.size() == windowSize;
    }
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * MovingAverage obj = new MovingAverage(size);
 * double param_1 = obj.next(val);
 */