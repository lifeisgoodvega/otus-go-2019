package list

import (
	"errors"
)

// Item is a node of double-linked list.
type Item struct {
	next, prev  *Item
	parent      *List
	invalidated *bool
	data        interface{}
}

// Value is a data from node.
func (i *Item) Value() interface{} {
	return i.data
}

// Next node.
func (i *Item) Next() *Item {
	return i.next
}

// Prev - previous node.
func (i *Item) Prev() *Item {
	return i.prev
}

// List is a double linked list for interface{} objects.
type List struct {
	length      int
	first, last *Item
}

// Len - get length of list
func (l *List) Len() int {
	return l.length
}

// First item of list.
func (l *List) First() *Item {
	return l.first
}

// Last item of list
func (l *List) Last() *Item {
	return l.last
}

// PushFront - push item at begining.
func (l *List) PushFront(v interface{}) {
	if l.first == nil {
		// zerovalue for bool is false
		l.first = &Item{data: v, parent: l, invalidated: new(bool)}
		l.last = l.first
	} else {
		formerFirst := l.first
		l.first = &Item{data: v, next: formerFirst, parent: l, invalidated: new(bool)}
		formerFirst.prev = l.first
	}
	l.length++
}

// PushBack - push item at end.
func (l *List) PushBack(v interface{}) {
	if l.first == nil {
		l.first = &Item{data: v, parent: l, invalidated: new(bool)}
		l.last = l.first
	} else {
		formerLast := l.last
		l.last = &Item{data: v, prev: formerLast, parent: l, invalidated: new(bool)}
		formerLast.next = l.last
	}
	l.length++
}

// Remove item from list.
func (l *List) Remove(i Item) error {
	if i.parent != l {
		return errors.New("Item must belong to list, which method is called")
	}
	if *i.invalidated {
		return errors.New("Item was deleted, therefore invalidated")
	}
	if i.prev != nil && i.next != nil {
		i.prev.next = i.next
		i.next.prev = i.prev
	} else if i.prev == nil && i.next != nil {
		i.next.prev = nil
		l.first = i.next

	} else if i.next == nil && i.prev != nil {
		i.prev.next = nil
		l.last = i.prev
	} else {
		l.first = nil
		l.last = nil
	}

	*i.invalidated = true
	l.length--
	return nil
}
