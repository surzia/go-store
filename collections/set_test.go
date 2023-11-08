package collections

import "testing"

func TestSet_Add(t *testing.T) {
	set := NewSet()

	set.Add(1)
	set.Add(2)
	set.Add(3)

	if set.Size() != 3 {
		t.Errorf("Expected set size: 3, got: %d", set.Size())
	}
}

func TestSet_Remove(t *testing.T) {
	set := NewSet()

	set.Add(1)
	set.Add(2)
	set.Add(3)

	set.Remove(2)

	if set.Size() != 2 {
		t.Errorf("Expected set size: 2, got: %d", set.Size())
	}
	if set.Contains(2) {
		t.Error("Expected set not to contain 2, but it does")
	}
}

func TestSet_Contains(t *testing.T) {
	set := NewSet()

	set.Add("apple")
	set.Add("orange")

	if !set.Contains("apple") {
		t.Error("Expected set to contain apple, but it doesn't")
	}
	if set.Contains("banana") {
		t.Error("Expected set not to contain banana, but it does")
	}
}

func TestSet_Size(t *testing.T) {
	set := NewSet()

	set.Add(1)
	set.Add(2)
	set.Add(3)

	if set.Size() != 3 {
		t.Errorf("Expected set size: 3, got: %d", set.Size())
	}
}
