package main

import(
"gorm.io/gorm"
	)	
type Nurse struct{
	gorm.Model
	FullName string `json:"fullname"`
	Email string `json:"email"`
	ContactNumber string `json:"contactnumber"`
	WorkingDays string `json:"workingdays"`
	DutyStartTime string `json:"dutystarttime"`
	DutyEndTime string `json:"dutyendtime"`
	Address string `json:"address"`
	Gender string `json:"gender"`
}