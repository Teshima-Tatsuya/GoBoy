package debug

import (
	"fmt"
)

func Debug(s string, args ...interface{}) {
	if !DEBUG {
		return
	}

	fmt.Printf("[DEBUG]"+s+"\n", args...)
}

func Info(s string, args ...interface{}) {
	fmt.Printf("[INFO]"+s+"\n", args...)
}

func Fatal(s string, args ...interface{}) {
	msg := fmt.Sprintf("[FATAL]"+s, args...)
	panic(msg)
}
