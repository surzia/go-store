package collections

import "testing"

func TestHashTable_SetGet(t *testing.T) {
	hashTable := NewHashTable()

	hashTable.Set("one", 1)
	hashTable.Set("two", 2)
	hashTable.Set("three", 3)

	val, exists := hashTable.Get("two")
	if !exists || val != 2 {
		t.Error("Get did not return expected value")
	}
}

func TestHashTable_Delete(t *testing.T) {
	hashTable := NewHashTable()

	hashTable.Set("one", 1)
	hashTable.Set("two", 2)

	hashTable.Delete("two")

	_, exists := hashTable.Get("two")
	if exists {
		t.Error("Expected 'two' key to be deleted, but it still exists")
	}
}

func TestHashTable_Exists(t *testing.T) {
	hashTable := NewHashTable()

	hashTable.Set("apple", "red")
	hashTable.Set("banana", "yellow")

	if !hashTable.Exists("apple") {
		t.Error("Expected 'apple' key to exist, but it doesn't")
	}
	if hashTable.Exists("grape") {
		t.Error("Expected 'grape' key not to exist, but it does")
	}
}
