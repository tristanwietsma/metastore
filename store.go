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

import "sync"

// Store is the base storage structure.
type Store struct {
	dataMap map[string]string
	subMap  map[string][]chan<- string
	sync.RWMutex
}

// Init initializes the Store.
func (s *Store) Init() {
	s.dataMap = make(map[string]string)
	s.subMap = make(map[string][]chan<- string)
}

// Get returns the value associated with a key; also returns a boolean indicating success or failure.
func (s *Store) Get(key string) (value string, ok bool) {
	s.RLock()
	defer s.RUnlock()
	value, ok = s.dataMap[key]
	return
}

// Set associates a key with a value.
func (s *Store) Set(key, value string) {
	s.Lock()
	defer s.Unlock()
	s.dataMap[key] = value
}

// Delete removes a key from storage.
func (s *Store) Delete(key string) {
	s.Lock()
	defer s.Unlock()
	delete(s.dataMap, key)
}

// Publish associates a key with a value and updates subscribers.
func (s *Store) Publish(key, value string) {
	s.Set(key, value)
	subs, ok := s.fetchSubscribers(key)
	if ok {
		for _, out := range subs {
			defer func(o chan<- string) {
				if r := recover(); r != nil {
					s.unsubscribe(key, o)
				}
			}(out)
			out <- value
		}
	}
}

// Subscribe associates an alert on an outgoing channel with a key.
func (s *Store) Subscribe(key string, outgoing chan<- string) {
	_, hasSubs := s.fetchSubscribers(key)
	s.Lock()
	defer s.Unlock()
	if hasSubs {
		s.subMap[key] = append(s.subMap[key], outgoing)
	} else {
		subs := []chan<- string{outgoing}
		s.subMap[key] = subs
	}
}

func (s *Store) unsubscribe(key string, outgoing chan<- string) {
	subs, hasSubs := s.fetchSubscribers(key)
	s.Lock()
	defer s.Unlock()
	if hasSubs {
		newSubs := []chan<- string{}
		for _, sub := range subs {
			if sub == outgoing {
				continue
			}
			newSubs = append(newSubs, sub)
		}
		s.subMap[key] = newSubs
	}
}

func (s *Store) fetchSubscribers(key string) ([]chan<- string, bool) {
	s.RLock()
	subs, hasSubs := s.subMap[key]
	s.RUnlock()
	return subs, hasSubs
}
