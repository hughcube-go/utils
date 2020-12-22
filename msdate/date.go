package msdate

import "time"

const YYYYMMDDHHIISS = "2006-01-02 15:04:05"
const YYYYMMDDHHIISSMST = "2006-01-02 15:04:05.999999999 z07:00"

func Parse(value string) (time.Time, error) {
	return time.ParseInLocation(YYYYMMDDHHIISS, value, time.Local)
}

func Format(t time.Time) string {
	return t.Format(YYYYMMDDHHIISS)
}

func ParseMst(value string) (time.Time, error) {
	return time.ParseInLocation(YYYYMMDDHHIISSMST, value, time.Local)
}

func FormatMst(t time.Time) string {
	return t.Format(YYYYMMDDHHIISSMST)
}
