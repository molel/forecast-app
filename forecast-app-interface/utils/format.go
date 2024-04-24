package utils

import (
	"fmt"
	"time"
)

func FormatTs(ts int64, period int32) string {
	switch period {
	case 24:
		return time.Unix(0, ts).Format("02 Jan 2006")
	case 7:
		year, week := time.Unix(0, ts).ISOWeek()
		return fmt.Sprintf("Week %d of %d", week, year)
	case 4:
		return time.Unix(0, ts).Format("Jan 2006")
	case 12:
		return time.Unix(0, ts).Format("2006")
	}
	return ""
}
