package myList_test

import (
	"testing"

	"github.com/mremond/go/myList"
	"fmt"
)

func TestNewList(t *testing.T) {
	list := myList.NewList()

	if list.Size != 0 {
		t.Error("New list should have a size of zero")
	}
}

type myStruct struct {
	value int
}

func TestSingleElemInsert(t *testing.T) {
	list := myList.NewList()
	data := 1
	list.InsertNext(nil, &data) // nil -> Insert in front

	if list.Size != 1 {
		t.Error("Incorrect size after InsertNext: ", list.Size, " instead of 1")
	}

	HData := list.Head.Data.(*int)
	if *HData != data {
		t.Error("Incorrect data at head: ", HData)
	}

	TData := list.Tail.Data.(*int)
	if *TData != data {
		t.Error("Incorrect data at tail: ", list.Tail.Data)
	}
}

func TestMultipleElemInsert(t *testing.T) {
	list := myList.NewList()
	data := [5]int{1, 2, 3, 4, 5}
	for i, _ := range data {
		list.InsertNext(list.Tail, &data[i])
	}

	if list.Size != len(data) {
		t.Error("Incorrect list size: ", list.Size," instead of ", len(data))
	}

	current_item := list.Head
	i := 0
	for current_item.Next != nil {
		current_data := current_item.Data.(*int)
		if *current_data != data[i] {
			t.Error("Incorrect element found: ", current_item.Data, " instead of ", data[i])
		}
		current_item = current_item.Next
		i++
	}

	// Check tail
	TData := list.Tail.Data.(*int)
	if *TData != data[len(data) - 1] {
		t.Error("Incorrect data at tail: ", list.Tail.Data)
	}
}

func TestRemoveAllFromHead(t *testing.T) {
	list := myList.NewList()
	data := [5]int{1, 2, 3, 4, 5}
	for i, _ := range data {
		list.InsertNext(list.Tail, &data[i])
	}

	size := list.Size
	for i := 0; i < size; i++ {
		fmt.Println("Elements to remove: ", list.Size)
		fmt.Println("Removing Element", list.Head.Data)
		if _, err := list.RemoveNext(nil); err != nil {
			t.Error("Error when removing item from list: ", err)
		}
		fmt.Println("Size ", list.Size)
	}

	if list.Size != 0 {
		t.Error("After removing all elements from list, list size should be zero: ", list.Size)
	}
}
