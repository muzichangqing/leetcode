class Solution {
public:
    string reverseWords(string s) {
        string tmp;
        string rs;
        int size = s.length();
        for (int pos = 0; pos <= size; pos++) {
            char c = s[pos];
            if (c == ' ' || c == '\0') {
                int tmp_size = tmp.length();
                for (int i = 0; i < tmp_size; i++) {
                    rs.push_back(tmp[tmp_size - 1 - i]);
                }
                rs.push_back(c);
                tmp.clear();
            } else {
                tmp.push_back(c);
            }
        }
        return rs;
    }
};
