// builtins/rm_test.go
package builtins

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDirectory(t *testing.T) {
	t.Parallel()

	t.Run("success: remove directory", func(t *testing.T) {
		// Create a temporary directory and files inside it
		tempDir := "temp_directory"
		err := os.Mkdir(tempDir, 0755)
		require.NoError(t, err, "failed to create temporary directory")
		defer os.Remove(tempDir)

		// Run RemoveDirectory on the specific directory
		err = RemoveDirectory(tempDir)
		require.NoError(t, err, "unexpected error")

		// Check if the directory was actually removed
		_, err = os.Stat(tempDir)
		require.Error(t, err, "expected directory not found")
	})

	t.Run("failure: missing operand", func(t *testing.T) {
		// Run RemoveDirectory without providing a directory
		err := RemoveDirectory()
		require.Error(t, err, "expected an error for missing operand")
		require.EqualError(t, err, ErrInvalidArgCount.Error()+": missing operand", "unexpected error message")
	})
}
