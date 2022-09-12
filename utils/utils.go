package utils

import (
	"time"
)

func CurrentTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	now := time.Now().In(loc)
	return now
}
