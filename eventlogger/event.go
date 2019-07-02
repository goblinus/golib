package eventlogger

import (
	"fmt"
	"time"
)

type OtusEvent interface {
	LogMessage() string
}

type Event struct {
	Id       int
	Name     string
	DateTime time.Time
}

func (e *Event) LogMessage() string {
	return fmt.Sprintf(
		"%s %s %d",
		e.DateTime.Format("2006-01-02"),
		e.Name,
		e.Id)
}

func CreateHwAcceptedEvent(grade int) *HwAccepted {
	return &HwAccepted{
		Event{
			Name:     "accepted",
			DateTime: time.Now()},
		grade}
}

func CreateHwSubmitterEvent(code, comment string) *HwSubmitter {
	return &HwSubmitter{
		Event{
			Name:     "commiter",
			DateTime: time.Now()},
		code,
		comment}
}

type HwAccepted struct {
	Event
	Grade int
}

func (ha *HwAccepted) LogMessage() string {
	return fmt.Sprintf(
		"%s %s %d %d",
		ha.DateTime.Format("2006-01-02"),
		ha.Name,
		ha.Id,
		ha.Grade)
}

type HwSubmitter struct {
	Event
	Code    string
	Comment string
}

func (hs *HwSubmitter) LogMessage() string {
	return fmt.Sprintf(
		"%s %s %d \"%s\"",
		hs.DateTime.Format("2006-01-02"),
		hs.Name,
		hs.Id,
		hs.Comment)
}
