import java.util.HashSet;
import java.util.LinkedList;
import java.util.Queue;
import java.util.Set;

class WallsAndGatesSolution {
    public void wallsAndGates(int[][] rooms) {

        if (rooms.length < 1) {
            return;
        }

        int m,n;
        m = rooms[0].length;
        n = rooms.length;

        for(int i = 0; i < n; i++) {
            for(int j = 0; j < m; j++) {
                int val = rooms[i][j];
                if (val == 0) {
                    bfs(rooms, i, j);
                }
            }
        }
    }

    private static void bfs(int[][] rooms, int x, int y) {
        Queue<Coordinate> queue = new LinkedList<Coordinate>();
        Set<Coordinate> used = new HashSet<Coordinate>();

        int maxX = rooms.length - 1;
        int maxY = rooms[0].length - 1;
        
        Coordinate rootCo = new Coordinate(x, y);
        queue.offer(rootCo);
        used.add(rootCo);

        int step = 0;
        while(!queue.isEmpty()) {
            step++;

            int currentSize = queue.size();

            for(int i = 0; i < currentSize; i++) {
                Coordinate nodeCo = queue.peek();
                
                if (nodeCo.y > 0) {
                    Coordinate uCo = new Coordinate(nodeCo.x, nodeCo.y - 1);
                    if (rooms[uCo.x][uCo.y] > 0 && !used.contains(uCo)) {
                        used.add(uCo);
                        if (rooms[uCo.x][uCo.y] > step) {
                            queue.offer(uCo);
                            rooms[uCo.x][uCo.y] = step;
                        }
                    }
                }
    
                if (nodeCo.y < maxY) {
                    Coordinate dCo = new Coordinate(nodeCo.x, nodeCo.y + 1);
                    if (rooms[dCo.x][dCo.y] > 0 && !used.contains(dCo)) {
                        used.add(dCo);
                        if (rooms[dCo.x][dCo.y] > step) {
                            queue.offer(dCo);
                            rooms[dCo.x][dCo.y] = step;
                        }
                    }
                }
    
                if (nodeCo.x > 0) {
                    Coordinate lCo = new Coordinate(nodeCo.x - 1, nodeCo.y);
                    if (rooms[lCo.x][lCo.y] > 0 && !used.contains(lCo)) {
                        used.add(lCo);
                        if (rooms[lCo.x][lCo.y] > step) {
                            queue.offer(lCo);
                            rooms[lCo.x][lCo.y] = step;
                        }
                    }
                }
    
                if (nodeCo.x < maxX) {
                    Coordinate rCo = new Coordinate(nodeCo.x + 1, nodeCo.y);
                    if (rooms[rCo.x][rCo.y] > 0 && !used.contains(rCo)) {
                        used.add(rCo);
                        if (rooms[rCo.x][rCo.y] > step) {
                            queue.offer(rCo);
                            rooms[rCo.x][rCo.y] = step;
                        }
                    }
                }
                queue.poll();
            }
        }
    }
}

class Coordinate {
    public int x;
    public int y; 

    public Coordinate(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public boolean equals(Object obj) {
        if (!(obj instanceof Coordinate)) {
            return false;
        }
        if (obj == this) {
            return true;
        }
        return x == ((Coordinate) obj).x && y == ((Coordinate) obj).y;
    }
}