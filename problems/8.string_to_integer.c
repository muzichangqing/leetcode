int myAtoi(char *str)
{
    int i = 0;
    int rs = 0;
    int int_32_max = 2147483647;
    int int_32_min = -2147483648;
    int int_32_d10 = 214748364;
    int flag = 1;
    while (str[i] == ' ')
    {
        i++;
    }
    if (str[i] == '+')
    {
        flag = 1;
        i++;
    }
    else if (str[i] == '-')
    {
        flag = -1;
        i++;
    }
    while (str[i] != '\0')
    {
        char c = str[i++];
        if (c < '0' || c > '9')
        {
            return flag < 0 ? 0 - rs : rs;
        }
        if (rs == int_32_d10 && c > '7' || rs > int_32_d10)
        {
            return flag < 0 ? int_32_min : int_32_max;
        }
        rs = (c - '0') + rs * 10; // 计算顺序不能反
    }
    return flag < 0 ? 0 - rs : rs;
}
