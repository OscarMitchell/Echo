package lib

import (
	"testing"
)

// TestNewSet tests the creation of an empty Set
func TestNewSet(t *testing.T) {
	s := NewSet[int]()
	if s == nil {
		t.Fatal("NewSet returned nil")
	}
	if s.Len() != 0 {
		t.Errorf("Expected empty set, got length %d", s.Len())
	}
	if !s.IsEmpty() {
		t.Error("Expected IsEmpty to return true for new set")
	}
}

// TestNewSetFrom tests creating a Set from a slice
func TestNewSetFrom(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{1}, 1},
		{"multiple elements", []int{1, 2, 3}, 3},
		{"duplicates", []int{1, 2, 2, 3, 3, 3}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSetFrom(tt.input)
			if s.Len() != tt.expected {
				t.Errorf("Expected length %d, got %d", tt.expected, s.Len())
			}
		})
	}
}

// TestAdd tests adding a single element
func TestAdd(t *testing.T) {
	s := NewSet[string]()
	s.Add("apple")

	if !s.Contains("apple") {
		t.Error("Expected set to contain 'apple'")
	}
	if s.Len() != 1 {
		t.Errorf("Expected length 1, got %d", s.Len())
	}

	// Adding duplicate
	s.Add("apple")
	if s.Len() != 1 {
		t.Errorf("Expected length to remain 1 after duplicate add, got %d", s.Len())
	}
}

// TestAddMany tests adding multiple elements
func TestAddMany(t *testing.T) {
	s := NewSet[int]()
	s.AddMany(1, 2, 3, 4, 5)

	if s.Len() != 5 {
		t.Errorf("Expected length 5, got %d", s.Len())
	}

	for i := 1; i <= 5; i++ {
		if !s.Contains(i) {
			t.Errorf("Expected set to contain %d", i)
		}
	}

	// Adding with duplicates
	s.AddMany(3, 4, 5, 6)
	if s.Len() != 6 {
		t.Errorf("Expected length 6, got %d", s.Len())
	}
}

// TestRemove tests removing a single element
func TestRemove(t *testing.T) {
	s := NewSetFrom([]int{1, 2, 3})

	s.Remove(2)
	if s.Contains(2) {
		t.Error("Expected element 2 to be removed")
	}
	if s.Len() != 2 {
		t.Errorf("Expected length 2, got %d", s.Len())
	}

	// Removing non-existent element should be no-op
	s.Remove(99)
	if s.Len() != 2 {
		t.Errorf("Expected length to remain 2, got %d", s.Len())
	}
}

// TestRemoveMany tests removing multiple elements
func TestRemoveMany(t *testing.T) {
	s := NewSetFrom([]int{1, 2, 3, 4, 5})

	s.RemoveMany(2, 3, 4)
	if s.Len() != 2 {
		t.Errorf("Expected length 2, got %d", s.Len())
	}

	if s.Contains(2) || s.Contains(3) || s.Contains(4) {
		t.Error("Expected elements 2, 3, 4 to be removed")
	}

	if !s.Contains(1) || !s.Contains(5) {
		t.Error("Expected elements 1 and 5 to remain")
	}
}

// TestClone tests cloning a Set
func TestClone(t *testing.T) {
	original := NewSetFrom([]int{1, 2, 3})
	clone := original.Clone()

	if !original.Equal(clone) {
		t.Error("Clone should be equal to original")
	}

	// Modify clone
	clone.Add(4)
	if original.Contains(4) {
		t.Error("Modifying clone should not affect original")
	}
	if clone.Len() != 4 {
		t.Errorf("Expected clone length 4, got %d", clone.Len())
	}
	if original.Len() != 3 {
		t.Errorf("Expected original length 3, got %d", original.Len())
	}
}

// TestClear tests clearing all elements
func TestClear(t *testing.T) {
	s := NewSetFrom([]int{1, 2, 3, 4, 5})
	s.Clear()

	if s.Len() != 0 {
		t.Errorf("Expected length 0 after clear, got %d", s.Len())
	}
	if !s.IsEmpty() {
		t.Error("Expected IsEmpty to return true after clear")
	}
}

