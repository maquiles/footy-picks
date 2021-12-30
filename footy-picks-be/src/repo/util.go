package repo

import "time"

func getCurrentTimestamp() string {
	return time.Now().Format("2006-01-02")
}
