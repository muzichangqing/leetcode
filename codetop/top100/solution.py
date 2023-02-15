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

if __name__ == '__main__':
    s = Solution()
    coins = [1, 2, 5]
    amount = 11
    print(s.coinChange(coins, amount))