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

import "hash/fnv"

type MetaStore struct {
	size   uint
	Bucket []Store
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
