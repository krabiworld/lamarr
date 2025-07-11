package uptime

import "time"

var startTime time.Time

func Init() {
	startTime = time.Now()
}

func Get() int64 {
	return startTime.Unix()
}
