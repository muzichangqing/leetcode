import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

public class CloneGraphSolution {
    public Node cloneGraph(Node node) {
        Map<Node, Node> map = new HashMap<Node, Node>();
        return dfs(node, map);
    }

    private Node dfs(Node node, Map<Node, Node> map) {
        if (node == null) {
            return null;
        }
        if (map.get(node) != null) {
            return map.get(node);
        }
        Node cloneNode = new Node();
        cloneNode.val = node.val;
        cloneNode.neighbors = new LinkedList<Node>();
        map.put(node, cloneNode);
        for (Node neighbor : node.neighbors) {
            cloneNode.neighbors.add(dfs(neighbor, map));
        }
        return cloneNode;
    }
}


class Node {
    public int val;
    public List<Node> neighbors;

    public Node() {}

    public Node(int _val,List<Node> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
};