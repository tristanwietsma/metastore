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

import "testing"

// TestGet will set a key and get back the value
func TestGet(t *testing.T) {
	var S Store
	S.Init()
	S.Set("key123", "value567")
	value, ok := S.Get("key123")
	if value != "value567" || !ok {
		t.Errorf("Set-Get failed.")
	}
}

// TestDelete will set a key, delete it, and verify it is gone
func TestDelete(t *testing.T) {
	var S Store
	S.Init()
	S.Set("key123", "value567")
	S.Delete("key123")
	value, ok := S.Get("key123")
	if ok {
		t.Errorf("Delete failed. Got back value '%s'.", value)
	}
}

// TestSubscribe will subscribe and publish on a key
func TestSubscribe(t *testing.T) {
	var S Store
	S.Init()
	recv := make(chan string)
	S.Subscribe("key123", recv)
	go S.Publish("key123", "value567")
	for {
		value := <-recv
		if value != "value567" {
			t.Errorf("Publish-Subscribe failed. Got back value '%s'.", value)
		}
		return
	}
}
