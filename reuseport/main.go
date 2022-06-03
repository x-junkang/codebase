package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"syscall"

	"golang.org/x/sys/unix"
)

var data = flag.String("data", "ok", "server response data")

func init() {
	flag.Parse()
	fmt.Println(*data)
}

func main() {
	cfg := net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) error {
			return c.Control(func(fd uintptr) {
				syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEADDR, 1)
				syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEPORT, 1)
			})
		},
	}
	tcp, err := cfg.Listen(context.Background(), "tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("listen failed", err)
		return
	}

	buf := make([]byte, 1024)
	for {
		conn, err := tcp.Accept()
		if err != nil {
			fmt.Println("accept failed", err)
			continue
		}
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("read failed", err)
				break
			}

			fmt.Println(string(buf[:n]))
			conn.Write([]byte(*data))
		}
	}
}
