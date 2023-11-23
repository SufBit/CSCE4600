// builtins/touch_test.go
package builtins

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTouch(t *testing.T) {
	t.Parallel()

	t.Run("success: create files", func(t *testing.T) {
		fileNames := []string{"file1.txt", "file2.txt", "file3.txt"}
		err := Touch(fileNames...)
		require.NoError(t, err, "unexpected error")

		// Check if the files were actually created
		for _, fileName := range fileNames {
			_, err := os.Stat(fileName)
			require.NoError(t, err, "expected file not found")

			// Clean up: remove the created files
			err = os.Remove(fileName)
			require.NoError(t, err, "cleanup failed")
		}
	})

	t.Run("failure: missing operand", func(t *testing.T) {
		err := Touch()
		require.Error(t, err, "expected an error for missing operand")
		require.EqualError(t, err, ErrInvalidArgCount.Error()+": missing operand", "unexpected error message")
	})
}
