package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/hoangduc02011998/golang-rpc/server/action"
)

// định nghĩa service struct
type HelloService struct{}

func (p *HelloService) Hello(req string, reply *string) error {
	*reply = "hello " + req

	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	rpc.RegisterName("KVStoreService", action.NewKVStoreService())

	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal("Listen TCP error: ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept error: ", err)
		}

		go rpc.ServeConn(conn)
	}
}
