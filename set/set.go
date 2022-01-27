package set

import (
	"sync"
)

type T int

type Set struct {
	sync.RWMutex
	sets map[T]bool
}

func NewSet() *Set {
	ret := &Set{}
	ret.sets = make(map[T]bool)

	return ret
}

// Add adds a new element to the Set. Returns true if t is not in set.
func (s *Set) Add(t T) bool {
	s.Lock()
	defer s.Unlock()

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
	s.Lock()
	defer s.Unlock()

	s.sets = make(map[T]bool)
}

// Delete removes the Item from the Set and returns true if t is in set
func (s *Set) Delete(t T) bool {
	s.Lock()
	defer s.Unlock()

	_, ok := s.sets[t]
	if ok {
		delete(s.sets, t)
	}

	return ok
}

// Contains returns true if the Set contains the t
func (s *Set) Contains(t T) bool {
	s.RLock()
	defer s.RUnlock()

	_, ok := s.sets[t]
	return ok
}

// All returns the all items stored
func (s *Set) All() []T {
	s.RLock()
	defer s.RUnlock()

	var ret []T
	for t := range s.sets {
		ret = append(ret, t)
	}

	return ret
}

// Size returns the size of the set
func (s *Set) Size() int {
	s.RLock()
	defer s.RUnlock()

	return len(s.sets)
}

// Union returns a new set with elements from both
// the given sets
func (s *Set) Union(t *Set) *Set {
	ret := &Set{}
	ret.sets = make(map[T]bool)

	s.RLock()
	for i := range s.sets {
		ret.sets[i] = true
	}
	s.RUnlock()

	t.RLock()
	for i := range t.sets {
		if _, ok := ret.sets[i]; !ok {
			ret.sets[i] = true
		}
	}
	t.RUnlock()

	return ret
}

func (s *Set) Intersection(t *Set) *Set {
	ret := &Set{}
	ret.sets = make(map[T]bool)

	s.RLock()
	t.RLock()
	defer s.RUnlock()
	defer t.RUnlock()

	for i := range t.sets {
		if _, ok := s.sets[i]; ok {
			ret.sets[i] = true
		}
	}

	return ret
}

func (s *Set) Difference(t *Set) *Set {
	ret := &Set{}
	ret.sets = make(map[T]bool)

	s.RLock()
	t.RLock()
	defer s.RUnlock()
	defer t.RUnlock()

	for i := range t.sets {
		if _, ok := s.sets[i]; !ok {
			ret.sets[i] = true
		}
	}

	return ret
}

func (s *Set) Subset(t *Set) bool {
	s.RLock()
	t.RLock()
	defer s.RUnlock()
	defer t.RUnlock()

	for i := range s.sets {
		if _, ok := t.sets[i]; !ok {
			return false
		}
	}

	return true
}
