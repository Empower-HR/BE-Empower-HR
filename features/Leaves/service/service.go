package service

import (
	leaves "be-empower-hr/features/Leaves"
	"errors"
	"log"
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

	// Update the leave status
	if err := s.leavesData.UpdateLeaveStatus(leaveID, updatesleaves); err != nil {
		log.Println("Error updating leave status:", err)
		return err
	}

	// If the leave request is approved, adjust the total leave balance
	if updatesleaves.Status == "approved" {
		// Fetch leave details
		leaveDetail, err := s.leavesData.GetLeavesDetail(leaveID)
		if err != nil {
			return err
		}
		leaveDays, err := CalculateLeaveDays(leaveDetail.StartDate, leaveDetail.EndDate)
		if err != nil {
			return err
		}

		var personalData leaves.PersonalDataEntity
		if err := s.leavesData.GetPersonalDataByID(leaveDetail.PersonalDataID, &personalData); err != nil {
			return err
		}

		if personalData.TotalLeave < leaveDays {
			return errors.New("insufficient leave balance")
		}

		// Update the total leave balance
		personalData.TotalLeave -= leaveDays
		if err := s.leavesData.UpdatePersonalData(personalData); err != nil {
			return err
		}
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

func CalculateLeaveDays(startDate, endDate string) (int, error) {
	layout := "02 January 2006"
	start, err := time.Parse(layout, startDate)
	if err != nil {
		return 0, err
	}
	end, err := time.Parse(layout, endDate)
	if err != nil {
		return 0, err
	}

	if end.Before(start) {
		return 0, errors.New("end date cannot be before start date")
	}

	duration := end.Sub(start)
	return int(duration.Hours() / 24), nil
}
