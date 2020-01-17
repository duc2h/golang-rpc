package action

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func DoClientWork(client *rpc.Client) {
	go func() {
		var keyChanged string

		err := client.Call("KVStoreService.Watch", 30, &keyChanged)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Watch: ", keyChanged)
	}()

	err := client.Call("KVStoreService.Set", [2]string{"abc", "abc-value"}, new(struct{}))

	err = client.Call("KVStoreService.Set", [2]string{"abc", "other"}, new(struct{}))

	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 30)
}
