/*
Copyright 2013 Tristan Wietsma

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metastore

import (
	"hash/fnv"
	"sync"
)

type MetaStore struct {
	size   uint
	Bucket []Store
	sync.RWMutex
}

func (m *MetaStore) Init(n uint) {
	m.size = n
	m.Bucket = make([]Store, n)
	var i uint
	for i = 0; i < n; i++ {
		m.Bucket[i].Init()
	}
}

func (m *MetaStore) GetHasher() func([]byte) uint {
	h := fnv.New32()
	hasher := func(kb []byte) uint {
		h.Write(kb)
		idx := uint(h.Sum32()) % m.size
		h.Reset()
		return idx
	}
	return hasher
}

func (m *MetaStore) Set(key, value string) {
	h := m.GetHasher()
	bucketId := h([]byte(key))
	m.Bucket[bucketId].Set(key, value)
}

func (m *MetaStore) Get(key string) (string, bool) {
	h := m.GetHasher()
	bucketId := h([]byte(key))
	return m.Bucket[bucketId].Get(key)
}

func (m *MetaStore) Delete(key string) {
	h := m.GetHasher()
	bucketId := h([]byte(key))
	m.Bucket[bucketId].Delete(key)
}

func (m *MetaStore) Publish(key, value string) {
	h := m.GetHasher()
	bucketId := h([]byte(key))
	m.Bucket[bucketId].Publish(key, value)
}

func (m *MetaStore) Subscribe(key string, outgoing chan<- string) {
	h := m.GetHasher()
	bucketId := h([]byte(key))
	m.Bucket[bucketId].Subscribe(key, outgoing)
}

func (m *MetaStore) Unsubscribe(key string, outgoing chan<- string) {
	h := m.GetHasher()
	bucketId := h([]byte(key))
	m.Bucket[bucketId].Unsubscribe(key, outgoing)
}

func (m *MetaStore) NumSubscribers(key string) int {
	h := m.GetHasher()
	bucketId := h([]byte(key))
	return m.Bucket[bucketId].NumSubscribers(key)
}

func (m *MetaStore) FlushAll() {
	m.Lock()
	defer m.Unlock()
	for i := range m.Bucket {
		m.Bucket[i].FlushAll()
	}
}
