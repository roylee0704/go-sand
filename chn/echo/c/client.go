package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:4444")

	if err != nil {
		log.Fatal("error dialing to host")
	}
	defer conn.Close()

	go io.Copy(conn, os.Stdin)
	io.Copy(os.Stdout, conn)

}
