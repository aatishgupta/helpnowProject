package handler

import (
	"helpnow/src/bindings"
	"helpnow/src/dbcore"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListScheduledAppointment(c *gin.Context) {
	queryDate := c.Query("date")
	if queryDate == "" {
		userResponse := make(map[string]interface{})
		userResponse["message"] = "Please provide date in Y-m-d format"
		userResponse["status"] = false
		c.JSON(http.StatusBadRequest, userResponse)
		return
	}
	schedule := dbcore.ListScheduledAppointment(queryDate)
	c.JSON(http.StatusOK, schedule)

}
func ScheduleTimeSlot(c *gin.Context) {
	appointmentSchedulePayload := bindings.DoctorSchedule{}
	userResponse := make(map[string]interface{})
	if err := c.ShouldBindJSON(&appointmentSchedulePayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	startTime, endTime, err := DoctorScheduleValidation(appointmentSchedulePayload)
	if err != nil {
		userResponse["message"] = err.Error()
		userResponse["status"] = false
		c.JSON(http.StatusBadRequest, userResponse)
		return

	} else {
		if dbcore.AddDoctorScheduleTime(startTime, endTime) {
			userResponse["message"] = "Your daily schedule timing added in db"
			userResponse["status"] = true
			c.JSON(http.StatusOK, userResponse)
			return

		} else {
			userResponse["message"] = "Some Error is there try after some time"
			userResponse["status"] = false
			c.JSON(http.StatusBadRequest, userResponse)
			return
		}

	}

}

func BookAppointment(c *gin.Context) {
	appointmentPayload := bindings.Appointment{}
	httpCode := http.StatusOK
	userResponse := make(map[string]interface{})
	if err := c.ShouldBindJSON(&appointmentPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := AppointmentValidation(appointmentPayload)
	if err != nil {
		userResponse["message"] = err.Error()
		userResponse["status"] = false
	} else {
		if dbcore.BookAppointment(appointmentPayload.StartTime, appointmentPayload.EndTime) {
			userResponse["message"] = "Your Appointment booked successfully"
			userResponse["status"] = true

		} else {
			userResponse["message"] = "There is problem try after some time"
			userResponse["status"] = false
		}

	}
	c.JSON(httpCode, userResponse)

}
func CancelAppointment(c *gin.Context) {
	id := c.Param("id")
	httpCode := http.StatusOK
	userResponse := make(map[string]interface{})
	if id == "" {
		httpCode = http.StatusBadRequest
		userResponse["message"] = "Please provide id to cancel appointment"
		userResponse["status"] = false
	} else {
		id, err := strconv.Atoi(id)
		if err != nil {
			httpCode = http.StatusBadRequest
			userResponse["message"] = "Please provide valid id to cancel appointment"
			userResponse["status"] = false

		} else {
			if dbcore.CancelAppointment(id) {
				userResponse["message"] = "Your Appointment cancelled successfully"
				userResponse["status"] = true

			} else {
				userResponse["message"] = "Error while updating your appointment"
				userResponse["status"] = false
				httpCode = http.StatusBadRequest
			}
		}
	}
	c.JSON(httpCode, userResponse)

}
