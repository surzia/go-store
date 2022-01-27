package set

import "testing"

var (
	t1 T = 1
	t2 T = 2
	t3 T = 3
)

func InitSet() *Set {
	set := &Set{}
	set.Add(t1)
	set.Add(t2)
	return set
}

func TestSet_Add(t *testing.T) {
	set := InitSet()
	ok := set.Add(t1)
	if ok {
		t.Errorf("There is already %d in set!", t1)
	}

	ok = set.Add(t3)
	if !ok {
		t.Errorf("There should be %d in set!", t3)
	}
}

func TestSet_Clear(t *testing.T) {
	set := InitSet()
	set.Clear()
	if size := set.Size(); size != 0 {
		t.Errorf("wrong count, expected 0 and got %d", size)
	}
}

func TestSet_Delete(t *testing.T) {
	set := InitSet()
	ok := set.Delete(t1)
	if !ok {
		t.Errorf("There should be %d in set!", t1)
	}

	ok = set.Delete(t3)
	if ok {
		t.Errorf("There should not be %d in set!", t3)
	}
}

func TestSet_Contains(t *testing.T) {
	set := InitSet()
	ok := set.Contains(t1)
	if !ok {
		t.Errorf("There should be %d in set!", t1)
	}

	ok = set.Contains(t2)
	if !ok {
		t.Errorf("There should be %d in set!", t2)
	}

	ok = set.Contains(t3)
	if ok {
		t.Errorf("There should not be %d in set!", t3)
	}
}

func TestSet_All(t *testing.T) {
	set := InitSet()
	items := set.All()
	if len(items) != 2 {
		t.Errorf("wrong count, expected 2 and got %d", len(items))
	}

	if items[0] != t1 && items[1] != t2 {
		t.Errorf("There should be %d and %d in set!", t1, t2)
	}
}

func TestSet_Size(t *testing.T) {
	set := InitSet()
	size := set.Size()
	if size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}

	set.Add(t3)
	size = set.Size()
	if size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}

	set.Delete(t3)
	size = set.Size()
	if size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}

	set.Delete(t2)
	size = set.Size()
	if size != 1 {
		t.Errorf("wrong count, expected 1 and got %d", size)
	}

	set.Delete(t1)
	size = set.Size()
	if size != 0 {
		t.Errorf("wrong count, expected 0 and got %d", size)
	}
}
