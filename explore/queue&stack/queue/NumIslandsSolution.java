import java.util.HashSet;
import java.util.LinkedList;
import java.util.Queue;
import java.util.Set;

class NumIslandsSolution {
    public int numIslands(char[][] grid) {
        int num = 0;
        if (grid.length == 0) {
            return num;
        }
        int mx = grid.length;
        int my = grid[0].length;

        Set<Island> islandsSet = new HashSet<Island>();

        for (int i = 0; i < mx; i++) {
            for (int j = 0; j < my; j ++) {
                if (grid[i][j] == '1' && !islandsSet.contains(new Island(i, j))) {
                    num++;
                    bfs(grid, islandsSet, i, j);
                }
            }
        }
        return num;
    }

    private void bfs(char[][] grid, Set<Island> islandsSet, int sx, int sy) {
        Queue<Island> queue = new LinkedList<Island>();
        Island root = new Island(sx, sy);
        queue.offer(root);
        islandsSet.add(root);

        int mx = grid.length - 1;
        int my = grid[0].length - 1;
        while (!queue.isEmpty()) {
            Island node = queue.poll();
            int x = node.x;
            int y = node.y;
            if (y > 0 && grid[x][y - 1] == '1') {
                Island up = new Island(x, y - 1);
                if (!islandsSet.contains(up)) {
                    queue.offer(up);
                    islandsSet.add(up);
                }
            }
            if (y < my && grid[x][y + 1] == '1') {
                Island down = new Island(x, y + 1);
                if (!islandsSet.contains(down)) {
                    queue.offer(down);
                    islandsSet.add(down);
                }
            }
            if (x > 0 && grid[x - 1][y] == '1') {
                Island left = new Island(x - 1, y);
                if (!islandsSet.contains(left)) {
                    queue.offer(left);
                    islandsSet.add(left);
                }
            }
            if (x < mx && grid[x + 1][y] == '1') {
                Island right = new Island(x + 1, y);
                if (!islandsSet.contains(right)) {
                    queue.offer(right);
                    islandsSet.add(right);
                }
            }
        }

    }
}


class Island {
    public int x;
    public int y; 

    public Island(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public boolean equals(Object obj) {
        if (!(obj instanceof Island)) {
            return false;
        }
        if (obj == this) {
            return true;
        }
        return x == ((Island) obj).x && y == ((Island) obj).y;
    }

    public int hashCode() {
        return x * y;
    }
}