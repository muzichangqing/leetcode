# -*- coding:utf-8 -*-

from typing import List, Optional


# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution(object):

    # 543. 二叉树的直径
    def diameterOfBinaryTree(self, root: TreeNode) -> int:
        """
        :type root: TreeNode
        :rtype: int
        """
        self.maxDiameter = 0

        def dfs(node: TreeNode) -> int:
            if not node:
                return 0
            l = dfs(node.left)
            r = dfs(node.right)
            self.maxDiameter = max(self.maxDiameter, l + r + 1)
            return max(l, r) + 1

        dfs(root)
        return self.maxDiameter - 1

    # 41. 缺失的第一个正数
    def firstMissingPositive(self, nums: List[int]) -> int:
        if 1 not in nums:
            return 1
        length = len(nums)
        for i in range(length):
            if nums[i] <= 0 or nums[i] > length:
                nums[i] = 1
        for i in range(length):
            num = abs(nums[i]) - 1
            nums[num] = -1 * abs(nums[num])
        for i in range(length):
            if nums[i] > 0:
                return i + 1
        return length + 1

    # 98. 验证二叉搜索树
    def isValidBST(self, root: TreeNode) -> bool:
        stack, inorder = [], float('-inf')
        while stack or root:
            while root:
                stack.append(root)
                root = root.left
            root = stack.pop()
            if root.val <= inorder:
                return False
            inorder = root.val
            root = root.right
        return True

    # 718. 最长重复子数组
    def findLength(self, nums1: List[int], nums2: List[int]) -> int:
        m, n = len(nums1), len(nums2)
        dp = [[0] * (n+1) for _ in range (m+1)]
        ans = 0
        for i in range(m-1, -1, -1):
            for j in range(n-1, -1, -1):
                if nums1[i] == nums2[j]:
                    dp[i][j] = dp[i+1][j+1] + 1
                    ans = max(dp[i][j], ans)
        return ans

    # 234. 回文链表
    def isPalindrome(self, head: ListNode) -> bool:
        nums = []
        cursor = head
        while cursor:
            nums.append(cursor.val)
            cursor = cursor.next
        length = len(nums)
        for i in range(int(length/2)):
            if nums[i] != nums[length - i - 1]:
                return False
        return True

    # 78. 子集
    def subsets(self, nums: List[int]) -> List[List[int]]:
        self.res = [[]]
        def backtrace(nums: List[int], path: List[int], index: int):
            if index == len(nums):
                return
            path.append(nums[index])
            self.res.append(path[:])
            backtrace(nums, path, index+1)
            path.pop()
            backtrace(nums, path, index+1)
        backtrace(nums, [], 0)
        return self.res

    # 112. 路径总和
    def hasPathSum(self, root: Optional[TreeNode], targetSum: int) -> bool:
        if not root:
            return False
        targetSum -= root.val
        if not root.left and not root.right:
            return targetSum == 0
        return self.hasPathSum(root.left, targetSum) or self.hasPathSum(root.right, targetSum)

    # 322. 零钱兑换
    def coinChange(self, coins: List[int], amount: int) -> int:
        dp = [float('inf')] * (amount + 1)
        dp[0] = 0
        for coin in coins:
            for i in range(coin, amount+1):
                dp[i] = min(dp[i], dp[i-coin] + 1)
        return dp[amount] if dp[amount] != float('inf') else -1

    # 32. 最长有效括号
    def longestValidParentheses(self, s: str) -> int: 
        s_len = len(s)
        stack = [-1] * (s_len+1)
        top = 0
        max_length = 0
        for i in range(s_len):
            if s[i] == '(':
                top += 1
                stack[top] = i
            else:
                if stack[top] >= 0 and s[stack[top]] == '(':
                    top -= 1
                    max_length = max(max_length, i - stack[top])
                else:
                    top = 0
                    stack[top] = i
        return max_length
    
    # 48. 旋转图像
    def rotate(self, matrix: List[List[int]]) -> None:
        n = len(matrix)
        for i in range(n // 2):
            for j in range((n + 1) // 2):
                matrix[i][j], matrix[n - j - 1][i], matrix[n - i - 1][n - j - 1], matrix[j][n - i - 1] \
                    = matrix[n - j - 1][i], matrix[n - i - 1][n - j - 1], matrix[j][n - i - 1], matrix[i][j]
        print(matrix)


if __name__ == '__main__':
    solution = Solution()
    print(solution.longestValidParentheses(")("))
    solution.rotate([[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]])