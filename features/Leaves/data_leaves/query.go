package dataleaves

import (
	leaves "be-empower-hr/features/Leaves"

	"gorm.io/gorm"
)

type leavesQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) leaves.DataLeavesInterface {
	return &leavesQuery{
		db: db,
	}
}

func (q *leavesQuery) RequestLeave(leave leaves.LeavesDataEntity) error {
	leaveData := LeavesData{
		StartDate:      leave.StartDate,
		EndDate:        leave.EndDate,
		Reason:         leave.Reason,
		Status:         "pending",
		TotalLeave:     leave.TotalLeave,
		PersonalDataID: leave.PersonalDataID,
	}
	return q.db.Create(&leaveData).Error
}

func (q *leavesQuery) UpdateLeaveStatus(leaveID uint, status string) error {
	var leaveData LeavesData

	// Find the leave data by ID
	if err := q.db.First(&leaveData, leaveID).Error; err != nil {
		return err
	}

	// Update the status
	leaveData.Status = status

	if status == "approved" {
		leaveData.TotalLeave--
	}

	// Save the changes
	return q.db.Save(&leaveData).Error
}

func (q *leavesQuery) GetLeaveHistory(personalDataID uint, page, pageSize int) ([]leaves.LeavesDataEntity, error) {
	var leaveEntities []leaves.LeavesDataEntity

	// Query langsung ke database dan scan ke dalam slice leaveEntities
	err := q.db.Table("leaves_data").
		Select("leaves_data.personal_data_id, personal_data.name, employment_data.job_position, leaves_data.start_date, leaves_data.end_date, leaves_data.reason, leaves_data.status").
		Joins("JOIN personal_data ON leaves_data.personal_data_id = personal_data.id").
		Joins("JOIN employment_data ON leaves_data.personal_data_id = employment_data.personal_data_id").
		Where("leaves_data.personal_data_id = ?", personalDataID).
		Scan(&leaveEntities).Error

	if err != nil {
		return nil, err
	}

	return leaveEntities, nil
}

func (q *leavesQuery) GetLeavesByStatus(personalDataID uint, status string) ([]leaves.LeavesDataEntity, error) {
	var leavesData []LeavesData
	err := q.db.Where("personal_data_id = ? AND status = ?", personalDataID, status).Find(&leavesData).Error
	if err != nil {
		return nil, err
	}

	var result []leaves.LeavesDataEntity
	for _, leave := range leavesData {
		result = append(result, leaves.LeavesDataEntity{
			LeavesID:       leave.ID,
			StartDate:      leave.StartDate,
			EndDate:        leave.EndDate,
			Reason:         leave.Reason,
			Status:         leave.Status,
			TotalLeave:     leave.TotalLeave,
			PersonalDataID: leave.PersonalDataID,
		})
	}
	return result, nil
}

// GetLeavesByDateRange
func (q *leavesQuery) GetLeavesByDateRange(personalDataID uint, startDate, endDate string) ([]leaves.LeavesDataEntity, error) {
	var leaveEntities []leaves.LeavesDataEntity

	err := q.db.Table("leaves_data").
		Select("leaves_data.id as leaves_id, leaves_data.personal_data_id, personal_data.name, employment_data.job_position, leaves_data.start_date, leaves_data.end_date, leaves_data.reason, leaves_data.status").
		Joins("JOIN personal_data ON leaves_data.personal_data_id = personal_data.id").
		Joins("JOIN employment_data ON leaves_data.personal_data_id = employment_data.personal_data_id").
		Where("leaves_data.personal_data_id = ? AND leaves_data.start_date >= ? AND leaves_data.end_date <= ?", personalDataID, startDate, endDate).
		Scan(&leaveEntities).Error

	if err != nil {
		return nil, err
	}

	return leaveEntities, nil
}
