package dateutils

import (
	"time"
)

const (
	//01 indicate months, 02 indicate days
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNowDbFormat()string{
	return GetNow().Format(apiDbLayout)
}
