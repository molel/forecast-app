package utils

import "time"

func FormatTs(ts int64) string {
	return time.Unix(0, ts).Format("02 Jan 2006")
}
