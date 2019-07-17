class Solution {
public:
    string reverseString(string s) {
        size_t size = s.length();
        string rs;
        for (int i = 0; i < size; i++) {
            rs.push_back(s[size - 1 - i]);
        }
        return rs;
    }
};
