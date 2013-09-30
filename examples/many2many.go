package main

import (
	"fmt"
	"time"
	ms "github.com/tristanwietsma/metastore"
)

func main() {

	var S ms.Store
	S.Init()

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	// create five publishers on the same key
	for pid := 0; pid < 5; pid++ {
		go func(pid int) {
			for i := 0;; i++ {
				time.Sleep(time.Second)
				val := fmt.Sprintf("%d_%d", pid,i)
				S.Publish("k123", val)
			}
		}(pid)
	}

	S.Subscribe("k123", c1)
	S.Subscribe("k123", c2)
	S.Subscribe("k123", c3)
	var v1, v2, v3 string
	for {
		select {
		case v1 = <-c1:
			fmt.Println("channel 1 got value:", v1)
		case v2 = <-c2:
			fmt.Println("channel 2 got value:", v2)
		case v3 = <-c3:
			fmt.Println("channel 3 got value:", v3)
		}
	}

}
