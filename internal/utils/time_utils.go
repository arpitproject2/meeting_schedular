package utils

import "time"

func IsTimeConflict(startTime, endTime, otherStartTime, otherEndTime time.Time) bool {
	return startTime.Before(otherEndTime) && endTime.After(otherStartTime)
}
