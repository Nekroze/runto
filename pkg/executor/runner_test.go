package executor

import "testing"

func TestRunner_Loop(t *testing.T) {
	type fields struct {
		Command   string
		Exec      Executor
		Condition RunnerStopCheck
	}
	tests := []struct {
		name             string
		fields           fields
		expectedAttempts uint
	}{
		{"Immediate End By Success", fields{
			Command:   "test",
			Exec:      genTestExecutor([]int{0}),
			Condition: StopOnSuccess,
		}, 1},
		{"Single Retry Then End By Success", fields{
			Command:   "test",
			Exec:      genTestExecutor([]int{1, 0}),
			Condition: StopOnSuccess,
		}, 2},
		{"Immediate End By Failure", fields{
			Command:   "test",
			Exec:      genTestExecutor([]int{1}),
			Condition: StopOnFailure,
		}, 1},
		{"Single Retry Then End By Failure", fields{
			Command:   "test",
			Exec:      genTestExecutor([]int{0, 1}),
			Condition: StopOnFailure,
		}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Runner{
				Command:   tt.fields.Command,
				Exec:      tt.fields.Exec,
				Condition: tt.fields.Condition,
			}
			r.Loop()
			if r.result.Attempts != tt.expectedAttempts {
				t.Fatalf("%d attempts instead of expcted %d", r.result.Attempts, tt.expectedAttempts)
			}
		})
	}
}