// TestLen tests the length method
func TestLen(t *testing.T) {
	s := NewSet[int]()
	if s.Len() != 0 {
		t.Errorf("Expected length 0, got %d", s.Len())
	}

	s.Add(1)
	if s.Len() != 1 {
		t.Errorf("Expected length 1, got %d", s.Len())
	}

	s.AddMany(2, 3, 4)
	if s.Len() != 4 {
		t.Errorf("Expected length 4, got %d", s.Len())
	}
}

// TestIsEmpty tests the IsEmpty method
func TestIsEmpty(t *testing.T) {
	s := NewSet[int]()
	if !s.IsEmpty() {
		t.Error("Expected new set to be empty")
	}

	s.Add(1)
	if s.IsEmpty() {
		t.Error("Expected set with element to not be empty")
	}

	s.Remove(1)
	if !s.IsEmpty() {
		t.Error("Expected set to be empty after removing only element")
	}
}

// TestContains tests the Contains method
func TestContains(t *testing.T) {
	s := NewSetFrom([]string{"apple", "banana", "cherry"})

	tests := []struct {
		elem     string
		expected bool
	}{
		{"apple", true},
		{"banana", true},
		{"cherry", true},
		{"orange", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.elem, func(t *testing.T) {
			if s.Contains(tt.elem) != tt.expected {
				t.Errorf("Contains(%q) = %v, expected %v", tt.elem, !tt.expected, tt.expected)
			}
		})
	}
}

// TestEqual tests the Equal method
func TestEqual(t *testing.T) {
	tests := []struct {
		name     string
		set1     *Set[int]
		set2     *Set[int]
		expected bool
	}{
		{
			"equal sets",
			NewSetFrom([]int{1, 2, 3}),
			NewSetFrom([]int{1, 2, 3}),
			true,
		},
		{
			"equal sets different order",
			NewSetFrom([]int{3, 2, 1}),
			NewSetFrom([]int{1, 2, 3}),
			true,
		},
		{
			"different sets",
			NewSetFrom([]int{1, 2, 3}),
			NewSetFrom([]int{1, 2, 4}),
			false,
		},
		{
			"different sizes",
			NewSetFrom([]int{1, 2}),
			NewSetFrom([]int{1, 2, 3}),
			false,
		},
		{
			"both empty",
			NewSet[int](),
			NewSet[int](),
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set1.Equal(tt.set2) != tt.expected {
				t.Errorf("Equal() = %v, expected %v", !tt.expected, tt.expected)
			}
		})
	}
}

// TestUnion tests the Union method
func TestUnion(t *testing.T) {
	tests := []struct {
		name     string
		set1     *Set[int]
		set2     *Set[int]
		expected []int
	}{
		{
			"disjoint sets",
			NewSetFrom([]int{1, 2}),
			NewSetFrom([]int{3, 4}),
			[]int{1, 2, 3, 4},
		},
		{
			"overlapping sets",
			NewSetFrom([]int{1, 2, 3}),
			NewSetFrom([]int{2, 3, 4}),
			[]int{1, 2, 3, 4},
		},
		{
			"one empty set",
			NewSetFrom([]int{1, 2, 3}),
			NewSet[int](),
			[]int{1, 2, 3},
		},
		{
			"both empty",
			NewSet[int](),
			NewSet[int](),
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.set1.Union(tt.set2)
			if result.Len() != len(tt.expected) {
				t.Errorf("Union length = %d, expected %d", result.Len(), len(tt.expected))
			}
			for _, elem := range tt.expected {
				if !result.Contains(elem) {
					t.Errorf("Union missing element %d", elem)
				}
			}
		})
	}
}

// TestIntersection tests the Intersection method
func TestIntersection(t *testing.T) {
	tests := []struct {
		name     string
		set1     *Set[int]
		set2     *Set[int]
		expected []int
	}{
		{
			"overlapping sets",
			NewSetFrom([]int{1, 2, 3, 4}),
			NewSetFrom([]int{3, 4, 5, 6}),
			[]int{3, 4},
		},
		{
			"disjoint sets",
			NewSetFrom([]int{1, 2}),
			NewSetFrom([]int{3, 4}),
			[]int{},
		},
		{
			"one empty set",
			NewSetFrom([]int{1, 2, 3}),
			NewSet[int](),
			[]int{},
		},
		{
			"identical sets",
			NewSetFrom([]int{1, 2, 3}),
			NewSetFrom([]int{1, 2, 3}),
			[]int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.set1.Intersection(tt.set2)
			if result.Len() != len(tt.expected) {
				t.Errorf("Intersection length = %d, expected %d", result.Len(), len(tt.expected))
			}
			for _, elem := range tt.expected {
				if !result.Contains(elem) {
					t.Errorf("Intersection missing element %d", elem)
				}
			}
		})
	}
}

