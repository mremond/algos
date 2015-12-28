package myList_test

import (
	"testing"

	"github.com/mremond/go/myList"
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

func TestSingleElem(t *testing.T) {
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

func TestMultipleElem(t *testing.T) {
	list := myList.NewList()
	data := [5]int{1, 2, 3, 4, 5}
	for i, _ := range data {
		list.InsertNext(list.Tail, &data[i])
	}

	if  list.Size != len(data) {
		t.Error("Incorrect list size: ", list.Size," instead of ", len(data))
	}

//	for item := list.H;
}
