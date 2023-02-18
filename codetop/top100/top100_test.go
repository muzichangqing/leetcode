package top100

import "testing"

func TestRestoreIPAddresses(t *testing.T) {
	s := "25525511135"
	ips := restoreIpAddresses(s)
	t.Log(ips)
}

func TestAdd(t *testing.T) {
	t.Log(string(add([]byte("8631"), []byte("0219"))))
	t.Log(string(reverseBytes([]byte("12345"))))
	t.Log(string(multiplyOneBit('2', []byte{'3'})))
	t.Log(multiply("123", "456"))
}

func TestCombinationSum(t *testing.T) {
	t.Log(combinationSum([]int{2, 3, 6, 7}, 7))
}
