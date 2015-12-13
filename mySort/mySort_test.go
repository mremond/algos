package mySort_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/mremond/algos/mySort"
)

// Testing InsertionSort

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestInsertionSortBasic(t *testing.T) {
	// Test with fully reversed list
	list := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	mySort.InsertionSort(list)
	checkOrder(list, t)
}

func TestInsertionSortRandom(t *testing.T) {
	list := buildRandomIntList(100)
	mySort.InsertionSort(list)
	checkOrder(list, t)
}

func benchmarkInsertionSort(i int, b *testing.B) {
	list := buildRandomIntList(i)
	for n := 0; n < b.N; n++ {
		mySort.InsertionSort(list)
	}
}

func BenchmarkInsertionSort10(b *testing.B)     { benchmarkInsertionSort(10, b) }
func BenchmarkInsertionSort100(b *testing.B)    { benchmarkInsertionSort(100, b) }
func BenchmarkInsertionSort1000(b *testing.B)   { benchmarkInsertionSort(1000, b) }
func BenchmarkInsertionSort10000(b *testing.B)  { benchmarkInsertionSort(10000, b) }
func BenchmarkInsertionSort100000(b *testing.B) { benchmarkInsertionSort(100000, b) }

// MergeSort

func TestMergeSortMin(t *testing.T) {
	// Test with minimal list
	list := []int{2, 1}
	mySort.MergeSort(list)
	checkOrder(list, t)
}

func TestMergeSortBasic(t *testing.T) {
	// Test with fully reversed list (2^4)
	list := []int{7, 8, 5, 9, 4, 2, 6, 10, 3, 1}
	mySort.MergeSort(list)
	checkOrder(list, t)
}

func TestMergeSortRandom(t *testing.T) {
	list := buildRandomIntList(100)
	mySort.MergeSort(list)
	checkOrder(list, t)
}

func benchmarkMergeSort(i int, b *testing.B) {
	list := buildRandomIntList(i)
	for n := 0; n < b.N; n++ {
		mySort.MergeSort(list)
	}
}

func BenchmarkMergeSort10(b *testing.B)      { benchmarkMergeSort(10, b) }
func BenchmarkMergeSort100(b *testing.B)     { benchmarkMergeSort(100, b) }
func BenchmarkMergeSort1000(b *testing.B)    { benchmarkMergeSort(1000, b) }
func BenchmarkMergeSort10000(b *testing.B)   { benchmarkMergeSort(10000, b) }
func BenchmarkMergeSort100000(b *testing.B)  { benchmarkMergeSort(100000, b) }
func BenchmarkMergeSort1000000(b *testing.B) { benchmarkMergeSort(1000000, b) }

// Helpers

func buildRandomIntList(size int) []int {
	myList := make([]int, size)
	for i := 0; i < size; i++ {
		myList[i] = rand.Int()
	}
	return myList
}

func checkOrder(sortedList []int, t *testing.T) {
	tmp := sortedList[0]
	for i := 1; i < len(sortedList)-2; i++ {
		if sortedList[i] < tmp {
			t.Error("List not sorted:", sortedList[i], "<", tmp)
			//			 break
		}
		tmp = sortedList[i]
	}
}
