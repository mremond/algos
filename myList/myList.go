/*
  Implementing and understanding complexity of linked list operations.
*/
package myList

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
