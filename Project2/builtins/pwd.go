// builtins/pwd.go
package builtins

import (
	"fmt"
	"io"
	"os"
)

// PrintWorkingDirectory prints the current working directory to the specified writer.
func PrintWorkingDirectory(w io.Writer) error {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Print the working directory to the specified writer
	_, err = fmt.Fprintln(w, wd)
	return err
}
