package eventlogger

import (
	"testing"
)

func TestEvent(t *testing.T) {
	var submitterStorage []*HwSubmitter
	var acceptedStorage []*HwAccepted

	submitterStorage = append(submitterStorage, createTestSubmitter("Some code of 1", "Comment of 1", 1111))
	submitterStorage = append(submitterStorage, createTestSubmitter("Some code of 2", "Comment of 2", 2222))
	submitterStorage = append(submitterStorage, createTestSubmitter("Some code of 3", "Comment of 3", 3333))
	for _, event := range submitterStorage {
		LogOtusEvent(event, os.Stdout)
	}

	acceptedStorage = append(acceptedStorage, createTestAccepted(4, 1111))
	acceptedStorage = append(acceptedStorage, createTestAccepted(5, 2222))
	for _, event := range acceptedStorage {
		LogOtusEvent(event, os.Stdout)
	}
}

func createTestSubmitter(code, comment string, id int) *HwSubmitter {
	result := CreateHwSubmitterEvent(code, comment)
	result.Id = id

	return result
}

func createTestAccepted(grade, id int) *HwAccepted {
	result := CreateHwAcceptedEvent(grade)
	result.Id = id

	return result
}