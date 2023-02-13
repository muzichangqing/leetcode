package top100

import "testing"

func TestRestoreIPAddresses(t *testing.T) {
	s := "25525511135"
	ips := restoreIpAddresses(s)
	t.Log(ips)
}
