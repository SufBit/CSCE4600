// builtins/ls_test.go
package builtins

import (
	"bytes"
	"os"
	"testing"
	"fmt"
	"github.com/stretchr/testify/require"
)

func TestListDirectory(t *testing.T) {
	t.Parallel()

	t.Run("success: list files in current directory", func(t *testing.T) {
		// Create temporary files in the current directory
		tempFiles := []string{"file1.txt", "file2.txt", "file3.txt"}
		for _, fileName := range tempFiles {
			_, err := os.Create(fileName)
			require.NoError(t, err, "failed to create temporary file")
			defer os.Remove(fileName)
		}

		// Run ListDirectory on the current directory
		var out bytes.Buffer
		err := ListDirectory(&out)
		require.NoError(t, err, "unexpected error")

		// Check if the output contains the temporary files
		for _, fileName := range tempFiles {
			require.Contains(t, out.String(), fileName, "output does not contain file name")
		}
	})

	t.Run("success: list files in a specific directory", func(t *testing.T) {
		// Create a temporary directory and files inside it
		tempDir := "temp_directory"
		err := os.Mkdir(tempDir, 0755)
		require.NoError(t, err, "failed to create temporary directory")
		defer os.Remove(tempDir)

		tempFiles := []string{"file1.txt", "file2.txt", "file3.txt"}
		for _, fileName := range tempFiles {
			_, err := os.Create(fmt.Sprintf("%s/%s", tempDir, fileName))
			require.NoError(t, err, "failed to create temporary file")
		}

		// Run ListDirectory on the specific directory
		var out bytes.Buffer
		err = ListDirectory(&out, tempDir)
		require.NoError(t, err, "unexpected error")

		// Check if the output contains the temporary files
		for _, fileName := range tempFiles {
			require.Contains(t, out.String(), fileName, "output does not contain file name")
		}
	})
}
