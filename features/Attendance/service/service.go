package service

import (
	"be-empower-hr/app/middlewares"
	att "be-empower-hr/features/Attendance"
	"be-empower-hr/utils"
	"be-empower-hr/utils/encrypts"
	"be-empower-hr/utils/pdf"
	"errors"
)

type attendanceService struct {
	qry    			  att.AQuery
	hashService       encrypts.HashInterface
	middlewareservice middlewares.MiddlewaresInterface
	accountUtility    utils.AccountUtilityInterface
	pdfUtility 		  pdf.PdfUtilityInterface
}

func New(ad att.AQuery, hash encrypts.HashInterface, mi middlewares.MiddlewaresInterface, au utils.AccountUtilityInterface, pu pdf.PdfUtilityInterface) att.AServices {
	return &attendanceService{
		qry:    			ad,
		hashService:       hash,
		middlewareservice: mi,
		accountUtility:    au,
		pdfUtility: pu,
	}

}

func (as *attendanceService) AddAtt(newAtt att.Attandance) error {
	// Check if a record already exists for the given personalID and date
    exists, err := as.qry.IsDateExists(newAtt.PersonalDataID, newAtt.Date)
    if err != nil {
        return err
    }
    if exists {
        return errors.New("attendance record already exists for this date")
    }

	err = as.qry.Create(newAtt)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat Clock In")
	}
	return nil
}
func (as *attendanceService) UpdateAtt(id uint,updateAtt att.Attandance) error {
	err := as.qry.Update(id, updateAtt)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat Clock Out")
	}
	return nil
}

func (as *attendanceService) DeleteAttByID(attID uint) error {
	err := as.qry.DeleteAttbyId(attID)
	if err != nil {
		return errors.New("error deleting attendance record")
	}
	return nil
}

func (as *attendanceService) GetAttByPersonalID(personalID uint, limit int, offset int) ([]att.Attandance, error) {
	attendances, err := as.qry.GetAttByPersonalID(personalID, limit, offset)
	if err != nil {
		return nil, errors.New("error retrieving attendance records")
	}
	return attendances, nil
}

func (as *attendanceService) GetAllAtt(limit int, offset int) ([]att.Attandance, error) {

    attendance, err := as.qry.GetAllAtt(limit, offset)
	if err != nil {
		return nil, errors.New("error retrieving attendance records")
	}
	return attendance, nil
}
func (ah *attendanceService) CountAllAtt() (int64, error) {
	count, err := ah.qry.GetTotalAttendancesCount()
	if err != nil {
		return 0, errors.New("terjadi kesalahan pada server saat menghitung total product")
	}
	return count, nil
}

// download
func (ah *attendanceService) DownloadAllAtt() error {
	attendance, err := ah.qry.GetAllAttDownload()
	if err != nil {
		return errors.New("error retrieving attendance records")
	}

	err = ah.pdfUtility.DownloadPdf(attendance, "Attendance.pdf")
	if err != nil {
		return errors.New("error download pdf")
	}

	return nil
}
