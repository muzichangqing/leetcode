package codetop

import "strings"

func validIPAddressV2(queryIP string) string {
	if strings.Contains(queryIP, ".") {
		if validIPv4(queryIP) {
			return "IPv4"
		}
		return "Neither"
	}
	if validIPv6(queryIP) {
		return "IPv6"
	}
	return "Neither"
}

func validIPv4(queryIP string) bool {
	segments := strings.Split(queryIP, ".")
	if len(segments) != 4 {
		return false
	}
	for _, segment := range segments {
		size := len(segment)
		switch {
		case size == 1:
			if segment[0] < '0' || segment[0] > '9' {
				return false
			}

		case size == 2:
			if segment[0] < '1' || segment[0] > '9' {
				return false
			}
			if segment[1] < '0' || segment[1] > '9' {
				return false
			}
		case size == 3:
			if segment[0] < '1' || segment[0] > '9' {
				return false
			}
			v := 0
			for i := 0; i < 3; i++ {
				v = v*10 + int(segment[i]-'0')
			}
			if v > 255 {
				return false
			}
		default:
			return false
		}
	}
	return true
}

func validIPv6(queryIP string) bool {
	segments := strings.Split(queryIP, ":")
	if len(segments) != 8 {
		return false
	}
	for _, segment := range segments {
		size := len(segment)
		if size < 1 || size > 4 {
			return false
		}
		for i := 0; i < size; i++ {
			b := segment[i]
			if b >= '0' && b <= '9' || b >= 'a' && b <= 'f' || b >= 'A' && b <= 'F' {
				continue
			}
			return false
		}
	}
	return true
}
