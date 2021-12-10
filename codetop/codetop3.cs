using System;

namespace codetop
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello World!");
            var left = new Node(2, new Node(1), new Node(3));
            var right = new Node(5);
            var root = new Node(4, left, right);

            var solution = new Solution();
            var doubleList = solution.TreeToDoublyList(root);

            var cursor = doubleList;
            if (cursor != null) {
                do
                {
                    Console.Write(cursor.val);
                    Console.Write("\t");
                    cursor = cursor.left;
                } while (cursor != doubleList && cursor != null);
                Console.WriteLine();
                cursor = doubleList;
                do
                {
                    Console.Write(cursor.val);
                    Console.Write("\t");
                    cursor = cursor.right;
                }
                while (cursor != doubleList && cursor != null);
                Console.WriteLine();
            } else
            {
                Console.WriteLine("root is null");
            }
        }
    }

    public class Node
    {
        public int val;
        public Node left;
        public Node right;

        public Node() { }

        public Node(int _val)
        {
            val = _val;
            left = null;
            right = null;
        }

        public Node(int _val, Node _left, Node _right)
        {
            val = _val;
            left = _left;
            right = _right;
        }
    }

    public class Solution
    {
        // 剑指 Offer 36. 二叉搜索树与双向链表
        public Node TreeToDoublyList(Node root)
        {
            if (root == null)
            {
                return null;
            }
            var left = TreeToDoublyList(root.left);
            var right = TreeToDoublyList(root.right);

            var head = root;

            head.left = head;
            head.right = head;

            if (left != null)
            {
                head = left;
                var leftTail = left.left;
                head.left = root;
                root.right = head;
                leftTail.right = root;
                root.left = leftTail;
            }
            if (right != null)
            {
                var tail = head.left;
                var rightTail = right.left;
                head.left = rightTail;
                rightTail.right = head;
                tail.right = right;
                right.left = tail;
            }

            return head;
        }
    }
}
