package main

import (
	"fmt"
	ms "github.com/tristanwietsma/metastore"
)

func main() {

	// declare and initialize MetaStore
	// with 1000 buckets
	var M ms.MetaStore
	M.Init(1000)

	h := M.GetHasher()

	// Set
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		idx := h([]byte(key))
		value := fmt.Sprintf("value%d", i)
		M.Bucket[idx].Set(key, value)
	}

	// Get
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		idx := h([]byte(key))
		value, ok := M.Bucket[idx].Get(key)
		if !ok {
			panic("oops, womething wrong...")
		}
		fmt.Printf("%s has value %s\n", key, value)
	}

}
