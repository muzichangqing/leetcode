from typing import List, Optional, Tuple
import functools
import random

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        @functools.lru_cache(amount)
        def dp(rem) -> int:
            if rem < 0: return -1
            if rem == 0: return 0
            mini = int(1e9)
            for coin in coins:
                res = dp(rem - coin)
                if res >= 0 and res < mini:
                    mini = res + 1
            return mini if mini < int(1e9) else -1
        if amount < 1: return 0
        return dp(amount)
    
    def numberOfPairs(self, nums: List[int]) -> List[int]:
        count = 0
        m = set([])
        for num in nums:
            if num in m:
                m.remove(num)
                count += 1
            else:
                m.add(num)
        return [count, len(nums)-count*2]
    
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        n = len(matrix)
        for i in range((n+1)>>1):
            for j in range(i, n-1-i):
                i2 = n-i-1
                j2 = n-j-1
                matrix[i][j], matrix[j][i2], matrix[i2][j2], matrix[j2][i] = \
                matrix[j2][i], matrix[i][j], matrix[j][i2], matrix[i2][j2]

    def compareVersion(self, version1: str, version2: str) -> int:
        m, n = len(version1), len(version2)
        i = j = 0
        while i < m or j < n:
            x = 0
            while i < m and version1[i] != '.':
                x = x * 10 + ord(version1[i]) - ord('0')
                i += 1
            i += 1
            y = 0
            while j < n and version2[j] != '.':
                y = y * 10 + ord(version2[j]) - ord('0')
                j += 1
            j += 1
            if x != y:
                return 1 if x > y else -1
        return 0

    def preorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        ans = []
        def traversal(node: Optional[TreeNode]):
            if node is None:
                return
            ans.append(node.val)
            traversal(node.left)
            traversal(node.right)
        traversal(root)
        return ans

    def subsets(self, nums: List[int]) -> List[List[int]]:
        ans = []
        n = len(nums)

        def dfs(idx: int, items: List[int]):
            if idx == n:
                ans.append(items.copy())
                return

            items.append(nums[idx])
            dfs(idx+1, items)
            items.pop()
            dfs(idx+1, items)   
               
        dfs(0, [])
        return ans

    def sumNumbers(self, root: Optional[TreeNode]) -> int:
        ans = 0
        if root is None:
            return ans

        def travel(node: TreeNode, sum: int):
            nonlocal ans
            sum = sum * 10 + node.val
            if node.left is not None:
                travel(node.left, sum)
            if node.right is not None:
                travel(node.right, sum)
            if node.left is None and node.right is None:
                ans += sum
        
        travel(root, 0)
        return ans

    def maxDepth(self, root: Optional[TreeNode]) -> int:
        max_depth = 0
        
        def dfs(node: Optional[TreeNode], depth: int):
            nonlocal max_depth
            if node is None:
                max_depth = max(max_depth, depth)
                return
            dfs(node.left, depth+1)
            dfs(node.right, depth + 1)

        dfs(root, 0)
        return max_depth

    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        queue = [root]
        should_break = False
        while not should_break:
            n = len(queue)
            for i in range(n//2):
                l, r = queue[i], queue[n-1-i]
                if l is None and r is None:
                    continue
                if l is None or r is None:
                    return False
                if l.val != r.val:
                    return False
            new_queue = []
            should_break = True
            for node in queue:
                l, r = None, None
                if node is not None:
                    l, r = node.left, node.right
                if l is not None or r is not None:
                    should_break = False
                new_queue.append(l)
                new_queue.append(r)
            queue = new_queue
        return True

    def diameterOfBinaryTree(self,root: Optional[TreeNode]) -> int:
        diameter = 0
        def dfs(node: Optional[TreeNode]) -> int:
            nonlocal diameter
            if node is None:
                return 0
            l, r = dfs(node.left), dfs(node.right)
            diameter = max(diameter, l + r)
            return max(l, r) + 1
        dfs(root)
        return diameter
    
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        def dfs(node: Optional[TreeNode]) -> Tuple[bool, float, float]:
            if node is None:
                return (True, float('inf'), float('-inf'))
            left_is_valid, left_min, left_max = dfs(node.left)
            if not left_is_valid:
                return (False, 0, 0)
            right_is_valid, right_min, right_max = dfs(node.right)
            if not right_is_valid:
                return (False, 0, 0)
            if node.val <= left_max or node.val >= right_min:
                return (False, 0, 0)
            if left_min == float('inf'):
                left_min = node.val
            if right_max == float('-inf'):
                right_max = node.val
            return (True, left_min, right_max)
        is_valid, _, _ = dfs(root)
        return is_valid 

    def minPathSum(self, grid: List[List[int]]) -> int:
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if i == j == 0: continue
                elif i == 0: grid[i][j] += grid[i][j-1]
                elif j == 0: grid[i][j] += grid[i-1][j]
                else: grid[i][j] += min(grid[i-1][j], grid[i][j-1])
        return grid[-1][-1]

    def rand10(self):
        x, y, n = 0, 0, 100
        while n > 40:
            x, y = rand7(), rand7()
            n = x + (y-1) * 7
        return 1 + (n-1) % 10

def rand7():
    return random.randint(1, 7)


if __name__ == '__main__':
    s = Solution()
    nums = [1, 2, 3]
    print(s.subsets(nums))
