// builtins/repeat_test.go

package builtins

import (
	"bytes"
	
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepeatCommand(t *testing.T) {
	t.Parallel()

	t.Run("success: repeat command", func(t *testing.T) {
		w := &bytes.Buffer{}
		args := []string{"3", "echo", "hello"}

		err := RepeatCommand(w, args...)
		require.NoError(t, err, "unexpected error")

		expectedOutput := "hellohellohello"
		require.Equal(t, expectedOutput, w.String(), "unexpected output")
	})

	t.Run("failure: invalid repeat count", func(t *testing.T) {
		w := &bytes.Buffer{}
		args := []string{"invalid", "echo", "hello"}

		err := RepeatCommand(w, args...)
		require.Error(t, err, "expected an error for invalid repeat count")
		require.EqualError(t, err, ErrInvalidRepeatCount.Error(), "unexpected error message")
	})

	t.Run("failure: not enough arguments", func(t *testing.T) {
		w := &bytes.Buffer{}
		args := []string{"1"} // Not enough arguments
	
		err := RepeatCommand(w, args...)
		require.Error(t, err, "expected an error for not enough arguments")
		require.ErrorIs(t, err, ErrInvalidArgCount, "unexpected error type")
	})
}
