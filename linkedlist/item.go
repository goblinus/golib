package linkedlist

func CreateItem(value interface{}, prev, next *Item) *Item {
	return &Item{value, prev, next}
}

//Базовая структуру элемента двусвязанного списка
type Item struct {
	value interface{}
	prev  *Item
	next  *Item
}

//Реализация интерфейса ItemLinker

func (i *Item) Value() interface{} {
	return i.value
}

func (i *Item) Prev() *Item {
	return i.prev
}

func (i *Item) Next() *Item {
	return i.next
}

func (i *Item) Remove() {
	if i.Prev() != nil && i.Next() != nil {
		i.prev.next = i.next
		i.next.prev = i.prev
	} else if i.Prev() == nil {
		i.next.prev = nil
	} else if i.Next() == nil {
		i.prev.next = nil
	}
}
