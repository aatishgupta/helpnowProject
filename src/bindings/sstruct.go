package bindings

type Appointment struct {
	StartTime int64 `json:"start_time"`
	EndTime   int64 `json:"end_time"`
}
type Schedule struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
type DoctorSchedule struct {
	StartAppointmentTime string `json:"start_appointment_time"`
	EndAppointmentTime   string `json:"end_appointment_time"`
}
