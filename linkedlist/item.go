package linkedlist

type ItemLinker interface {
	Value() interface{}
	Prev() *Item
	Next() *Item
	Remove()
}

//Базовая структуру элемента двусвязанного списка
type Item struct {
	value interface{}
	prev *Item
	next *Item
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

}
