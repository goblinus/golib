package eventlogger

import (
	"io"
	"fmt"
)

func LogOtusEvent(event OtusEvent, w io.Writer) {
	fmt.Fprintf(w, event.LogMessage())
}