package top100

func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	if m > n {
		return multiply(num2, num1)
	}
	if num1 == "0" {
		return "0"
	}
	num2Bytes := reverseBytes([]byte(num2))
	ans := []byte{'0'}
	for i := m - 1; i >= 0; i-- {
		product := multiplyOneBit(num1[i], num2Bytes)
		newProduct := make([]byte, m-1-i, m-1-i+len(product))
		for i := range newProduct {
			newProduct[i] = '0'
		}
		newProduct = append(newProduct, product...)
		ans = add(ans, newProduct)
	}
	return string(reverseBytes(ans))
}

func reverseBytes(bs []byte) []byte {
	n := len(bs)
	for i := 0; i < n/2; i++ {
		bs[i], bs[n-1-i] = bs[n-1-i], bs[i]
	}
	return bs
}

func multiplyOneBit(num1 byte, num2 []byte) []byte {
	m := len(num2)
	ans := make([]byte, m+1)
	var carry byte
	num1 -= '0'
	for i, v := range num2 {
		product := num1*(v-'0') + carry
		ans[i] = product%10 + '0'
		carry = product / 10
	}
	ans[len(ans)-1] = carry + '0'
	if carry == 0 {
		return ans[:len(ans)-1]
	}
	return ans
}

func add(nums1, nums2 []byte) []byte {
	m, n := len(nums1), len(nums2)
	sum := make([]byte, max(m, n)+1)
	i, j := 0, 0
	var carry byte
	for ; i < m && i < n; i, j = i+1, j+1 {
		num := nums1[i] + nums2[j] + carry - '0'*2
		sum[j] = num%10 + '0'
		carry = num / 10
	}
	for ; i < m; i, j = i+1, j+1 {
		num := nums1[i] + carry - '0'
		sum[j] = num%10 + '0'
		carry = num / 10
	}
	for ; i < n; i, j = i+1, j+1 {
		num := nums2[i] + carry - '0'
		sum[j] = num%10 + '0'
		carry = num / 10
	}
	sum[j] = carry + '0'
	if carry == 0 {
		return sum[:j]
	}
	return sum
}
