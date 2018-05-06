package executor

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Runner struct {
	Command   string
	Exec      Executor
	Condition RunnerStopCheck
	result    *RunnerResult
}

func (r *Runner) exec() {
	log.Printf("%s iteration %d\n", os.Args[0], r.result.Attempts)
	result := r.Exec(r.Command)
	r.result.Record(result)
	if len(strings.TrimSpace(result.Output)) > 0 {
		fmt.Println(result.Output)
	}
}

func (r *Runner) Loop() *RunnerResult {
	r.result = &RunnerResult{
		Command:     r.Command,
		Performance: &TimingStatistics{},
	}
	for {
		r.exec()
		if r.Condition(r.result) {
			return r.result
		}
	}
}
