class Solution {
public:
    int numJewelsInStones(string J, string S) {
        int len_j = J.length();
        int len_s = S.length();
        int result = 0;

        for(int i=0; i<len_s; i++) {
            for(int j=0; j<len_j; j++) {
                if (S[i] == J[j]) {
                    result++;
                }
            }
        }
        return result;
    }
};
