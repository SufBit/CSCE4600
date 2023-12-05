// builtins/repeat.go

package builtins

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"strconv"
)

// ErrInvalidRepeatCount is an error for invalid repeat count.
var ErrInvalidRepeatCount = fmt.Errorf("%w: invalid repeat count", ErrInvalidArgCount)

// RepeatCommand implements the 'repeat' shell builtin.
func RepeatCommand(w io.Writer, args ...string) error {
	if len(args) < 2 {
		return fmt.Errorf("%w: expected at least two arguments", ErrInvalidArgCount)
	}

	count, err := strconv.Atoi(args[0])
	if err != nil || count <= 0 {
		return ErrInvalidRepeatCount
	}

	cmdArgs := args[1:]

	for i := 0; i < count; i++ {
		// Run the command and capture the output.
		output, err := runCommand(cmdArgs...)
		if err != nil {
			return err
		}

		// Print the command output to the specified writer.
		fmt.Fprint(w, output)
	}

	return nil
}

// runCommand executes a shell command and returns its output.
func runCommand(args ...string) (string, error) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

