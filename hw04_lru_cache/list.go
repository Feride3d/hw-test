package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	elemCount int
	head      *ListItem
	tail      *ListItem
}

func NewList() List {
	return new(list)
}

// Amount of items in the list.
func (l *list) Len() int {
	return l.elemCount
}

// The first item of the list.
func (l *list) Front() *ListItem {
	return l.head
}

// The last item of the list.
func (l *list) Back() *ListItem {
	return l.tail
}

// Add new item to the front of the list.
func (l *list) PushFront(v interface{}) *ListItem {
	newItem := new(ListItem)
	if l.elemCount == 0 {
		newItem.Prev, newItem.Next, newItem.Value = nil, nil, v
		l.tail, l.head = newItem, newItem
		l.elemCount++
	} else {
		newItem.Value = v
		newItem.Next = l.head
		newItem.Prev = nil
		l.head.Prev = newItem
		l.head = newItem
		l.elemCount++
	}
	return newItem
}

// Add new item to the end of the list.
func (l *list) PushBack(v interface{}) *ListItem {
	newItem := new(ListItem)
	if l.elemCount == 0 {
		newItem.Prev, newItem.Next, newItem.Value = nil, nil, v
		l.tail, l.head = newItem, newItem
		l.elemCount++
	} else {
		newItem.Value = v
		newItem.Prev = l.tail
		newItem.Next = nil
		l.tail.Next = newItem
		l.tail = newItem
		l.elemCount++
	}
	return newItem
}

// Remove list item.
func (l *list) Remove(i *ListItem) {
	switch {
	case l.elemCount == 1:
		l.head, l.tail = nil, nil
	case i.Prev == nil:
		l.head = i.Next
		l.head.Prev = nil
	case i.Next == nil:
		l.tail = i.Prev
		l.tail.Next = nil
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	l.elemCount--
}

// Move list item to the front of the list.
func (l *list) MoveToFront(i *ListItem) {
	if i == l.head { // for head item
		return
	}
	if i == l.tail { // for tail item
		i.Prev.Next = nil
		l.tail = i.Prev
		i.Prev = nil
		i.Next = l.head
		l.head.Prev = i
		l.head = i
	} else { // for other items
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
		i.Prev = nil
		i.Next = l.head
		l.head.Prev = i
		l.head = i
	}
}
