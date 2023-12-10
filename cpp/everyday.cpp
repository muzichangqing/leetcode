class Solution {
public:
    int climbStairs(int n) {
        int dp1 = 1;
        int dp2 = 2;
        if (n == 1) {
            return dp1;
        }
        if (n == 2) {
            return dp2;
        }
        for (int i = 3; i <= n; i++) {
            int tmp = dp1;
            dp1 = dp2;
            dp2 = dp1 + tmp;
        }
        return dp2;
    }
};