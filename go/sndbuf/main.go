package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"syscall"
)

func getSndbufSize(conn interface{}) (int, error) {
	var fd int
	switch v := conn.(type) {
	case *net.TCPConn:
		f, err := v.File()
		if err != nil {
			return 0, err
		}
		fd = int(f.Fd())

	case *net.TCPListener:
		f, err := v.File()
		if err != nil {
			return 0, err
		}
		fd = int(f.Fd())
	}
	v, err := syscall.GetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_SNDBUF)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func enobufer(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(getSndbufSize(conn))
	b := make([]byte, 1024)
	for i := 0; i < 1024; i++ {
		b[i] = 'a'
	}
	for {
		_, err = conn.Write(b)
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	if len(os.Args) != 1 && os.Args[1] != "" {
		log.Println(fmt.Sprintf("client mode for %s", os.Args[1]))
		enobufer(os.Args[1])
	}

	l, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	fmt.Println(getSndbufSize(l))
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		defer conn.Close()

		fmt.Println(getSndbufSize(conn))
	}
}
