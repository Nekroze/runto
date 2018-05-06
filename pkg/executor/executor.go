package executor

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

type ExecutionResult struct {
	Command  string
	Output   string
	ExitCode int
	Duration time.Duration
}
type Executor func(string) *ExecutionResult

func Call(command string) (out *ExecutionResult) {
	out = &ExecutionResult{Command: command}
	defer func(start time.Time) {
		out.Duration = time.Since(start)
	}(time.Now())

	cparts := strings.Fields(command)
	head, tail := cparts[0], cparts[1:]
	stdoe, err := exec.Command(head, tail...).CombinedOutput()

	out.Output = string(stdoe)

	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				out.ExitCode = status.ExitStatus()
			}
		} else {
			panic(fmt.Errorf("Failed to execute os command due to error: %s", err))
		}
	}
	return out
}
