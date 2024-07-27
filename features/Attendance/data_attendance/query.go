package dataattendance

import (
	attendance "be-empower-hr/features/Attendance"

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
func (am *AttandanceModel) GetAttByPersonalID(personalID uint, limit int, offset int) ([]attendance.AttendanceDetail, error) {
	var results []attendance.AttendanceDetail

	query := `
    SELECT 
        pd.name,
		at.personal_data_id,
        sc.schedule_in, 
        sc.schedule_out, 
        at.clock_in, 
        at.clock_out,
		at.date,
		at.id
    FROM
        personal_data AS pd
    JOIN 
        schedule_data AS sc ON sc.company_id = pd.company_id
    LEFT JOIN 
        attandances AS at ON at.personal_data_id = pd.id
    WHERE 
       at.personal_data_id = ? AND at.deleted_at IS NULL 
    LIMIT ? OFFSET ?`

	err := am.db.Raw(query, personalID, limit, offset).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil

}

func (am *AttandanceModel) GetAllAttbyDate(date string, limit int, offset int) ([]attendance.AttendanceDetail, error) {
	var results []attendance.AttendanceDetail

	query := `
    SELECT 
        pd.name,
		at.personal_data_id, 
        sc.schedule_in, 
        sc.schedule_out, 
        at.clock_in, 
        at.clock_out,
		at.date,
		at.id
    FROM
        personal_data AS pd
    JOIN 
        schedule_data AS sc ON sc.company_id = pd.company_id
    LEFT JOIN 
        attandances AS at ON at.personal_data_id = pd.id AND at.date = ?
    WHERE 
        at.deleted_at IS NULL
    LIMIT ? OFFSET ?`

	err := am.db.Raw(query, date, limit, offset).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
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
func (am *AttandanceModel) GetTotalAttendancesCountbyDate(date string) (int64, error) {
	var count int64
	err := am.db.Model(&attendance.Attandance{}).
		Where("deleted_at IS NULL AND date = ?", date).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (am *AttandanceModel) GetAllAttDownload() ([]attendance.Attandance, error) {
	var attendances []attendance.Attandance
	err := am.db.Where("deleted_at IS NULL").Find(&attendances).Error
	if err != nil {
		return nil, err
	}
	return attendances, nil
}

// New Get Test
func (am *AttandanceModel) GetAttendanceDetails(limit int, offset int) ([]attendance.AttendanceDetail, error) {
	var results []attendance.AttendanceDetail

	query := `
    SELECT 
        pd.name, 
		at.personal_data_id,
        sc.schedule_in, 
        sc.schedule_out, 
        at.clock_in, 
        at.clock_out,
		at.date,
		at.id
    FROM
        personal_data AS pd
    JOIN 
        schedule_data AS sc ON sc.company_id = pd.company_id
    LEFT JOIN 
        attandances AS at ON at.personal_data_id = pd.id
    WHERE 
        at.deleted_at IS NULL
    LIMIT ? OFFSET ?`

	err := am.db.Raw(query, limit, offset).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}
