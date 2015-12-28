/*
  Implementing and understanding complexity of linked list operations.
*/
package myList

type Elem struct {
	Data interface{}
	Next *Elem
}

type List struct {
	H    *Elem // Head
	T    *Elem // Tail: Shortcut to last element (=> Not equivalent to cdr in lisp)
	Size uint
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
		newElem.Next = list.H
		list.H = newElem
		if list.Size == 0 {
			list.T = newElem // If list size was 0, first and last elements are the same
		}
	} else { // Insert element somewhere other than head
		newElem.Next = elem.Next
		elem.Next = newElem
		if elem.Next == nil {
			list.T = newElem // We are inserting at the end
		}
	}

	list.Size++
}

func (list List) Head() *Elem {
	return list.H
}

// Should we return the tail wrapped in a list instead of directly a pointer to the first element of the tail ?
func (list List) Tail() *Elem {
	return list.T
}
