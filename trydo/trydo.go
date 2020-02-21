package trydo

import (
	"fmt"
	"time"
)

func TryTimes(do DoFn, times int) error {
	var retErr error
	for tryTime := 1; tryTime <= times; tryTime++ {
		err := do()
		if err != nil {
			fmtErr := fmt.Errorf("try_time:%v with err:%v", tryTime, err)
			if retErr == nil {
				retErr = fmtErr
			} else {
				retErr = fmt.Errorf("%v; %v", fmtErr, retErr)
			}
		} else {
			return nil
		}
	}
	return retErr
}

func TryWithIntervals(do DoFn, intervals ...time.Duration) error {
	times := len(intervals) + 1
	prependedIntervals := append([]time.Duration{0}, intervals...)
	var tryTime int
	wrappedDo := func() error {
		time.Sleep(prependedIntervals[tryTime])
		tryTime++
		return do()
	}
	return TryTimes(wrappedDo, times)
}
