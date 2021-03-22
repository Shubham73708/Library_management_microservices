package main

import (
	"bufio"
	"fmt"
	"net"
)

func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", message)
}

func main() {
	len, err := net.Listen("tcp", ":5050")
	check(err, "Welcome to the server :")

	for {
		conn, err := len.Accept()
		check(err, "Connection establish.")

		go func() {
			buf := bufio.NewReader(conn)
			for {
				name, err := buf.ReadString('\n')
				if err != nil {
					fmt.Printf("Welcome here.\n")
					break
				}

				conn.Write([]byte("Hi, " + name))
			}
		}()
	}
}
