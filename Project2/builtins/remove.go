package builtins

import (
	"fmt"
	"os"
)

func RemoveDirectory(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("%w: missing operand", ErrInvalidArgCount)
	}

	for _, dir := range args {
		err := os.RemoveAll(dir)
		if err != nil {
			return err
		}
	}

	return nil
}