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
