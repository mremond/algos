/* Reference: Introduction to Algorithms: Third Edition

When playing with sorting algorithms, you may be interested by discussion on that Golang
ticket: https://github.com/golang/go/issues/467
*/
package mySort

import "math"

// O(n^2)
// Sorts the list in-place: memory = O(n)
func InsertionSort(list []int) {
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
}

// O(n log n)
// Sorts the list in-place, but requires scratch slices.
// Interesting discussion on how to avoid scratch slices, at the expense of processing complexity:
// http://penguin.ewu.edu/cscd300/Topic/AdvSorting/MergeSorts/InPlace.html
func MergeSort(list []int) {
	mergeSort(list, 1, len(list))
}

func mergeSort(list []int, p, r int) {
	if p < r { // Only need sorting if list is longer than one
		q := (p + r) / 2
		mergeSort(list, p, q)
		mergeSort(list, q+1, r)
		merge(list, p, q, r)
	}
}

// Merge use a scratch slices as temporary buffer.
// May be optimised in Go ?
func merge(list []int, p, q, r int) {
	// Set up temporary scratch slices
	n1 := q - p + 1
	n2 := r - q
	L := make([]int, n1+1)
	R := make([]int, n2+1)
	for i := 0; i < n1; i++ {
		L[i] = list[p+i-1]
	}
	for j := 0; j < n2; j++ {
		R[j] = list[q+j]
	}
	// Trick to avoid testing if a list has been exausted: We add a "infinity" marker
	// at end of list.
	L[n1] = math.MaxInt64
	R[n2] = math.MaxInt64

	i := 0
	j := 0
	for k := p - 1; k < r; k++ {
		if L[i] <= R[j] {
			list[k] = L[i]
			i++
		} else {
			list[k] = R[j]
			j++
		}
	}
}
