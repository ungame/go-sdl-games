package global

import "time"

const TARGET_TICKS_PER_SECOND = 60

var delta float64

func GetDeltaTime() float64 {
	return delta
}

func SetDeltaTime(frameStart time.Time) (elapsed time.Duration) {
	elapsed = time.Since(frameStart)
	delta = elapsed.Seconds() * TARGET_TICKS_PER_SECOND
	return elapsed
}
