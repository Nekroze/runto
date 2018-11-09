package executor

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

type RunnerResult struct {
	Command      string
	Attempts     uint
	Performance  *TimingStatistics
	LastExitCode int
}

func (rr *RunnerResult) String() string {
	return tabString([]string{
		"# Command",
		fmt.Sprintf("Command:\t%s", rr.Command),
		fmt.Sprintf("Attempts:\t%d", rr.Attempts),
		fmt.Sprintf("Last Exit:\t%d", rr.LastExitCode),
		"# Performance",
		rr.Performance.String(),
	})
}

func (rr *RunnerResult) Record(result *ExecutionResult) {
	rr.Attempts++
	rr.Performance.Record(result.Duration)
	rr.LastExitCode = result.ExitCode
}

func tabString(lines []string) string {
	buf := new(bytes.Buffer)
	w := tabwriter.NewWriter(buf, 0, 0, 1, ' ', 0)
	for _, line := range lines {
		fmt.Fprintf(w, "%s\n", line)
	}
	w.Flush()
	return buf.String()
}
