[![Build Status](https://travis-ci.org/tristanwietsma/metastore.png?branch=master)](https://travis-ci.org/tristanwietsma/metastore)

MetaStore
=========

**Store** is an abstraction over a string map that supports get, set, delete, publish, and subscribe.

**MetaStore** ("a Store of Stores") is an abstraction over a Store that breaks the key space into buckets. By doing so, we get finer lock granularity when deployed in a concurrent environment, such as in [JackDB](https://github.com/tristanwietsma/jackdb).

Store
-----

A Store is, as mentioned, a wrapper over a string map. It provides safe access and tracks subscribers.

* Create a Store

* Set a value to a key

* Get a value at a key

* Delete a key

* Publish a value to a key

* Subscribe to changes on a key

MetaStore
---------

MetaStore splits the key space into buckets and provides a means to determine which bucket a key-value pair resides in.
