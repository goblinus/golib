package linkedlist

import (
	"testing"
)

const FirstElementValue 	= "first"
const SecondElementValue 	= "second"
const LastElementValue		= "last"

func TestList(t *testing.T) {
	list := CreateList()
	list.PushFront(struct{Data string}{SecondElementValue})
	list.PushFront(struct{Data string}{FirstElementValue})
	list.PushBack(struct{Data string}{LastElementValue})
	
	counter := 0
	currentElement := list.First()
	for currentElement != nil {
		counter++
		currentElement = currentElement.Next()
	}

	//1) Проверяем кол-во элементов тестовой последовательности (двусвязаного списка)
	if counter != list.Len() {
		t.Errorf("Кол-во элементов тестовой последовательности не соответствует контрольному значению (3 элемента)")
	}

	//2) Проверяем значения элементов тестовой последовательности: элементы получаем из объекта "список"
	if list.First().Value().(struct{Data string}).Data != FirstElementValue {
		t.Errorf("Значение первого элемента тестовой последовательности (двусвязного списка) не соответствует контрлоьному значению")
	}

	if list.First().Next().Value().(struct{Data string}).Data != SecondElementValue {
		t.Errorf("Значение второго элемента тестовой последовательности (двусвязного списка) не соответствует контрлоьному значению")
	}

	if list.Last().Value().(struct{Data string}).Data != LastElementValue {
		t.Errorf("Значение последнего (третьего) элемента тестовой последовательности (двусвязного списка) не соответствует контрлоьному значению")
	}
}