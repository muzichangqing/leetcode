import java.util.HashSet;
import java.util.LinkedList;
import java.util.Queue;
import java.util.Set;

public class PerfectSquares {
    public int numSquares(int n) {
        Queue<Node> queue = new LinkedList<Node>();
        Set<Integer> set = new HashSet<Integer>();

        queue.offer(new Node(n, 1));
        set.add(n);
        while(!queue.isEmpty()) {
            Node node = queue.poll();;
            int val = node.val;
            int step = node.step;
            for(int i = 1; i * i <= val; i++) {
                int nextVal = val - i * i;
                if (nextVal == 0) {
                    return step;
                }
                if (!set.contains(nextVal)) {
                    queue.offer(new Node(nextVal, step+1));
                    set.add(nextVal);
                }
            }
        }
        return -1;
    }
}

class Node {
    public int val;
    public int step;

    public Node(int val, int step) {
        this.val = val;
        this.step = step;
    }
}