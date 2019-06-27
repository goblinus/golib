package linkedlist

import (
	"fmt"
	"testing"
)

const UpperLimit = 5
const FirstIteration = 0
const CheckedQuantity = 6

func TestItem(t *testing.T) {

	//Создаем набор связанных элементов в цикле
	var currentElement *Item
	var firstElement *Item
	var lastElement *Item

	storage := make([]*Item, 10)

	var counter int
	for counter = 0; counter < 4; counter++ {

		if currentElement == nil {
			currentElement = CreateItem(
				struct{ Data int }{counter},
				nil,
				nil)
			firstElement = currentElement
		}

		currentElement.next = CreateItem(
			struct{ Data int }{counter + 1},
			currentElement,
			nil)

		storage[counter] = currentElement
		currentElement = currentElement.next
	}

	//Обрабатываем последнюю итерацию
	counter++
	storage[counter] = currentElement
	fmt.Printf("ddd: %v\n", storage[counter])
	lastElement = storage[counter+1]

	fmt.Printf("Первый элемент последовательности: %p: %v\n", firstElement, firstElement)
	fmt.Printf("Последний элемент последовательности: %p: %v\n", lastElement, lastElement)
	fmt.Printf("Длина контрольного массива: %d\n", counter)
	for i := 0; i < CheckedQuantity; i++ {
		fmt.Printf("%p: %v\n", storage[i], storage[i])
	}
	//1) Тестируем корректность создания первого и последнего элемента последовательности
	// if firstElement == nil {
	// 	t.Errorf("Не создан первый элемент последовательности")
	// }

	// if lastElement == nil {
	// 	t.Errorf("Не создан последний элемент последовательности")
	// }

	// if firstElement.Prev() != nil {
	// 	t.Errorf("Ошибка при создании первого элемента последовательности: не должен иметь ссылки на предыдущий элемент")
	// }

	// if lastElement.Next() != nil {
	// 	t.Errorf("Ошибка при создании последнего элемента последовательности: не должен иметь ссылки на следующий элемент")
	// }

	//2) Проверяем остальные члены созданной последовательности
	// fmt.Printf("%v\n", storage)

	// fmt.Printf("%v: %v\n", storage[FirstIteration].Next(), firstElement.Next())
	// fmt.Printf("%v: %v\n", storage[UpperLimit], lastElement)
}
