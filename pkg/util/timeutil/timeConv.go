package timeutil

import "time"

// was used to parse in localtime, but now used as usual parsing
// to not break already existed dependencies
func ParseLocalTime(t string, format string) (time.Time, error) {
	// loc, err := time.LoadLocation("Asia/Makassar")
	// if err != nil {
	// 	return time.Time{}, err
	// }
	// return time.ParseInLocation(format, t, loc)
	return time.Parse(format, t)
}

func FormatLocalTime(t time.Time, format string) string {
	// loc, _ := time.LoadLocation("Asia/Makassar")
	// return t.In(loc).Format(format)
	return t.Format(format)
}
