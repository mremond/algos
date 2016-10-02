package ch1

// isUnique returns true is all chars in the string are unique or false if one
// or more duplicate is found. isUnique is case sensitive, i.e. 'a' and 'A' are consider different.
func isUnique(s string) bool {
	index := make(map[rune]bool)
	for _, c := range s {
		if index[c] {
			return false
		}
		index[c] = true
	}
	return true
}
