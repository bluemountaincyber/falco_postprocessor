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
func WriteToStdOut(output []byte) {
	output = append(output, '\n')
	if _, err := os.Stdout.Write(output); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to stdout: %v\n", err)
		os.Exit(1)
	}
}
