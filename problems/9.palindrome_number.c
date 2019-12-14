#include <stdbool.h>

bool isPalindrome(int x) {
    if (x < 0) return false;
    if (x == 0) return true;
    int y = 0;
    int tmp = x;
    while (x != 0) {
         // 判断是否越界
        if (y == 214748364 && x % 10 > 7 || y > 214748364) {
            return false;
        }
        y = y * 10 + x % 10;
        x /= 10;

    }
    if (y == tmp) return true;
    else return false;
}