// TestDifference tests the Difference method
func TestDifference(t *testing.T) {
	tests := []struct {
		name     string
		set1     *Set[int]
		set2     *Set[int]
		expected []int
	}{
		{
			"overlapping sets",
			NewSetFrom([]int{1, 2, 3, 4}),
			NewSetFrom([]int{3, 4, 5, 6}),
			[]int{1, 2},
		},
		{
			"disjoint sets",
			NewSetFrom([]int{1, 2}),
			NewSetFrom([]int{3, 4}),
			[]int{1, 2},
		},
		{
			"subset",
			NewSetFrom([]int{1, 2, 3}),
			NewSetFrom([]int{2}),
			[]int{1, 3},
		},
		{
			"empty difference",
			NewSetFrom([]int{1, 2}),
			NewSetFrom([]int{1, 2, 3}),
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.set1.Difference(tt.set2)
			if result.Len() != len(tt.expected) {
				t.Errorf("Difference length = %d, expected %d", result.Len(), len(tt.expected))
			}
			for _, elem := range tt.expected {
				if !result.Contains(elem) {
					t.Errorf("Difference missing element %d", elem)
				}
			}
		})
	}
}

// TestSymmetricDifference tests the SymmetricDifference method
func TestSymmetricDifference(t *testing.T) {
	tests := []struct {
		name     string
		set1     *Set[int]
		set2     *Set[int]
		expected []int
	}{
		{
			"overlapping sets",
			NewSetFrom([]int{1, 2, 3}),
			NewSetFrom([]int{2, 3, 4}),
			[]int{1, 4},
		},
		{
			"disjoint sets",
			NewSetFrom([]int{1, 2}),
			NewSetFrom([]int{3, 4}),
			[]int{1, 2, 3, 4},
		},
		{
			"identical sets",
			NewSetFrom([]int{1, 2, 3}),
			NewSetFrom([]int{1, 2, 3}),
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.set1.SymmetricDifference(tt.set2)
			if result.Len() != len(tt.expected) {
				t.Errorf("SymmetricDifference length = %d, expected %d", result.Len(), len(tt.expected))
			}
			for _, elem := range tt.expected {
				if !result.Contains(elem) {
					t.Errorf("SymmetricDifference missing element %d", elem)
				}
			}
		})
	}
}

// TestIsSubset tests the IsSubset method
func TestIsSubset(t *testing.T) {
	tests := []struct {
		name     string
		set1     *Set[int]
		set2     *Set[int]
		expected bool
	}{
		{
			"proper subset",
			NewSetFrom([]int{1, 2}),
			NewSetFrom([]int{1, 2, 3, 4}),
			true,
		},
		{
			"equal sets",
			NewSetFrom([]int{1, 2, 3}),
			NewSetFrom([]int{1, 2, 3}),
			true,
		},
		{
			"not subset",
			NewSetFrom([]int{1, 2, 5}),
			NewSetFrom([]int{1, 2, 3, 4}),
			false,
		},
		{
			"empty set is subset",
			NewSet[int](),
			NewSetFrom([]int{1, 2, 3}),
			true,
		},
		{
			"superset not subset",
			NewSetFrom([]int{1, 2, 3, 4}),
			NewSetFrom([]int{1, 2}),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set1.IsSubset(tt.set2) != tt.expected {
				t.Errorf("IsSubset() = %v, expected %v", !tt.expected, tt.expected)
			}
		})
	}
}

