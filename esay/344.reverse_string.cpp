class Solution {
public:
    string reverseString(string s) {
        size_t size = s.length();
        string rs(size, ' ');
        for (int i = 0; i < size; i++) {
            rs[i] = s[size - 1 - i];
        }
        return rs;
    }
};
