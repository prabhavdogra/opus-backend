package utils

import (
	"fmt"
	"time"
)

func GetCurrentTimestamp() int64 {
	return time.Now().UTC().UnixNano() / int64(time.Millisecond)
}

func CalculateTimestampDifference(ts1, ts2 int64) int64 {
	return ts2 - ts1
}

func Must(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