// TestIsSuperset tests the IsSuperset method
func TestIsSuperset(t *testing.T) {
	tests := []struct {
		name     string
		set1     *Set[int]
		set2     *Set[int]
		expected bool
	}{
		{
			"proper superset",
			NewSetFrom([]int{1, 2, 3, 4}),
			NewSetFrom([]int{1, 2}),
			true,
		},
		{
			"equal sets",
			NewSetFrom([]int{1, 2, 3}),
			NewSetFrom([]int{1, 2, 3}),
			true,
		},
		{
			"not superset",
			NewSetFrom([]int{1, 2, 3}),
			NewSetFrom([]int{1, 2, 5}),
			false,
		},
		{
			"any set is superset of empty",
			NewSetFrom([]int{1, 2, 3}),
			NewSet[int](),
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set1.IsSuperset(tt.set2) != tt.expected {
				t.Errorf("IsSuperset() = %v, expected %v", !tt.expected, tt.expected)
			}
		})
	}
}

// TestIsDisjoint tests the IsDisjoint method
func TestIsDisjoint(t *testing.T) {
	tests := []struct {
		name     string
		set1     *Set[int]
		set2     *Set[int]
		expected bool
	}{
		{
			"disjoint sets",
			NewSetFrom([]int{1, 2, 3}),
			NewSetFrom([]int{4, 5, 6}),
			true,
		},
		{
			"overlapping sets",
			NewSetFrom([]int{1, 2, 3}),
			NewSetFrom([]int{3, 4, 5}),
			false,
		},
		{
			"empty sets are disjoint",
			NewSet[int](),
			NewSet[int](),
			true,
		},
		{
			"empty and non-empty are disjoint",
			NewSet[int](),
			NewSetFrom([]int{1, 2, 3}),
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set1.IsDisjoint(tt.set2) != tt.expected {
				t.Errorf("IsDisjoint() = %v, expected %v", !tt.expected, tt.expected)
			}
		})
	}
}

// TestAll tests the iterator method
func TestAll(t *testing.T) {
	s := NewSetFrom([]int{1, 2, 3, 4, 5})

	count := 0
	seen := make(map[int]bool)

	for elem := range s.All() {
		count++
		seen[elem] = true
	}

	if count != 5 {
		t.Errorf("Expected to iterate over 5 elements, got %d", count)
	}

	for i := 1; i <= 5; i++ {
		if !seen[i] {
			t.Errorf("Expected to see element %d during iteration", i)
		}
	}
}

// TestAllEarlyExit tests that the iterator respects early exit
func TestAllEarlyExit(t *testing.T) {
	s := NewSetFrom([]int{1, 2, 3, 4, 5})

	count := 0
	for range s.All() {
		count++
		if count == 3 {
			break
		}
	}

	if count != 3 {
		t.Errorf("Expected to break after 3 iterations, got %d", count)
	}
}

// TestString tests the String method
func TestString(t *testing.T) {
	tests := []struct {
		name     string
		set      *Set[int]
		contains []string
	}{
		{
			"empty set",
			NewSet[int](),
			[]string{"Set{}"},
		},
		{
			"single element",
			NewSetFrom([]int{1}),
			[]string{"Set{", "1", "}"},
		},
		{
			"multiple elements",
			NewSetFrom([]int{1, 2, 3}),
			[]string{"Set{", "1", "2", "3", "}"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := tt.set.String()
			for _, substr := range tt.contains {
				if !contains(str, substr) {
					t.Errorf("String() = %q, expected to contain %q", str, substr)
				}
			}
		})
	}
}

// TestStringTypes tests String with different types
func TestStringTypes(t *testing.T) {
	strSet := NewSetFrom([]string{"apple", "banana"})
	str := strSet.String()
	if !contains(str, "Set{") || !contains(str, "}") {
		t.Errorf("String set string representation malformed: %s", str)
	}
}

// TestSetWithDifferentTypes tests that Set works with different comparable types
func TestSetWithDifferentTypes(t *testing.T) {
	// Test with strings
	strSet := NewSetFrom([]string{"a", "b", "c"})
	if strSet.Len() != 3 {
		t.Errorf("String set length = %d, expected 3", strSet.Len())
	}

	// Test with float64
	floatSet := NewSetFrom([]float64{1.1, 2.2, 3.3})
	if floatSet.Len() != 3 {
		t.Errorf("Float set length = %d, expected 3", floatSet.Len())
	}

	// Test with bool
	boolSet := NewSetFrom([]bool{true, false, true})
	if boolSet.Len() != 2 {
		t.Errorf("Bool set length = %d, expected 2", boolSet.Len())
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || indexOfSubstring(s, substr) >= 0)
}

func indexOfSubstring(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
