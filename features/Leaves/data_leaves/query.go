package dataleaves

import (
	leaves "be-empower-hr/features/Leaves"
	"be-empower-hr/utils"

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

func (q *leavesQuery) UpdateLeaveStatus(leaveID uint, updatesleaves leaves.LeavesDataEntity) error {
	var leaveData LeavesData

	if err := q.db.First(&leaveData, leaveID).Error; err != nil {
		return err
	}

	return q.db.Save(&leaveData).Error
}

func (q *leavesQuery) GetLeaveHistory(personalDataID uint, page, pageSize int) ([]leaves.LeavesDataEntity, error) {
	var leaveEntities []leaves.LeavesDataEntity
	pagination := utils.NewPagination(page, pageSize)

	err := q.db.Table("leaves_data").
		Select("leaves_data.id AS leaves_id, leaves_data.personal_data_id, personal_data.name, employment_data.job_position, leaves_data.start_date, leaves_data.end_date, leaves_data.reason, leaves_data.status, leaves_data.total_leave").
		Joins("JOIN personal_data ON leaves_data.personal_data_id = personal_data.id").
		Joins("JOIN employment_data ON leaves_data.personal_data_id = employment_data.personal_data_id").
		Where("leaves_data.personal_data_id = ?", personalDataID).
		Where("personal_data.role = ?", "employee").
		Offset(pagination.Offset()).
		Limit(pagination.PageSize).
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

// GetLeavesDetail implements leaves.DataLeavesInterface.
func (q *leavesQuery) GetLeavesDetail(leaveID uint) (*leaves.LeavesDataEntity, error) {
	var leaveEntity leaves.LeavesDataEntity

	err := q.db.Table("leaves_data").
		Select("leaves_data.id as leaves_id, leaves_data.personal_data_id, leaves_data.start_date, leaves_data.end_date, leaves_data.reason, leaves_data.status, leaves_data.total_leave, personal_data.name, employment_data.job_position").
		Joins("JOIN personal_data ON leaves_data.personal_data_id = personal_data.id").
		Joins("JOIN employment_data ON leaves_data.personal_data_id = employment_data.personal_data_id").
		Where("leaves_data.id = ?", leaveID).
		Scan(&leaveEntity).Error

	if err != nil {
		return nil, err
	}

	return &leaveEntity, nil
}
func (q *leavesQuery) CountTotalUsers(leaveID uint) (int64, error) {
	var count int64
	if err := q.db.Model(&PersonalData{}).Where("company_id = ?", leaveID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (uq *leavesQuery) CountPendingLeaves(leavesID uint) (int64, error) {
	var count int64
	if err := uq.db.Model(&LeavesData{}).
		Where("personal_data_id IN (SELECT id FROM personal_data WHERE company_id = ?)", leavesID).
		Where("status = ?", "pending").
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// GetLeavesEmployees implements leaves.DataLeavesInterface.
func (q *leavesQuery) GetLeaveHistoryEmployee(personalDataID uint, page, pageSize int) ([]leaves.LeavesDataEntity, error) {
	var leaveEntities []leaves.LeavesDataEntity
	pagination := utils.NewPagination(page, pageSize)
	err := q.db.Table("leaves_data").
		Select("leaves_data.id AS leaves_id, leaves_data.personal_data_id, personal_data.name, employment_data.job_position, leaves_data.start_date, leaves_data.end_date, leaves_data.reason, leaves_data.status, leaves_data.total_leave").
		Joins("JOIN personal_data ON leaves_data.personal_data_id = personal_data.id").
		Joins("JOIN employment_data ON leaves_data.personal_data_id = employment_data.personal_data_id").
		Where("leaves_data.personal_data_id = ?", personalDataID).
		Offset(pagination.Offset()).
		Limit(pagination.PageSize).
		Scan(&leaveEntities).Error

	if err != nil {
		return nil, err
	}

	return leaveEntities, nil
}
