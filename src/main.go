package main

import (
	"helpnow/src/dbcore"
	"helpnow/src/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	dbcore.CreateAppointmentTable()
	dbcore.CreateScheduleTable()

	router := gin.Default()
	v1 := router.Group("/v1/appointment")
	v1.POST("/book", handler.BookAppointment)
	v1.GET("/list", handler.ListScheduledAppointment)
	v1.POST("/schedule", handler.ScheduleTimeSlot)
	v1.PUT("/cancel/:id", handler.CancelAppointment)
	router.Run(":8080")
}
