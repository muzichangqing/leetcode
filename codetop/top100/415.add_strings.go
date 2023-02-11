package top100

func addStrings(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	z := m + 1
	if n > m {
		z = n + 1
	}
	ans := make([]byte, z)
	flag := byte(0)
	i, j, k := m-1, n-1, z-1
	for ; i >= 0 && j >= 0; i, j, k = i-1, j-1, k-1 {
		num := num1[i] + num2[j] - '0' - '0' + flag
		flag = num / 10
		ans[k] = num%10 + '0'
	}
	for ; i >= 0; i, k = i-1, k-1 {
		num := num1[i] - '0' + flag
		flag = num / 10
		ans[k] = num%10 + '0'
	}
	for ; j >= 0; j, k = j-1, k-1 {
		num := num2[j] - '0' + flag
		flag = num / 10
		ans[k] = num%10 + '0'
	}
	if flag != 0 {
		ans[k] = '1'
	}
	if ans[0] == 0 {
		ans = ans[1:]
	}
	return string(ans)
}
