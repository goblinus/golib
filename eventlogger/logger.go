package eventlogger

import (
	"fmt"
	"io"
)

func LogOtusEvent(event OtusEvent, w io.Writer) {
	fmt.Fprintln(w, event.LogMessage())
}
