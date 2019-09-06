public class Solution
{
    public int RomanToInt(string s)
    {
        int len = s.Length;
        int[] values = new int[len];

        Dictionary<char, int> romanIntMap = new Dictionary<char, int>(){
            {'I', 1}, {'V', 5}, {'X', 10}, {'L', 50}, {'C', 100}, {'D', 500},
            {'M', 1000}
        };

        char[] char_arr = s.ToCharArray();
        for (int i = len - 1; i >= 0; i--) {
            char c = char_arr[i];
            int current_value = romanIntMap[c];
            int last_value = i + 1 < len ? values[i + 1] : -1;
            if (last_value > current_value) {
                current_value = 0 - current_value;
            }
            values[i] = current_value;
        }

        int result = 0;
        foreach(int value in values) {
            result += value;
        }

        return result;
    }
}