package hashmap

import "testing"

var (
	k1 Key   = "key1"
	k2 Key   = "key2"
	k3 Key   = "key3"
	v1 Value = "value1"
	v2 Value = "value2"
	v3 Value = "value3"
)

func InitHashMap() *HashMap {
	hashmap := NewHashMap()
	hashmap.hashmap[hash(k1)] = v1
	hashmap.hashmap[hash(k2)] = v2

	return hashmap
}

func TestHashMap_Put(t *testing.T) {
	hashmap := InitHashMap()
	if size := hashmap.Size(); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}

	v := hashmap.Get(k3)
	if v != nil {
		t.Errorf("key k3 is not in hashmap")
	}

	hashmap.Put(k3, v3)
	if size := hashmap.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
	v = hashmap.Get(k3)
	if *v != v3 {
		t.Errorf("key k3 is in hashmap")
	}
}

func TestHashMap_Remove(t *testing.T) {
	hashmap := InitHashMap()

	ret := hashmap.Remove(k1)
	if !ret {
		t.Errorf("removing key k1 should return true")
	}

	ret = hashmap.Remove(k3)
	if ret {
		t.Errorf("removing key k1 should return false")
	}
}
