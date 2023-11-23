package builtins

import (
    "fmt"
    "os"
)

func MakeDirectory(args ...string) error {
    if len(args) == 0 {
        return fmt.Errorf("%w: missing operand", ErrInvalidArgCount)
    }

    err := os.Mkdir(args[0], 0755) // 0755 is a common directory permission.
    return err
}
