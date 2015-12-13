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
	checkOrder(mySort.InsertionSort(list), t)
}

func TestInsertionSortRandom(t *testing.T) {
	list := buildRandomIntList(100)
	checkOrder(mySort.InsertionSort(list), t)
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
