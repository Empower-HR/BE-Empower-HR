package service

import (
	leaves "be-empower-hr/features/Leaves"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

type leavesService struct {
	leavesData leaves.DataLeavesInterface
}

func New(ld leaves.DataLeavesInterface) leaves.ServiceLeavesInterface {
	return &leavesService{
		leavesData: ld,
	}
}

// UpdateLeaveStatus implements leaves.ServiceLeavesInterface.
func (s *leavesService) UpdateLeaveStatus(userID uint, leaveID uint, updatesleaves leaves.LeavesDataEntity) error {
	var user leaves.PersonalDataEntity
	if err := s.leavesData.GetPersonalDataByID(userID, &user); err != nil {
		return err
	}

	// Validate the status value
	if updatesleaves.Status != "" && updatesleaves.Status != "approved" && updatesleaves.Status != "rejected" {
		return errors.New("invalid status")
	}

	// Fetch leave details
	leaveDetail, err := s.leavesData.GetLeavesDetail(leaveID)
	if err != nil {
		log.Println("Error fetching leave details:", err)
		return err
	}

	// Calculate leave days
	leaveDays, err := CalculateLeaveDays(leaveDetail.StartDate, leaveDetail.EndDate)
	if err != nil {
		log.Println("Error calculating leave days:", err)
		return err
	}

	log.Printf("UserID: %d, LeaveID: %d, Status: %s, Requested Leave Days: %d", userID, leaveID, updatesleaves.Status, leaveDays)

	// Check if the leave request is approved and adjust the total leave balance
	if updatesleaves.Status == "approved" {
		var leaveData leaves.LeavesDataEntity
		if err := s.leavesData.GetLeavesDataByID(leaveDetail.LeavesID, &leaveData); err != nil {
			log.Println("Error fetching leave data:", err)
			return err
		}

		log.Printf("UserID: %d, Current Total Leave: %d, Requested Leave Days: %d", userID, leaveData.TotalLeave, leaveDays)

		// Check if the user has sufficient leave balance
		if leaveData.TotalLeave < leaveDays {
			log.Println("Insufficient leave balance")
			return fmt.Errorf("insufficient leave balance: you have %d days left but requested %d days", leaveData.TotalLeave, leaveDays)
		}

		// Update the total leave balance
		leaveData.TotalLeave -= leaveDays
		if err := s.leavesData.UpdateLeaveData(leaveData); err != nil {
			log.Println("Error updating leave data:", err)
			return err
		}
	}

	// Update the leave status
	if err := s.leavesData.UpdateLeaveStatus(leaveID, updatesleaves); err != nil {
		log.Println("Error updating leave status:", err)
		return err
	}

	return nil
}

// RequestLeave implements leaves.ServiceLeavesInterface.
func (s *leavesService) RequestLeave(userID uint, leave leaves.LeavesDataEntity) error {
	err := s.leavesData.RequestLeave(leave)
	if err != nil {
		log.Println("Error requesting leave:", err)
		return err
	}

	return nil
}

// ViewLeaveHistory implements leaves.ServiceLeavesInterface.
func (s *leavesService) ViewLeaveHistory(companyID uint, personalDataID uint, page, pageSize int, status string, startDate, endDate string) ([]leaves.LeavesDataEntity, error) {
	var leaveEntities []leaves.LeavesDataEntity
	var err error

	if status != "" {
		leaveEntities, err = s.leavesData.GetLeavesByStatus(personalDataID, status)
		if err != nil {
			log.Println("Error getting leaves by status:", err)
			return nil, err
		}
	} else if startDate != "" && endDate != "" {
		leaveEntities, err = s.leavesData.GetLeavesByDateRange(personalDataID, startDate, endDate)
		if err != nil {
			log.Println("Error getting leaves by date range:", err)
			return nil, err
		}
	} else {
		leaveEntities, err = s.leavesData.GetLeaveHistory(companyID, personalDataID, page, pageSize)
		if err != nil {
			log.Println("Error viewing leave history:", err)
			return nil, err
		}
	}

	return leaveEntities, nil
}

// GetLeavesByID implements leaves.ServiceLeavesInterface.
func (s *leavesService) GetLeavesByID(leaveID uint) (leaves *leaves.LeavesDataEntity, err error) {
	leaveEntity, err := s.leavesData.GetLeavesDetail(leaveID)
	if err != nil {
		log.Println("Error getting leave detail:", err)
		return nil, err
	}

	return leaveEntity, nil
}

// GetLeavesEmployees implements leaves.ServiceLeavesInterface.
func (s *leavesService) ViewLeaveHistoryEmployee(personalDataID uint, page, pageSize int, status string, startDate, endDate string) ([]leaves.LeavesDataEntity, error) {
	var leaveEntities []leaves.LeavesDataEntity
	var err error

	if status != "" {
		leaveEntities, err = s.leavesData.GetLeavesByStatus(personalDataID, status)
		if err != nil {
			log.Println("Error getting leaves by status:", err)
			return nil, err
		}
	} else if startDate != "" && endDate != "" {
		leaveEntities, err = s.leavesData.GetLeavesByDateRange(personalDataID, startDate, endDate)
		if err != nil {
			log.Println("Error getting leaves by date range:", err)
			return nil, err
		}
	} else {
		leaveEntities, err = s.leavesData.GetLeaveHistoryEmployee(personalDataID, page, pageSize)
		if err != nil {
			log.Println("Error viewing leave history:", err)
			return nil, err
		}
	}

	return leaveEntities, nil
}

func (s *leavesService) Dashboard(companyID uint) (*leaves.DashboardLeavesStats, error) {
	totalUsers, err := s.leavesData.CountTotalUsers(companyID)
	if err != nil {
		log.Printf("Error counting total users: %v", err)
		return nil, err
	}

	pendingLeaves, err := s.leavesData.CountPendingLeaves(companyID)
	if err != nil {
		log.Printf("Error counting pending leaves: %v", err)
		return nil, err
	}
	return &leaves.DashboardLeavesStats{
		TotalUsers:    totalUsers,
		LeavesPending: pendingLeaves,
	}, nil
}

var bulanMap = map[string]string{
	"Januari":   "January",
	"Februari":  "February",
	"Maret":     "March",
	"April":     "April",
	"Mei":       "May",
	"Juni":      "June",
	"Juli":      "July",
	"Agustus":   "August",
	"September": "September",
	"Oktober":   "October",
	"November":  "November",
	"Desember":  "December",
}

func ConvertIndonesiaMonthToEnglish(dateStr string) (string, error) {
	parts := strings.Split(dateStr, " ")
	if len(parts) != 3 {
		return "", errors.New("format tanggal tidak valid")
	}
	day := parts[0]
	monthIndo := parts[1]
	year := parts[2]

	monthEng, ok := bulanMap[monthIndo]
	if !ok {
		return "", fmt.Errorf("bulan %s tidak valid", monthIndo)
	}

	return fmt.Sprintf("%s %s %s", day, monthEng, year), nil
}

func CalculateLeaveDays(startDate, endDate string) (int, error) {
	layout := "02 January 2006"

	startDateEng, err := ConvertIndonesiaMonthToEnglish(startDate)
	if err != nil {
		return 0, err
	}
	endDateEng, err := ConvertIndonesiaMonthToEnglish(endDate)
	if err != nil {
		return 0, err
	}

	start, err := time.Parse(layout, startDateEng)
	if err != nil {
		return 0, err
	}

	end, err := time.Parse(layout, endDateEng)
	if err != nil {
		return 0, err
	}

	if end.Before(start) {
		return 0, errors.New("tanggal akhir tidak boleh sebelum tanggal mulai")
	}

	duration := end.Sub(start)
	return int(duration.Hours()/24) + 1, nil
}
