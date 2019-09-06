public class Solution {
    public string LongestCommonPrefix(string[] strs) 
    {
        if (strs.Length == 0) {
            return "";
        }
        string common_prefix = strs[0];
        char[] common_prefix_arr = common_prefix.ToCharArray();
        int common_prefix_len = common_prefix.Length;

        foreach(string str in strs) {
            if (common_prefix_len == 0) {
                break;
            }
            char[] str_arr = str.ToCharArray();
            common_prefix_len = common_prefix_len > str_arr.Length ? str_arr.Length : common_prefix_len;
            for (int i = 0; i < common_prefix_len; i++) {
                if (str_arr[i] != common_prefix[i]) {
                    common_prefix_len = i;
                    continue;
                }
            }
        }

        string result = "";
        for (int i = 0; i < common_prefix_len; i++) {
            result += common_prefix[i];
        }

        return result;
    }
}