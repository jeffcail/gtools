package main

import (
	"fmt"
	"github.com/jeffcail/gtools"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	defer conn.Close()
	if err != nil {
		fmt.Printf("Connect tcp err: %v", err)
		return
	}

	err = gtools.Encode(conn, "Hi mary!!!")
	if err != nil {
		fmt.Printf("Unpack err: %v", err)
		return
	}
}
