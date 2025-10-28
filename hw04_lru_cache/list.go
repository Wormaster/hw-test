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
	len         int
	first, last *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	n := &ListItem{Value: v, Next: l.first}
	if l.first != nil {
		l.first.Prev = n
	} else {
		l.last = n
	}
	l.first = n
	l.len++
	return n
}

func (l *list) PushBack(v interface{}) *ListItem {
	n := &ListItem{Value: v, Prev: l.last}
	if l.last != nil {
		l.last.Next = n
	} else {
		l.first = n
	}
	l.last = n
	l.len++
	return n
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.first = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.last = i.Prev
	}
	i.Next, i.Prev = nil, nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.first = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.last = i.Prev
	}

	i.Prev = nil
	i.Next = l.first
	if l.first != nil {
		l.first.Prev = i
	}
	l.first = i
	if l.last == nil {
		l.last = i
	}
}

func NewList() List {
	return new(list)
}
