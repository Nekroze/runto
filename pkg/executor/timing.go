package executor

import (
	"fmt"
	"time"
)

type TimingStatistics struct {
	Fastest time.Duration
	Slowest time.Duration
	Average time.Duration
	history []time.Duration
}

func (t *TimingStatistics) String() string {
	return tabString([]string{
		fmt.Sprintf("Fastest:\t%s", t.Fastest),
		fmt.Sprintf("Slowest:\t%s", t.Slowest),
		fmt.Sprintf("Average:\t%s", t.Average),
	})
}

func (t *TimingStatistics) Record(in time.Duration) {
	t.updateMinMax(in)
	t.updateAverage(in)
}

func (t *TimingStatistics) updateMinMax(in time.Duration) {
	if t.Fastest == 0 || in < t.Fastest {
		t.Fastest = in
	}
	if t.Slowest == 0 || in > t.Slowest {
		t.Slowest = in
	}
}

func (t *TimingStatistics) updateAverage(in time.Duration) {
	t.history = append(t.history, in)
	t.Average = AverageDuration(t.history)
}

func AverageDuration(in []time.Duration) time.Duration {
	var total time.Duration
	for _, v := range in {
		total += v
	}
	return time.Duration(float64(total) / float64(len(in)))
}
