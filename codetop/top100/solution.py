from typing import List, Optional
import functools

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

if __name__ == '__main__':
    s = Solution()
    nums = [1, 2, 3]
    print(s.subsets(nums))
