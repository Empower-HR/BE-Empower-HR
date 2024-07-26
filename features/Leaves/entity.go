package leaves

import "time"

type LeavesDataEntity struct {
	LeavesID       uint
	StartDate      time.Time
	EndDate        time.Time
	Reason         string
	Status         string
	TotalLeave     int
	PersonalDataID uint
}
