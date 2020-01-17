package main

import (
	"log"
	"net/rpc"

	"github.com/hoangduc02011998/golang-rpc/client/action"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("diating: ", err)
	}
	action.DoClientWork(client)

	// var reply string

	// err = client.Call("HelloService.Hello", "World 11", &reply)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(reply)
	
}
