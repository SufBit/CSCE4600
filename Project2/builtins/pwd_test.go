// builtins/pwd_test.go
package builtins

import (
	"bytes"
	"testing"
	"os"
	"fmt"
	"github.com/stretchr/testify/require"
)

func TestPrintWorkingDirectory(t *testing.T) {
	t.Parallel()

	t.Run("success: print working directory", func(t *testing.T) {
        // Run PrintWorkingDirectory
        var out bytes.Buffer
        err := PrintWorkingDirectory(&out)
        require.NoError(t, err, "unexpected error")

        // Print the current working directory for debugging
        wd, _ := os.Getwd()
        fmt.Println("Current working directory:", wd)

        // Check if the output contains the current working directory
        require.Contains(t, out.String(), wd, "output does not contain working directory")
    })
}
