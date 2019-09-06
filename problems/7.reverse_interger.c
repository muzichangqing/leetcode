int reverse(int x) {
    long result = 0;
    unsigned max = INT_MAX + 1;
    int sign = x > 0 ? 1 : -1;
    x = x > 0 ? x : -x;
    while (x > 0) {
        result = result * 10 + x % 10;
        x /= 10;
        if (sign == 1 && result > max - 1) {
            return 0;
        } else if (sign == -1 && result > max) {
            return 0;
        }
    }
    return result * sign;
}
