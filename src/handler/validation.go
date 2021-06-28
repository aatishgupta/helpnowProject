package handler

import (
	"errors"
	"helpnow/src/bindings"
	"helpnow/src/dbcore"
	"helpnow/src/dblogger"
	"helpnow/src/utils"
	"strconv"
	"strings"
)

var (
	logHandler dblogger.DBLogger
)

func DoctorScheduleValidation(payload bindings.DoctorSchedule) (int64, int64, error) {
	sArray := strings.Split(payload.StartAppointmentTime, ":")
	eArray := strings.Split(payload.EndAppointmentTime, ":")
	if len(sArray) > 2 || len(eArray) > 2 {
		return 0, 0, errors.New("Please provide time ins HH:mm format")
	}
	h, err := strconv.Atoi(sArray[0])
	if err != nil {
		return 0, 0, err
	}
	m, err := strconv.Atoi(sArray[1])
	if err != nil {
		return 0, 0, err
	}
	eh, err := strconv.Atoi(eArray[0])
	if err != nil {
		return 0, 0, err
	}
	em, err := strconv.Atoi(eArray[1])
	if err != nil {
		return 0, 0, err
	}

	startTime := int64(h*60*60 + m*60)
	endTime := int64(eh*60*60 + em*60)
	return startTime, endTime, nil

}
func AppointmentValidation(payload bindings.Appointment) error {
	loggerStruct := logHandler.GetLogger(nil)
	logger := loggerStruct.Logger
	defer logHandler.Close()

	epochTime := utils.GetEpochTime()
	if payload.StartTime < epochTime {
		return errors.New("Start time must be greater than current time")
	}
	timeDiff := payload.EndTime - payload.StartTime
	totalAppointMentTime := int64(900) // which 15 minutes
	logger.Println("HEEEEEE")
	logger.Println(timeDiff)
	logger.Println(totalAppointMentTime)

	if timeDiff != totalAppointMentTime {
		return errors.New("Appointment time must be 15 minutes only")
	}
	if !dbcore.IsAppointmentAvailable(payload.StartTime) {
		return errors.New("Appointment for this slot is not available")
	}
	return nil
}
