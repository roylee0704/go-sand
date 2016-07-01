package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:4444")
	if err != nil {
		log.Fatal("error in listening")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error in accepting")
		}
		go handleConn(conn) // one connection per goroutine
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		go echo(c, scanner.Text())
	}
}

func echo(c net.Conn, shout string) {
	fmt.Fprintln(c, strings.ToUpper(shout))
	time.Sleep(10 * time.Second)
	fmt.Fprintln(c, shout)
	time.Sleep(10 * time.Second)
	fmt.Fprintln(c, strings.ToLower(shout))
}
