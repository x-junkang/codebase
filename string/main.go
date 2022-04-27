package main

import (
	"fmt"
	"strings"
)

func MapToKVString(data map[string]string) string {
	if data == nil {
		return ""
	}
	res := strings.Builder{}
	for k, v := range data {
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

func main() {
	data := map[string]string{
		"hello": "world",
		"test":  "value",
	}
	s := MapToKVString(data)
	fmt.Println("data: ", s)
	data2 := map[string]string{}
	s2 := MapToKVString(data2)
	fmt.Println("data: ", s2)
}
