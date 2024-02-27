package utils

import (
	"fmt"
	"log"
	"net"
	"sync"
)

var (
	Port = "8989"
)

func Server() {
	listen, err := net.Listen("tcp", ""+":"+Port)
	if err != nil {
		log.Fatal(err)
	} else if err == nil {
		fmt.Println("Now listening to port: " + Port)
	}

	var mutex sync.Mutex

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		if NbrClents == 10 {
			conn.Write([]byte("Chat is close try latter ..."))
			conn.Close()
		}
		NbrClents++
		go ClientHandler(conn, &mutex)
	}

}
