// echo.go
package builtins

import (
	"fmt"
	"io"
	"strings"
)

func Echo(w io.Writer, args ...string) error {
	// Join the arguments with a space and write to the output.
	_, err := fmt.Fprintln(w, strings.Join(args, " "))
	return err
}
