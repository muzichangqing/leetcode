package top100

import "strconv"

func restoreIpAddresses(s string) []string {
	ips := make([]string, 0)
	ip := [4]int{}
	size := len(s)
	var backtrace func(int, int)
	backtrace = func(offset, ip_i int) {
		if ip_i == 4 {
			if offset == size {
				ip_s := ""
				for _, v := range ip {
					ip_s += strconv.Itoa(v)
					ip_s += "."
				}
				ips = append(ips, ip_s[:len(ip_s)-1])
			}
			return
		}
		if offset == size {
			return
		}
		if s[offset] == '0' {
			ip[ip_i] = 0
			backtrace(offset+1, ip_i+1)
			return
		}
		num := 0
		for i := offset; i < size; i++ {
			num = num*10 + int(s[i]-'0')
			if num > 0xFF {
				break
			}
			ip[ip_i] = num
			backtrace(i+1, ip_i+1)
		}
	}
	backtrace(0, 0)
	return ips
}
