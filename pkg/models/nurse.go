package models

import (
	"gorm.io/gorm"
)

type Nurse struct {
	gorm.Model
	FullName           string `json:"full_name" validate:"required min=2 max=30"`
	Email              string `json:"email" validate:"email,required"`
	ContactNumber      string `json:"contact_number" validate:"required min=8 max=10"`
	WorkingDays        string `json:"working_days" validate:"required"`
	WorkingShift       string `json:"working_shift" validate:"required , eq=morning|eq=day|eq=night"`
	DutyStartTime      string `json:"duty_start_time" validate:"required"`
	DutyEndTime        string `json:"duty_end_time" validate:"required"`
	Address            string `json:"address"`
	Gender             string `json:"gender" validate:"eq:Male | eq:Female | eq:Other"`
	CreatedAtFormatted string `json:"created_at_formatted"`
}
