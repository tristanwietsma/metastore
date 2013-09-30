package main

import (
	"fmt"
	"time"
	ms "github.com/tristanwietsma/metastore"
)

func main() {

	var S ms.Store
	S.Init()

	recv := make(chan string)

	go func() {
		for i := 0;; i++ {
			time.Sleep(time.Second)
			S.Publish("key123", fmt.Sprintf("value%d", i))
		}
	}()

	S.Subscribe("key123", recv)
	for {
		fmt.Println("key123 has value", <-recv)
	}

}
