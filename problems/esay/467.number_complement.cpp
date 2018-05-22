class Solution {
public:
    int findComplement(int num) {
        int result = 0;
        int tmp;
        int index = 0;
        while(num != 0) {
            tmp = num % 2 == 1 ? 0 : 1;
            result += tmp * pow(2, index);
            num /= 2;
            index ++;
        }
        return result;
    }

    int pow(int num, int exp) {
        int result = 1;
        while(exp > 0) {
            result *= num;
            exp --;
        }
        return result;
    }
};
