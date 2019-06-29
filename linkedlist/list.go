package linkedlist

func CreateList() *List {
	return &List{}
}

type List struct {
	length int
	firstElement *Item
	lastElement *Item
}

func (l *List) Len() int {
	return l.length
}

func (l *List) First() *Item {
	return l.firstElement
}

func (l *List) Last() *Item {
	return l.lastElement
}

func (l *List) PushFront(v interface{}) {
	newItem := CreateItem(v, nil, nil)
	if l.firstElement == nil {
		l.firstElement = newItem
		l.lastElement  = newItem
	} else {
		newItem.next 	= l.firstElement
		l.firstElement 	= newItem
		l.firstElement.prev = newItem
	}

	l.length++
}

func (l *List) PushBack(v interface{}) {
	newItem := CreateItem(v, nil, nil)
	if l.firstElement == nil && l.lastElement == nil {
		l.firstElement = newItem
		l.lastElement = newItem
	} else {
		newItem.prev = l.lastElement
		l.lastElement.next = newItem
		l.lastElement = newItem
	}
	
	l.length++
}