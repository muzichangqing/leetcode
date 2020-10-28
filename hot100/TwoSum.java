package hot100;

import java.util.HashMap;
import java.util.Map;

public class TwoSum {
    public int[] twoSum1(int[] nums, int target) {
        int[] indexs = {0, 0};

        for(int i = 0; i < nums.length - 1; i++) {
            for(int j = i + 1; j < nums.length; j++) {
                if (nums[i] + nums[j] == target) {
                    indexs[0] = i;
                    indexs[1] = j;
                    break;
                }
            }
        }
        return indexs;
    }

    public int[] twoSum2(int[] nums, int target) {
        Map<Integer, Integer> hashTable = new HashMap<Integer, Integer>();

        for (int i=0; i < nums.length; i++) {
            if (hashTable.containsKey(target - nums[i])) {
                return new int[]{i, hashTable.get(target - nums[i])};
            }
            hashTable.put(nums[i], i);
        }
        return new int[0];
    }
}