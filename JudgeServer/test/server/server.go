package main

import (
	"fmt"
	"github.com/afanke/OJO/utils/tcp"
)

var msg = make(chan string, 1)
var recvmsg = make(chan string, 1)

func handle(conn tcp.Conn) {
	defer conn.Close()
	go recv()
	for {
		select {
		case s := <-msg:
			_, err := conn.Send([]byte(s))
			if err != nil {
				fmt.Printf("error:%v\n", err)
				return
			}
		case s := <-recvmsg:
			fmt.Println(s)
		}
	}
}
func recv(conn tcp.Conn) {
	for {
		_, bytes, err := conn.Recv()
		if err != nil {
			fmt.Printf("error:%v\n", err)
			return
		}
		recvmsg <- string(bytes)
	}
}

func read() {
	var s string
	for {
		_, err := fmt.Scanln(&s)
		if err != nil {
			fmt.Printf("error:%v\n", err)
			return
		}
		msg <- s
	}
}

func main() {
	go read()
	listen, err := tcp.Listen(":8888")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	for {
		Conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("error:%v\n", err)
			return
		}
		go handle(Conn)
	}
}
