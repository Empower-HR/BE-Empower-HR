package dataleaves

import (
	"time"

	"gorm.io/gorm"
)

type LeavesData struct {
	gorm.Model
	StartDate      time.Time
	EndDate        time.Time
	Reason         string
	Status         string
	TotalLeave     int
	PersonalDataID uint
}
