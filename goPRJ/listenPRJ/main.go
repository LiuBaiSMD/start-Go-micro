package main
import (
	"fmt"
	"log"
	"net"
	"bufio"
	"io"
)
func handleConnection(conn net.Conn) {
	br := bufio.NewReader(conn)
	fmt.Println("------>")
	for{
		data, err := br.ReadString('\n')
		if err == io.EOF{
			break
		}
		fmt.Printf("%s", data)
		fmt.Fprintf(conn, "OK\n")
	}
	conn.Close()
}
func main() {
	ln, err := net.Listen("tcp", ":8092")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("get client connection error: ", err)
		}
		go handleConnection(conn)
	}
}
