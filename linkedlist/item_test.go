package linkedlist

import (
	"fmt"
	"testing"
)
const UpperLimit 		= 5
const FirstIteration 	= 0
const CheckedQuantity 	= 6
func TestItem(t *testing.T) {
	
	//Создаем набор связанных элементов в цикле
	var currentElement *Item
	var firstElement *Item
	var lastElement *Item
	storage := make([]*Item, UpperLimit)
	
	for i := 0; i < UpperLimit; i++ {
		
		currentElement = &Item{
			struct{Data int}{i}, 								//создаем элемента списка для текущей интерации
			currentElement, 									//элемент current становится предыдущим 
			&Item{struct{Data int}{i+1}, currentElement, nil}}	//создаем сразу следующий элемент в списке	

		//Сохраняем для дальнейшей проверки
		if i == FirstIteration {
			firstElement = currentElement
		}

		fmt.Printf("creating: %v\n", currentElement)
		storage[i] = currentElement
	}
	
	lastElement = currentElement.Next()
	storage = append(storage, lastElement)
	
	
	//1) Тестируем корректность создания первого и последнего элемента последовательности
	if firstElement == nil {
		t.Errorf("Не создан первый элемент последовательности")
	}

	if lastElement == nil {
		t.Errorf("Не создан последний элемент последовательности")
	}

	if firstElement.Prev() != nil {
		t.Errorf("Ошибка при создании первого элемента последовательности: не должен иметь ссылки на предыдущий элемент")
	}

	if lastElement.Next() != nil {
		t.Errorf("Ошибка при создании последнего элемента последовательности: не должен иметь ссылки на следующий элемент")
	}

	//2) Проверяем остальные члены созданной последовательности
	fmt.Printf("%v\n", storage)

	fmt.Printf("%v: %v\n", storage[FirstIteration].Next(), firstElement.Next())
	fmt.Printf("%v: %v\n", storage[UpperLimit], lastElement)
}
