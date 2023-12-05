// which_test.go
package builtins

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWhichCommand(t *testing.T) {
	t.Parallel()

	t.Run("which command with an existing executable", func(t *testing.T) {
		w := &bytes.Buffer{}
		args := []string{"ls"} // Assuming 'ls' is a common executable on Unix-like systems

		err := Which(w, args...)
		require.NoError(t, err, "unexpected error")

		// Assert the expected output or any other validation
		// Note: Adjust the expected output based on your system and available executables
		require.Contains(t, w.String(), "/bin/ls\n")
	})

	t.Run("which command with a non-existing executable", func(t *testing.T) {
		w := &bytes.Buffer{}
		args := []string{"nonexistent_executable"}

		err := Which(w, args...)
		require.NoError(t, err, "unexpected error")

		// Assert the expected output or any other validation
		require.Empty(t, w.String())
	})

	t.Run("which command with multiple existing executables", func(t *testing.T) {
		w := &bytes.Buffer{}
		args := []string{"ls", "pwd"}

		err := Which(w, args...)
		require.NoError(t, err, "unexpected error")

		// Assert the expected output or any other validation
		// Note: Adjust the expected output based on your system and available executables
		require.Contains(t, w.String(), "/bin/ls\n")
		require.Contains(t, w.String(), "/bin/pwd\n")
	})

	t.Run("which command with multiple executables, some non-existing", func(t *testing.T) {
		w := &bytes.Buffer{}
		args := []string{"ls", "nonexistent_executable", "pwd"}
	
		err := Which(w, args...)
		require.NoError(t, err, "unexpected error")
	
		// Assert the expected output or any other validation
		expectedOutput := "/bin/ls\n/bin/pwd\n"
		require.Equal(t, expectedOutput, w.String(), "unexpected output")
	})
	
}
