package service

import (
	leaves "be-empower-hr/features/Leaves"
	"errors"
	"log"
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
func (s *leavesService) UpdateLeaveStatus(leaveID uint, updatesleaves leaves.LeavesDataEntity) error {
	if updatesleaves.Status != "approved" && updatesleaves.Status != "rejected" {
		return errors.New("invalid status")
	}

	err := s.leavesData.UpdateLeaveStatus(leaveID, updatesleaves)
	if err != nil {
		log.Println("Error updating leave status:", err)
		return err
	}

	return nil
}

// RequestLeave implements leaves.ServiceLeavesInterface.
func (s *leavesService) RequestLeave(leave leaves.LeavesDataEntity) error {
	err := s.leavesData.RequestLeave(leave)
	if err != nil {
		log.Println("Error requesting leave:", err)
		return err
	}

	return nil
}

// ViewLeaveHistory implements leaves.ServiceLeavesInterface.
func (s *leavesService) ViewLeaveHistory(personalDataID uint, page, pageSize int, status string, startDate, endDate string) ([]leaves.LeavesDataEntity, error) {
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
		leaveEntities, err = s.leavesData.GetLeaveHistory(personalDataID, page, pageSize)
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
		leaveEntities, err = s.leavesData.GetLeaveHistory(personalDataID, page, pageSize)
		if err != nil {
			log.Println("Error viewing leave history:", err)
			return nil, err
		}
	}

	return leaveEntities, nil
}
