/*
  Implementing and understanding complexity of linked list operations.

  Implementation inspired by "Mastering Algorithms in C", Kyle Loudon.

  Benefit is that InsertNext and RemoveNext have execution cost of O(1).

  Edge cases to get right when managing linked list:

  - Insert on first and last position
  - Delete on first and last position
  - Delete on empty list

  What about passing a reference element that is not in the list ?
  We cannot easily check integrity

*/
package myList
import "fmt"

type Elem struct {
	Data interface{}
	Next *Elem
}

type List struct {
	Head *Elem // Head
	Tail *Elem // Tail: Shortcut to last element (=> Not equivalent to cdr in lisp)
	Size int
}

func NewList() *List {
	return new(List)
}

// InsertNext insert data in list after element.
// If the elem to insert after is nil, we insert the element in front
func (list *List) InsertNext(elem *Elem, data interface{}) {
	// Create new element
	newElem := new(Elem)
	newElem.Data = data

	// Insert element in front
	if elem == nil {
		newElem.Next = list.Head
		list.Head = newElem
		if list.Size == 0 {
			list.Tail = newElem // If list size was 0, first and last elements are the same
		}
	} else { // Insert element somewhere other than head
		if elem.Next == nil {
			list.Tail = newElem // We are inserting at the end
		}
		newElem.Next = elem.Next
		elem.Next = newElem
	}

	list.Size++
}

// RemoveNext delete an element from the list after the one passed and returns the delete element data
// We cannot directly remove any element from the list as the list is a single linked list. We cannot get the previous
// element.
// Element nil is used to remove from head of list.
func (list *List) RemoveNext(elem *Elem) (data interface{}, err error) {
	var old_element *Elem

	if list.Size == 0 {
		err = fmt.Errorf("Cannot remove element from empty list")
		return
	}

	if elem == nil { // If we are removing first element from list
		old_element = list.Head
		data = old_element.Data

		if list.Size == 1 { // If list was only one element long before removal
			list.Head = nil
			list.Tail = nil
		}
	} else { // Handle removal from somewhere other than head
		if elem.Next == nil {
			err = fmt.Errorf("Element is last in the list, cannot remove next")
			return
		}
		old_element = elem.Next
		data = old_element.Data
		elem.Next = old_element.Next

		if elem.Next == nil { // This is now the last element, we update tail shortcut
			list.Tail = elem
		}
	}
	list.Size--
	return
}

// TODO: Remove a given element from list. This is a helper in which we need to iterate list from Head to find the element to delete