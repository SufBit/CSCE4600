// echo_test.go
package builtins

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEchoCommand(t *testing.T) {
	t.Parallel()

	t.Run("echo command", func(t *testing.T) {
		w := &bytes.Buffer{}
		args := []string{"Hello,", "World!"}

		err := Echo(w, args...)
		require.NoError(t, err, "unexpected error")

		// Assert the expected output or any other validation
		require.Equal(t, "Hello, World!\n", w.String())
	})

	t.Run("echo command with multiple arguments", func(t *testing.T) {
		w := &bytes.Buffer{}
		args := []string{"This", "is", "a", "test."}

		err := Echo(w, args...)
		require.NoError(t, err, "unexpected error")

		// Assert the expected output or any other validation
		require.Equal(t, "This is a test.\n", w.String())
	})

	t.Run("echo command with no arguments", func(t *testing.T) {
		w := &bytes.Buffer{}
	
		err := Echo(w)
		require.NoError(t, err, "unexpected error")
	
		// Assert the expected output or any other validation
		require.Equal(t, "\n", w.String())
	})
	

	t.Run("echo command with special characters", func(t *testing.T) {
		w := &bytes.Buffer{}
		args := []string{"$100", "@username", "escaped\nnewline"}

		err := Echo(w, args...)
		require.NoError(t, err, "unexpected error")

		// Assert the expected output or any other validation
		require.Equal(t, "$100 @username escaped\nnewline\n", w.String())
	})
}
