package builtins

import (
    "fmt"
    "os"
)

func Touch(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("%w: missing operand", ErrInvalidArgCount)
	}

	for _, fileName := range args {
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	return nil
}