// builtins/type_test.go

package builtins

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTypeCommand(t *testing.T) {
	t.Parallel()

	t.Run("built-in command", func(t *testing.T) {
		w := &bytes.Buffer{}
		err := TypeCommand(w, "cd")
		require.NoError(t, err, "unexpected error")

		expectedOutput := "cd is a built-in command\n"
		require.Equal(t, expectedOutput, w.String(), "unexpected output")
	})

	t.Run("external command", func(t *testing.T) {
		w := &bytes.Buffer{}
		command := "ls"

		err := TypeCommand(w, command)
		require.NoError(t, err, "unexpected error")

		// Check if the actual output contains the expected command and path
		expectedOutput := command + " is a file located at /bin/ls\n"
		actualOutput := w.String()

		// Use strings.Contains to check if the actual output contains the expected substring
		require.Contains(t, actualOutput, expectedOutput, "unexpected output")
	})

	t.Run("non-existent command", func(t *testing.T) {
		w := &bytes.Buffer{}
		err := TypeCommand(w, "nonexistentcommand")
		require.Error(t, err, "expected an error for a non-existent command")

		expectedError := "type: nonexistentcommand: not found"
		require.EqualError(t, err, expectedError, "unexpected error message")
	})
}
