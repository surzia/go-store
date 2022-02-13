package hashmap

import "testing"

func TestHash(t *testing.T) {
	var key Key = "testKey"
	hashKey := hash(key)
	if hashKey != 105951707245 {
		t.Errorf("wrong hash, expect 105951707245 but got %d", hashKey)
	}
}
