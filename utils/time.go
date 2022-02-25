package utils

import "time"

// FormatTime formats given time to as string format.
func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// ParseTime parses the formatted string time.
func ParseTime(s string) time.Time {
	t, _ := time.Parse(time.RFC3339, s)
	return t
}

//UnixToTime converts a epoch timestamp to a time object
func UnixToTime(ms uint64) time.Time {
	return time.Unix(0, int64(ms)*int64(time.Millisecond))
}

const invertTsNum = 10000000000

// InvertTimestamp makes large numbers smaller.
func InvertTimestamp(ts int64) int64 {
	return invertTsNum - ts
}
