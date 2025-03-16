package cli

import (
	"fmt"
	"os"
)

func ErrorExit(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "ERROR: "+format+"\n", args...)
	os.Exit(1)
}

func PrintDryRun(msg string) {
	fmt.Printf("[DRY-RUN] %s\n", msg)
}

func PrintDebug(format string, args ...interface{}) {
	fmt.Printf("[DEBUG] "+format+"\n", args...)
}
