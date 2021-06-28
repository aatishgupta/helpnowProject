package dblogger

import (
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

type DBLogger struct {
	Logger  *log.Logger
	LogFile *os.File
}

func (dbLogger *DBLogger) GetLogger(fileobj *os.File) *DBLogger {
	logPath := "/var/log/helpnow/helpnow.log"
	var f1 *os.File
	var err error
	if fileobj == nil || reflect.TypeOf(fileobj) != reflect.TypeOf(f1) {
		f1, err = os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			fmt.Println("error opening file: %v", err)
			log.Fatalf("error opening file: %v", f1)
		}

	} else {
		f1 = fileobj
	}

	wrt := io.MultiWriter(f1)
	logger := log.New(wrt, "", log.LstdFlags)
	logger.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	dbLogger.Logger = logger
	dbLogger.LogFile = f1
	return dbLogger
}

func (dbLogger *DBLogger) Close() {
	dbLogger.LogFile.Close()
	dbLogger.Logger = &log.Logger{}
}
