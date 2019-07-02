package eventlogger

import (
	"fmt"
	"strings"
	"testing"
)

const TestEventId = 5555

func TestEvent(t *testing.T) {
	//HwSubmitter: Сравниваем сообщение, которое генерирует событие, и соответствующее ему контрольное сообщение 
	if ok := hwSubmitterTest(TestEventId); !ok {
		t.Errorf("Ошибка построения строки лога для события HwSubmitter")
	}
	
	//HwAccepted: Сравниваем сообщение, которое генерирует событие, и соответствующее ему контрольное сообщение
	if ok := hwAcceptedTest(TestEventId); !ok {
		t.Errorf("Ошибка построения строки лога для события HwAccepted")
	}		
}

//Выносим в отдельную функцию создание тестового события HwSubmitter и проверку его работоспособности
func hwSubmitterTest(testEventId int) bool {
	hwSubmitterCode := "Some code"
	hwSubmitterComment := "Отправляем домашнее задание на проверку"
	lms := CreateHwSubmitterEvent(
		hwSubmitterCode,
		hwSubmitterComment)
	
	lms.Id = testEventId
	checkedEventMessage := fmt.Sprintf(
		"%s %s %d \"%s\"",
		lms.DateTime.Format("2006-01-02"),
		lms.Name,
		lms.Id,
		hwSubmitterComment)

	if EqualValue != strings.Compare(lms.LogMessage(), checkedEventMessage) {
		return false
	}

	return true
}

//Выносим в отдельную функцию создание события HwAccepted и проверку его работоспособности
func hwAcceptedTest(testEventId int) bool {
	hwAcceptedGrade := 5
	lma := CreateHwAcceptedEvent(hwAcceptedGrade)

	lma.Id = testEventId
	checkedEventMessage := fmt.Sprintf(
		"%s %s %d %d",
		lma.DateTime.Format("2006-01-02"),
		lma.Name,
		lma.Id,
		hwAcceptedGrade)

	if EqualValue != strings.Compare(lma.LogMessage(), checkedEventMessage) {
		return false
	}

	return true
}