// mkdir_test.go
package builtins

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeDirectory(t *testing.T) {
	t.Parallel()

	t.Run("success: create directory", func(t *testing.T) {
		dirName := "test_directory"
		err := MakeDirectory(dirName)
		require.NoError(t, err, "unexpected error")

		// Check if the directory was actually created
		_, err = os.Stat(dirName)
		require.NoError(t, err, "expected directory not found")

		// Clean up: remove the created directory
		err = os.Remove(dirName)
		require.NoError(t, err, "cleanup failed")
	})

	t.Run("failure: missing operand", func(t *testing.T) {
		err := MakeDirectory()
		require.Error(t, err, "expected an error for missing operand")
		require.EqualError(t, err, ErrInvalidArgCount.Error()+": missing operand", "unexpected error message")
	})

	t.Run("failure: invalid directory name", func(t *testing.T) {
		invalidDirName := "/invalid/directory/name"
		err := MakeDirectory(invalidDirName)
		require.Error(t, err, "expected an error for invalid directory name")
	})
}
