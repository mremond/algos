package ch1

import "testing"

type testUnit struct {
	input    string
	expected bool
}

var uniqueTestSet = []testUnit{
	testUnit{input: "ABCDEF", expected: true},
	testUnit{input: "ABCDEFF", expected: false},
}

func TestIsUniqueMap(t *testing.T) {
	for _, u := range uniqueTestSet {
		actual := isUnique(u.input)
		if actual != u.expected {
			t.Errorf("isUnique(%q): %v != %v", u.input, actual, u.expected)
		}
	}
}
