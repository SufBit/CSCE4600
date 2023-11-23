// builtins/ls.go
package builtins

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ListDirectory(w io.Writer, args ...string) error {
	dirPath := "."
	if len(args) > 0 {
		dirPath = args[0]
	}

	dir, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	var fileNames []string
	for _, entry := range dir {
		fileNames = append(fileNames, entry.Name())
	}

	_, err = fmt.Fprintln(w, strings.Join(fileNames, "\t"))
	return err
}
