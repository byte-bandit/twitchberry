package stream

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Stream(t *testing.T) {
	t.Run("Should open notepad", func(t *testing.T) {
		c := exec.Command("notepad.exe")
		require.NoError(t, c.Start())
		require.NoError(t, c.Process.Kill())
		require.NoError(t, c.Wait())
	})
}
