package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		io.WriteString(conn, "\n Hello from TCP server \n")
		fmt.Fprintln(conn, "How is your day ?")
		fmt.Fprintf(conn, "%v", "Well I hope!")
		conn.Close()

	}
}
