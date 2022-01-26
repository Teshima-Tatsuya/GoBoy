package debug

import (
	"fmt"
)

func Debug(s string, args ...interface{}) {
	if !DEBUG {
		return
	}

	fmt.Printf("[DEBUG]"+s, args...)
}

func Fatal(s string, args ...interface{}) {
	msg := fmt.Sprintf("[FATAL]"+s, args...)
	panic(msg)
}
