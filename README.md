[![Build Status](https://travis-ci.org/tristanwietsma/metastore.png?branch=master)](https://travis-ci.org/tristanwietsma/metastore)

MetaStore
=========

**Store** is an abstraction over a string map that supports get, set, delete, publish, and subscribe.

**MetaStore** ("a Store of Stores") is an abstraction over a Store that breaks the key space into buckets. By doing so, we get finer lock granularity when deployed in a concurrent environment, such as in [JackDB](https://github.com/tristanwietsma/jackdb).

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

Publish a value to a key:

    s.Publish(key, value)

Subscribe to changes on a key:

    recv := make(chan string)
    s.Subscribe(key, recv)
    for {
        fmt.Println(<-recv)
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
