// which.go
package builtins

import (
	"fmt"
	"io"
	"os/exec"
)

func Which(w io.Writer, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("%w: missing operand", ErrInvalidArgCount)
	}

	for _, cmd := range args {
		path, err := exec.LookPath(cmd)
		if err == nil {
			fmt.Fprintln(w, path)
		}
	}

	return nil
}
