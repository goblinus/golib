package eventlogger

import (
	"bytes"
	"fmt"
	"bytes"
	"time"
	"testing"
)

const EventId = 1111
const GradeValue = 5
const EqualValue = 0

func TestEvent(t *testing.T) {

	//Тестируем логирование события HwSubmitter
	checkedMessage := fmt.Sprintf(
		"%s submitted %d \"Comment for code\"\n",
		time.Now().Format("2006-01-02"),
		EventId)

	buffer = &bytes.Buffer{}
	submitterEvent := CreateHwSubmitterEvent(
		"Some code", 
		"Comment for code")
	submitterEvent.Id = EventId
	LogOtusEvent(submitterEvent, buffer)

	if EqualValue != strings.Compare(buffer.String(), checkedMessage) {
		t.Errorf("Контрольное сообщение и сообщение логгера для события HwSubmitter не совпадают")
	}

	//Тестируем логирование события HwAccepted
	checkedMessage = fmt.Sprintf(
		"%s accepted %d %d\n",
		time.Now().Format("2006-01-02"),
		EventId,
		GradeValue)

	buffer = &bytes.Buffer{}
	acceptedEvent := CreateHwAcceptedEvent(GradeValue)
	acceptedEvent.Id = EventId
	LogOtusEvent(acceptedEvent, buffer)

	if EqualValue != strings.Compare(buffer.String(), checkedMessage) {
		t.Errorf("Контрольное сообщение и сообщение логгера для события HwAccepted не совпадают")
	}

}
