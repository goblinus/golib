package eventlogger

import (
	"fmt"
	"testing"
)

const TestEventId = 5555

func TestEvent(t *testing.T) {
	hwSubmitterTest(TestEventId)
	hwAcceptedTest(TestEventId)		
}

func hwSubmitterTest(testEventId int) {
	hwSubmitterCode := "Some code"
	hwSubmitterComment := "Отправляем домашнее задание на проверку"
	lms := CreateHwSubmitterEvent(
		hwSubmitterCode,
		hwSubmitterComment)
	
	lms.Event.Id = testEventId
	checkedEventMessage := fmt.Sprintf(
		"%s %s %d \"%s\"",
		lms.DateTime.Format("2006-01-02"),
		lms.Name,
		hwSubmitterComment)

	if lms.LogMessage() != checkedEventMessage {
		fmt.Errorf("Ошибка построения строки лога для события HwSubmitter")
	}
}

func hwAcceptedTest(testEventId int) {
	hwAcceptedGrade := 5
	lma := CreateHwAcceptedEvent(hwAcceptedGrade)

	lma.Event.Id = testEventId
	checkedEventMessage := fmt.Sprintf(
		"%s %s %d %d",
		lma.DateTime.Format("2006-01-02"),
		lma.Name,
		lma.Id,
		hwAcceptedGrade)

	if lma.LogMessage() != checkedEventMessage {
		t.Errorf("Ошибка построения строки лога для события HwAccepted")
	}
}