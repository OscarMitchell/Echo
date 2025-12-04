package lib

import (
	"fmt"
	"iter"
	"maps"
	"strings"
)

// Set is a wrapper around a standard map and allows for most common operations
// expected of a set collection
type Set[T comparable] struct {
	elements map[T]struct{}
}

// NewSet creates a new empty Set and returns a pointer to it
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]struct{})}
}

// NewSetFrom creates a new Set from the elements of a slice
func NewSetFrom[T comparable](elems []T) *Set[T] {
	new := NewSet[T]()
	new.AddMany(elems...)

	return new
}

// Add adds the given element into the Set
func (s *Set[T]) Add(elem T) {
	s.elements[elem] = struct{}{}
}

// AddMany adds the given element(s) into the Set
func (s *Set[T]) AddMany(elems ...T) {
	for _, elem := range elems {
		s.elements[elem] = struct{}{}
	}
}

// Remove deletes the given element(s) from the Set. If an element is not in the
// Set then Remove is a no op
func (s *Set[T]) Remove(elem T) {
	delete(s.elements, elem)
}

// RemoveMany deletes the given element(s) from the Set. If an element is not in
// the Set then Remove is a no op
func (s *Set[T]) RemoveMany(elems ...T) {
	for _, elem := range elems {
		delete(s.elements, elem)
	}
}

// Clone creates a new Set with identical elements to the base Set
func (s *Set[T]) Clone() *Set[T] {
	return &Set[T]{maps.Clone(s.elements)}
}

// Clear deletes every element in the Set
func (s *Set[T]) Clear() {
	s.elements = make(map[T]struct{})
}

// Len returns the number of elements in the Set
func (s *Set[T]) Len() int {
	return len(s.elements)
}

// IsEmpty returns True if the length of the Set is != 0
func (s *Set[T]) IsEmpty() bool {
	return s.Len() == 0
}

// Contains returns whether the given element is present in the Set
func (s *Set[T]) Contains(elem T) bool {
	_, ok := s.elements[elem]
	return ok
}

// Equal returns true if the elements of the base Set are match the elements of
// the other Set
func (s *Set[T]) Equal(other *Set[T]) bool {
	return maps.Equal(s.elements, other.elements)
}

// Union returns a pointer to a new Set with both the base Set and the other
// Sets elements
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	new := &Set[T]{elements: maps.Clone(s.elements)}
	for elem := range other.elements {
		new.elements[elem] = struct{}{}
	}

	return new
}

// Intersection returns a pointer to a new Set with only the elements that are
// present in both the base Set and the other Set
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	new := NewSet[T]()
	for elem := range other.elements {
		if s.Contains(elem) {
			new.elements[elem] = struct{}{}
		}
	}

	return new
}

// Difference returns a pointer to a new Set with only the elements that are
// Present in the base Set but not the other Set
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	new := NewSet[T]()
	for elem := range s.elements {
		if !other.Contains(elem) {
			new.elements[elem] = struct{}{}
		}
	}

	return new
}

// SymmetricDifference returns a pointer to a new Set with only the elements
// that are unique in both the base Set and the other Set
func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	return s.Difference(other).Union(other.Difference(s))
}

// IsSubset returns true if all of the elements of the base Set are also
// present in the other Set
func (s *Set[T]) IsSubset(other *Set[T]) bool {
	for elem := range s.elements {
		if !other.Contains(elem) {
			return false
		}
	}

	return true
}

// IsSuperset returns true if the base Set at least contains all of the elements
// present in the other Set
func (s *Set[T]) IsSuperset(other *Set[T]) bool {
	return other.IsSubset(s)
}

// IsDisjoint returns true if the elements of both the base Set and the other
// Set are entirely unique from each other
func (s *Set[T]) IsDisjoint(other *Set[T]) bool {
	for elem := range s.elements {
		if other.Contains(elem) {
			return false
		}
	}

	return true
}

// All returns a sequence to use in a for loop with the range keyword
func (s *Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range s.elements {
			if !yield(item) {
				return
			}
		}
	}
}

// String returns a string representation of the Set
func (s *Set[T]) String() string {
	if s.IsEmpty() {
		return "Set{}"
	}

	var sb strings.Builder
	sb.Grow(s.Len() * 10)

	_, _ = sb.WriteString("Set{")

	first := true
	for elem := range s.elements {
		if !first {
			_, _ = sb.WriteString(", ")
		}
		_, _ = fmt.Fprintf(&sb, "%v", elem)
		first = false
	}
	_, _ = sb.WriteString("}")

	return sb.String()
}
