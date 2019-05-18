package koala

import (
	"fmt"
	"os"
	"strings"
)

func debugPrint(format string, a ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Fprintf(os.Stderr, "[KOALA-DEBUG] " + format, a)
}

func debugPrintf(format string, a interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Fprintf(os.Stdout, format, a)
}
