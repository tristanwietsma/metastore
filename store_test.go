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

func TestGet(t *testing.T) {
	var S Store
	S.Init()
	S.Set("key123", "value567")
	value, ok := S.Get("key123")
	if value != "value567" || !ok {
		t.Errorf("Get failed.")
	}
}

