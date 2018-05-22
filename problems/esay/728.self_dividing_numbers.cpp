class Solution {
public:
    std::vector<int> selfDividingNumbers(int left, int right) {
        std::vector<int> result;
        for(int i=left; i<=right; i++) {
            int num = i;
            int tmp = i % 10;
            while(num != 0) {
                if (tmp == 0) {
                    break;
                }
                if (i % tmp != 0) {
                    break;
                }
                num /= 10;
                tmp = num % 10;
            }
            if (num == 0) {
                result.push_back(i);
            }
        }
        return result;
    }
};
