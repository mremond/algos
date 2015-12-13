// Reference: Introduction to Algorithms: Third Edition
package mySort

// O(n2)
func InsertionSort(list []int) []int {
	// Start by comparing second element with existing list
	for j := 1; j < len(list); j++ {
		key := list[j]
		var i int
		for i = j - 1; i >= 0 && list[i] > key; {
			list[i+1] = list[i]
			i--
		}
		list[i+1] = key
	}
	return list
}

// O(n)
// Merge sort
