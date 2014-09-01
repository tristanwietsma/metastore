MetaStore
=========

[![Build Status](https://travis-ci.org/tristanwietsma/metastore.png?branch=master)](https://travis-ci.org/tristanwietsma/metastore) [![GoDoc](https://godoc.org/github.com/tristanwietsma/metastore?status.svg)](http://godoc.org/github.com/tristanwietsma/metastore)


**Store** is an abstraction over a string map that supports get, set, delete, publish, and subscribe.

**MetaStore** is an abstraction over a Store that breaks the key space into buckets. By doing so, we get finer lock granularity when deployed in a concurrent environment, such as in [Jack](https://github.com/tristanwietsma/jack).

Store
-----

A Store is, as mentioned, a wrapper over a string map. It provides safe access and tracks subscribers. Here is an overview:

Create a Store:

    var s metastore.Store
    s.Init()

Set a value to a key:

    s.Set(key, value)

Get a value at a key:

    value, ok := s.Get(key)

Delete a key:

    s.Delete(key)

Delete everything (including subscriptions):

    s.FlushAll()

Publish a value to a key:

    s.Publish(key, value)

In terms of state change, Publish is equivalent to Set. However, Publish also updates subscribers with the new value, while Set does not.

Subscribe to changes on a key:

    recv := make(chan string)
    s.Subscribe(key, recv)
    for {
        fmt.Println(<-recv)
    }

As it stands, each subscription requires a unique channel for the receiver to know which key an update is associated with. This can introduce some overhead, but the alternative would be sending the key as well and parsing... I don't like parsing. Besides, if you really want to conserve resources, you can add an identifier to the value and do it yourself.

To unsubscribe, you can close the channel. Store will drop the channel from the subscriber list.

For finer control, call Unsubscribe:

    S.Unsubscribe(key, recv)

Want to know how many are subscriptions a key has?

    nInt := S.NumSubscribers(key)
    fmt.Printf("%s has %d subscribers\n", key, nInt)

Check out this many-to-many pub/sub pattern:

    var S ms.Store
    S.Init()
	
    // create five publishers on the same key
    for pid := 0; pid < 5; pid++ {
    	go func(pid int) {
			for i := 0; ; i++ {
				time.Sleep(time.Second)
				val := fmt.Sprintf("%d_%d", pid, i)
				S.Publish("k123", val)
			}
		}(pid)
	}

    c1 := make(chan string)
    c2 := make(chan string)
    c3 := make(chan string)
	
    S.Subscribe("k123", c1)
    S.Subscribe("k123", c2)
    S.Subscribe("k123", c3)
    var v string
    for {
    	select {
    	case v = <-c1:
    		fmt.Println("channel 1 got value:", v)
    	case v = <-c2:
    		fmt.Println("channel 2 got value:", v)
    	case v = <-c3:
    		fmt.Println("channel 3 got value:", v)
    	}
    }

MetaStore
---------

MetaStore splits the key space into buckets and provides a means to determine which bucket a key-value pair resides in.

Create a MetaStore and split the key space over 1000 buckets:

    var m metastore.MetaStore
    m.Init(1000)

Get a hash function to determine which bucket a key goes in:

    h := m.GetHasher()
    bucketId := h([]byte("the key"))

Set a value to a key:

    bucketId := h([]byte(key))
    m.Bucket[bucketId].Set(key, value)

MetaStore also implements all of Store's public methods, so you don't need to go to the trouble of calculating the bucket index. In other words, this is valid code, too:

    m.Set(key, value)
    value, ok := m.Get(key)
    m.Delete(key)
    m.Publish(key, value)
    m.Subscribe(key, recv)
    m.Unsubscribe(key, recv)
    nInt := m.NumSubscribers(key)
    m.FlushAll()
