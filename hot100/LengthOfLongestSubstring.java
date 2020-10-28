package hot100;

import java.util.HashMap;
import java.util.Map;

public class LengthOfLongestSubstring {
    public int lengthOfLongestSubstring(String s) {
        int slow, fast, max;
        slow = fast = max = 0;
        Map<Character, Integer> hashTable = new HashMap<Character, Integer>();

        for (fast = 0; fast < s.length(); fast++) {
            char c = s.charAt(fast);
            if (hashTable.containsKey(c)) {
                int index = hashTable.get(c);
                if (index >= slow) {
                    if (fast - slow > max) {
                        max = fast - slow;
                    }
                    slow = hashTable.get(c) + 1;
                }
            }
            hashTable.put(c, fast);
        }
        if (fast - slow > max) {
            max = fast - slow;
        }
        return max;
    }
}
