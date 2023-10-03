package utils

import "time"


func DateTimeFormatter(date string) (result string, err error) {
	// format := "2006-01-02 15:04:05"
	t, errParse := time.Parse("2006-01-02 15:04:05", date)
	return t.String(), errParse
}

func TimeFormatter(date string) (result string, err error) {
	format := "15:04:05"
    resultParse, errParse := time.Parse(format, date)
	return resultParse.String(), errParse
}

func GetDifferenceTime(dateStr1, dateStr2 string) (years, months, days, hours, minutes, seconds int64) {
	format := "2006-01-02 15:04:05"
	date1,_ := time.Parse(format, dateStr1)
	date2,_ := time.Parse(format, dateStr2)

	if date2.After(date1) {
		date1, date2 = date2, date1
	}

	difference := date1.Sub(date2)
	years = int64(difference.Hours()/24/365)
	months = int64(difference.Hours()/24/30)
	days = int64(difference.Hours()/24)
	hours = int64(difference.Hours())
	minutes = int64(difference.Minutes())
	seconds = int64(difference.Seconds())
	return
}