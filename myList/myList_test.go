package myList_test

import (
	"testing"

	"github.com/mremond/algos/myList"
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
		t.Error("Incorrect size after InsertNext: ", list.Size, "instead of 1")
	}

	HData := list.H.Data.(*int)
	if *HData != data {
		t.Error("Incorrect data at head: ", HData)
	}

	TData := list.T.Data.(*int)
	if *TData != data {
		t.Error("Incorrect data at tail: ", list.T.Data)
	}

}
