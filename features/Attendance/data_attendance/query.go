package dataattendance

import (
	"be-empower-hr/features/Attendance"
	"gorm.io/gorm"
)

type AttandanceModel struct {
	db *gorm.DB
}


func NewAttandancesModel(connection *gorm.DB) attendance.AQuery {
	return &AttandanceModel{
		db: connection,
	}
}

// Create Att
func (am *AttandanceModel) Create(newAtt attendance.Attandance) error {
	cnv := AttandanceInput(newAtt)
	return am.db.Create(&cnv).Error
}

func (am *AttandanceModel) IsDateExists(personalID uint, date string) (bool, error) {
	var count int64
	err := am.db.Model(&attendance.Attandance{}).
		Where("personal_data_id = ? AND date = ? AND deleted_at IS NULL", personalID, date).
		Count(&count).Error
	if err != nil {
		return false, err 
	}
	return count > 0, nil
}

// Update updates an existing attendance record
func (am *AttandanceModel) Update(id uint, updatedAtt attendance.Attandance) error {
	// Create a map with fields to update
	updateData := map[string]interface{}{
		"clock_out": updatedAtt.Clock_out,
		"status":    updatedAtt.Status,
		"long":      updatedAtt.Long,
		"lat":       updatedAtt.Lat,
		"notes":     updatedAtt.Notes,
	}
	err := am.db.Model(&attendance.Attandance{}).Where("id = ? AND deleted_at IS NULL", id).Updates(updateData).Error
	if err != nil {
		return err
	}

	return nil
}
// Get Att
func (am *AttandanceModel) GetAttByPersonalID(personalID uint, limit int, offset int) ([]attendance.Attandance, error) {
	var attandances []attendance.Attandance

	err := am.db.Where("personal_data_id = ?", personalID).Limit(limit).Offset(offset).Find(&attandances).Error
	if err != nil {
		return nil, err
	}

	return attandances, nil
}
func (am *AttandanceModel) GetAllAtt(limit int, offset int) ([]attendance.Attandance, error) {
	var attendances []attendance.Attandance
	err := am.db.Where("deleted_at IS NULL").Limit(limit).Offset(offset).Find(&attendances).Error
	if err != nil {
		return nil, err
	}
	return attendances, nil
}

// delete Att
func (am *AttandanceModel) DeleteAttbyId(attId uint) error {
	var attandances attendance.Attandance
	err := am.db.First(&attandances, attId).Error
	if err != nil {
		return err
	}
	return am.db.Delete(&attandances).Error
}


func (am *AttandanceModel) GetTotalAttendancesCount() (int64, error) {
    var count int64
    err := am.db.Model(&attendance.Attandance{}).
        Where("deleted_at IS NULL").
        Count(&count).Error
    if err != nil {
        return 0, err
    }
    return count, nil
}