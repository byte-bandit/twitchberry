package stream

import (
	"fmt"
	"os/exec"
)

type Stream struct {
	cmd *exec.Cmd
}

func New() *Stream {
	return &Stream{}
}

func (s *Stream) Start() error {
	if s.cmd != nil {
		return fmt.Errorf("stream already started")
	}

	s.cmd = exec.Command("paint.exe")
	return s.cmd.Start()
}

func (s *Stream) Stop() error {
	if s.cmd == nil {
		return fmt.Errorf("stream not running")
	}

	return s.cmd.Wait()
}
