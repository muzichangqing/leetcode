import java.util.HashSet;
import java.util.LinkedList;
import java.util.Queue;
import java.util.Set;

class OpenLockSolution {
    public int openLock(String[] deadends, String target) {
        Set<String> deadendsSet = new HashSet<String>();
        for (String deadend : deadends) {
            deadendsSet.add(deadend);
        }
        String root = "0000";
        if (deadendsSet.contains(root)) {
            return -1;
        }
        Queue<String> queue = new LinkedList<String>();
        queue.offer(root);
        Set<String> used = new HashSet<String>();
        used.add(root);

        int step = 0;
        while(!queue.isEmpty()) {
            step++;
            for(int x = queue.size(); x > 0; x--) {
                String str = queue.poll();
                for(int i = 0; i < str.length(); i++) {
                    char[] arr = str.toCharArray();
                    char c = arr[i];
                    char cp = (char)(((int)c - 48 + 1 + 10) % 10 + 48);
                    arr[i] = cp;
                    String strp = String.valueOf(arr);
                    if(strp.equals(target)) {
                        return step; 
                    }
                    if (!deadendsSet.contains(strp) && !used.contains(strp)) {
                        queue.offer(strp);
                    }
                    used.add(strp);
    
                    char cs = (char)(((int)c - 48 - 1 + 10) % 10 + 48);
                    arr[i] = cs;
                    String strs = String.valueOf(arr);
                    if(strs.equals(target)) {
                        return step; 
                    }
                    if (!deadendsSet.contains(strs) && !used.contains(strs)) {
                        queue.offer(strs);
                    }
                    used.add(strs);
                }
            }
        }
        return -1;
    }
}