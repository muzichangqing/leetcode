public class Solution {
    public int findTargetSumWays(int[] nums, int S) {
        return dfs(nums, 0, S, 1, 0) + dfs(nums, 0, S, 2, 0);
    }

    private int dfs(int[] nums, int index, int target, int opt, int sum) {
        if (index == nums.length - 1) {
            if (opt == 1 && sum + nums[index] == target) {
                return 1;
            }
            if (opt == 2 && sum - nums[index] == target) {
                return 1;
            }
            return 0;
        }
        if (opt == 1) {
            sum += nums[index];
        } else {
            sum -= nums[index];
        }
        return dfs(nums, index + 1, target, 1, sum) + dfs(nums, index + 1, target, 2, sum);
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        int[] nums = {1, 1, 1, 1, 1};
        System.out.println(s.findTargetSumWays(nums, 3));
    }
}