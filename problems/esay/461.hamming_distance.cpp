class Solution {
public:
    int hammingDistance(int x, int y) {
        int tmp = x ^ y;
        int distance = 0;
        while (tmp != 0) {
            if (tmp % 2 == 1) distance ++;
            tmp /= 2;
        }
        return distance;
    }
};
