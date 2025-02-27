package outputs

import (
	"fmt"
	"os"
)

// WriteToStdOut writes the output to stdout
//
// The input to this function is a byte slice representing the output.
//
// An expected usage might be:
//
//	WriteToStdOut(output)
func WriteToFile(LOGFILE string, output []byte) {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening log file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	if _, err := f.WriteString(string(output) + "\n"); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to log file: %v\n", err)
		os.Exit(1)
	}
}
