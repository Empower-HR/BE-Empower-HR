package service

import (
	leaves "be-empower-hr/features/Leaves"
	"be-empower-hr/utils"
	"errors"
	"log"
)

type leavesService struct {
	leavesData     leaves.DataLeavesInterface
	accountUtility utils.AccountUtilityInterface
}

func New(ud leaves.DataLeavesInterface, au utils.AccountUtilityInterface) leaves.ServiceLeavesInterface {
	return &leavesService{
		leavesData:     ud,
		accountUtility: au,
	}
}

// UpdateLeaveStatus implements leaves.ServiceLeavesInterface.
func (s *leavesService) UpdateLeaveStatus(leaveID uint, status string) error {
	if status != "approved" && status != "rejected" {
		return errors.New("invalid status")
	}

	err := s.leavesData.UpdateLeaveStatus(leaveID, status)
	if err != nil {
		log.Println("Error updating leave status:", err)
		return err
	}

	// Additional logic if leave is approved
	// if status == "approved" {
	// 	err = s.leavesData.DecrementTotalLeave(leaveID)
	// 	if err != nil {
	// 		log.Println("Error decrementing total leave:", err)
	// 		return err
	// 	}
	// }

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
func (s *leavesService) ViewLeaveHistory(personalDataID uint, page, pageSize int) ([]leaves.LeavesDataEntity, error) {
	panic("Not implemented")
}
