package utils

import "strings"

func MapToKVString(data map[string]string) string {
	if data == nil {
		return ""
	}
	res := strings.Builder{}
	for k, v := range data {
		if k == "" || v == "" {
			continue
		}
		res.WriteString(k)
		res.WriteByte('=')
		res.WriteString(v)
		res.WriteByte('&')
	}
	ans := res.String()
	if len(ans) > 0 {
		return ans[:len(ans)-1]
	}
	return ""
}
