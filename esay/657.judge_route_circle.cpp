class Solution {
public:
    bool judgeCircle(string moves) {
        int index = 0;
        char c = moves[index];
        int count_u = 0, count_d = 0, count_r = 0, count_l = 0;
        while (c != '\0') {
            if (c == 'U') count_u++;
            if (c == 'D') count_d++;
            if (c == 'L') count_l++;
            if (c == 'R') count_r++;
            c = moves[++index];
        }
        if (count_d == count_u && count_l == count_r) return true;
        return false;
    }
};
