from typing import List
import functools


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


if __name__ == '__main__':
    s = Solution()
    coins = [1, 2, 5]
    amount = 11
    print(s.coinChange(coins, amount))

    nums = [1,3,2,1,3,2,2]
    print(s.numberOfPairs(nums))

    matrix = [[1,2,3],[4,5,6],[7,8,9]]
    s.rotate(matrix)
    print(matrix)