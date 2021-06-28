package dbcore

import (
	"database/sql"
	"helpnow/src/bindings"
	"helpnow/src/dblogger"
	"helpnow/src/utils"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const APPOINTMENT = "appointment"
const SCHEDULE = "schedule"

var (
	logHandler dblogger.DBLogger
)

func DBConn() *sql.DB {
	db, err := sql.Open("sqlite3", "./helpnow.db")
	if err != nil {
		panic(err)
	}
	return db
}
func CreateScheduleTable() {
	scheduleTable := "CREATE TABLE IF NOT EXISTS " + SCHEDULE + " (id INTEGER PRIMARY KEY,appointment_start_time INTEGER NOT NULL,appointment_end_time INTEGER NOT NULL);"
	db := DBConn()
	statement, err := db.Prepare(scheduleTable) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("appointmet table created")
}
func CreateAppointmentTable() {
	appointmentTable := "CREATE TABLE IF NOT EXISTS " + APPOINTMENT + " (id INTEGER PRIMARY KEY,start_time INTEGER NOT NULL,end_time INTEGER NOT NULL);"
	db := DBConn()
	statement, err := db.Prepare(appointmentTable) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("appointmet table created")
}
func ListScheduledAppointment(dateParam string) []bindings.Schedule {
	db := DBConn()
	loggerStruct := logHandler.GetLogger(nil)
	logger := loggerStruct.Logger
	defer logHandler.Close()
	s := bindings.Schedule{}
	sArray := []bindings.Schedule{}
	dayStartTime := utils.GetTimeFromEpoch(dateParam, false)
	dayEndTime := utils.GetTimeFromEpoch(dateParam, true)
	logger.Println(dayStartTime, dayEndTime)
	filterQuery := "SELECT  datetime(start_time,'unixepoch', 'localtime') , datetime(end_time, 'unixepoch', 'localtime') FROM " + APPOINTMENT + " WHERE status='booked' and (start_time between ? and ? ) and (end_time between ? and ?)"
	statement, err := db.Prepare(filterQuery)

	logger.Println(filterQuery)
	if err != nil {
		logger.Println(err)
	}
	rows, err := statement.Query(dayStartTime, dayEndTime, dayStartTime, dayEndTime)
	if err != nil {
		logger.Println(err)
	}

	var startDateTime string
	var endDateTime string
	defer rows.Close()
	if rows != nil {
		for rows.Next() {
			err = rows.Scan(&startDateTime, &endDateTime)
			if err == nil {
				s.StartTime = startDateTime
				s.EndTime = endDateTime
				sArray = append(sArray, s)
			} else {
				logger.Println(err)
			}
		}

	}
	return sArray

}
func IsAppointmentAvailable(startTime int64) bool {
	db := DBConn()
	loggerStruct := logHandler.GetLogger(nil)
	logger := loggerStruct.Logger
	defer logHandler.Close()
	var output int
	row, err := db.Prepare("SELECT COUNT(start_time) FROM " + APPOINTMENT + " WHERE status = ? and (? between start_time and end_time)")
	if err != nil {
		logger.Println(err)
		return false
	}
	defer row.Close()
	err = row.QueryRow("booked", startTime).Scan(&output)
	if err != nil {
		return false
	}
	if output > 0 {
		return false
	}
	return true
}
func BookAppointment(startTime int64, endTime int64) bool {
	db := DBConn()
	loggerStruct := logHandler.GetLogger(nil)
	logger := loggerStruct.Logger
	defer logHandler.Close()
	returnVal := true
	insertAppointment := "INSERT INTO " + APPOINTMENT + "(start_time, end_time, status) VALUES (?, ?, ?)"
	statement, err := db.Prepare(insertAppointment)
	if err != nil {
		logger.Println(err.Error())
		returnVal = false
	}
	_, err = statement.Exec(startTime, endTime, "booked")
	if err != nil {
		logger.Println(err.Error())
		returnVal = false
	}
	return returnVal
}
func CancelAppointment(id int) bool {
	db := DBConn()
	loggerStruct := logHandler.GetLogger(nil)
	logger := loggerStruct.Logger
	defer logHandler.Close()
	returnVal := true
	cancelAppointment := "UPDATE " + APPOINTMENT + " SET status = ? WHERE id= ? ;"
	statement, err := db.Prepare(cancelAppointment)
	if err != nil {
		logger.Println(err.Error())
		logger.Println(id)
		logger.Println(err.Error())
		returnVal = false
	}
	_, err = statement.Exec("cancelled", id)
	if err != nil {
		logger.Println(err.Error())
		returnVal = false
	}
	return returnVal
}
func AddDoctorScheduleTime(startTime int64, endTime int64) bool {
	db := DBConn()
	loggerStruct := logHandler.GetLogger(nil)
	logger := loggerStruct.Logger
	defer logHandler.Close()
	returnVal := true
	insertAppointmentTime := "INSERT INTO " + SCHEDULE + "(appointment_start_time, appointment_end_time) VALUES (?, ?)"
	statement, err := db.Prepare(insertAppointmentTime)
	if err != nil {
		logger.Println(err.Error())
		returnVal = false
	}
	_, err = statement.Exec(startTime, endTime)
	if err != nil {
		logger.Println(err.Error())
		returnVal = false
	}
	return returnVal
}
