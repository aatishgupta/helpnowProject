package utils

import (
	"helpnow/src/dblogger"
	"time"
)

var (
	logHandler dblogger.DBLogger
)

func GetEpochTime() int64 {
	return time.Now().Unix()
}
func GetTimeFromEpoch(dateParam string, isEndTime bool) int64 {
	loggerStruct := logHandler.GetLogger(nil)
	logger := loggerStruct.Logger
	defer logHandler.Close()
	layout := "2006-01-02"
	thetime, err := time.Parse(layout, dateParam)
	if err != nil {
		panic("Can't parse time format")
	}
	logger.Println(thetime.Unix())
	t := thetime.Unix() - 19800
	if isEndTime {
		t = t + 86399
	}
	return t
}
