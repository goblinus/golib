package linkedlist

import (
	"math/rand"
	"testing"
	"time"
)

const UpperLimit = 5
const FirstIteration = 0

func TestItem(t *testing.T) {

	//Создаем набор связанных элементов в цикле
	var currentElement *Item
	var firstElement *Item

	var counter int
	var storage []int

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for counter = 0; counter < UpperLimit; counter++ {

		if currentElement == nil {
			value := random.Intn(10000)
			currentElement = CreateItem(
				struct{ Data int }{value},
				nil,
				nil)
			firstElement = currentElement
			storage = append(storage, value)
		}

		value := random.Intn(10000)
		currentElement.next = CreateItem(
			struct{ Data int }{value},
			currentElement,
			nil)

		storage = append(storage, value)
		currentElement = currentElement.Next()
	}

	//1) Проверяем наличие входной точки последовательности для последующего перебора
	if firstElement == nil {
		t.Errorf("Пустая последовательность: нет первого элемента контрольной последовательности")
	}

	counter = 0
	checkQnty := len(storage)
	currentElement = firstElement

	//2) Сравниваем значения элементов последовательности с соответствующими контрольными значениями (массив storage)
	for currentElement != nil {
		itemValue := currentElement.Value().(struct{ Data int }).Data
		if itemValue != storage[counter] {
			t.Errorf(
				"Текущее значение элемента не соответствует контрольному значению: %d : %d\n",
				itemValue,
				storage[counter])
		}

		counter++
		currentElement = currentElement.Next()
	}

	//3) Проверяем общее кол-во элементов последовательности с кол-вом контрольных элементов
	if counter != checkQnty {
		t.Errorf("Кол-во элементов контрольной последовательности не соответствует кол-ву элементов тестируемого списка")
	}

	//4) Тестируем удаление элемента последовательности
	deletedElement := firstElement.Next()
	deletedElement.Remove()
	currentElement = firstElement
	for currentElement != nil {
		currentElement = currentElement.Next()
	}

	shouldBeValue := storage[2]
	if (firstElement.Next().Value().(struct{ Data int })).Data != shouldBeValue {
		t.Errorf("Ошибка при удалении второго элемента контрольной последовательности")
	}
}
