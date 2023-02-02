package everyday

func decodeMessage(key string, message string) string {
	m := make([]byte, 26)
	for i, j := 0, 0; i < len(key); i++ {
		b := key[i]
		if b == ' ' || m[b-'a'] != 0 {
			continue
		}
		m[b-'a'] = 'a' + byte(j)
		j++
	}
	rs := make([]byte, len(message))
	for i := 0; i < len(message); i++ {
		b := message[i]
		if b == ' ' {
			rs[i] = b
		} else {
			rs[i] = m[b-'a']
		}
	}
	return string(rs)
}
