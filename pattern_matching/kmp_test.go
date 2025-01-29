package pattern_matching

import (
	"reflect"
	"testing"
)

func TestKMP(t *testing.T) {
	tests := []struct {
		text     string
		pattern  string
		expected []int
	}{
		{"ababcabcabababd", "ababd", []int{10}},
		{"ababcabcabababd", "abc", []int{2, 5}},
		{"aaaaa", "aa", []int{0, 1, 2, 3}},
		{"", "a", nil},
		{"a", "", nil},
		{"a", "a", []int{0}},
		{"abc", "abcd", nil},
	}

	for _, test := range tests {
		result := KMP(test.text, test.pattern)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("KMP(%q, %q) = %v; expected %v", test.text, test.pattern, result, test.expected)
		}
	}
}
