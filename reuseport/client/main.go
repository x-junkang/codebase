package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
	n, err := conn.Write([]byte("hello"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("write %d bytes\n", n)
	buf := make([]byte, 16)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("read %d bytes, data: %s\n", n, buf[:n])
}
