package set

type T int

type Set struct {
	sets map[T]bool
}

// Add adds a new element to the Set. Returns true if t is not in set.
func (s *Set) Add(t T) bool {
	if s.sets == nil {
		s.sets = make(map[T]bool)
	}

	_, ok := s.sets[t]
	if !ok {
		s.sets[t] = true
	}

	return !ok
}

// Clear removes all elements from the Set
func (s *Set) Clear() {
	s.sets = make(map[T]bool)
}

// Delete removes the Item from the Set and returns true if t is in set
func (s *Set) Delete(t T) bool {
	_, ok := s.sets[t]
	if ok {
		delete(s.sets, t)
	}

	return ok
}

// Contains returns true if the Set contains the t
func (s *Set) Contains(t T) bool {
	_, ok := s.sets[t]
	return ok
}

// All returns the all items stored
func (s *Set) All() []T {
	var ret []T
	for t := range s.sets {
		ret = append(ret, t)
	}

	return ret
}

// Size returns the size of the set
func (s *Set) Size() int {
	return len(s.sets)
}
