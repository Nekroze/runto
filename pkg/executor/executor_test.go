package executor

import "time"

func genTestExecutor(exitSequence []int) Executor {
	pos := 0
	return func(cmd string) *ExecutionResult {
		out := &ExecutionResult{
			Command:  cmd,
			ExitCode: exitSequence[pos],
			Duration: time.Second,
		}
		pos++
		if pos == len(exitSequence) {
			pos = 0
		}
		return out
	}
}
