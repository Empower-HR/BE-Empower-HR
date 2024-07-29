package service

import (
	"be-empower-hr/app/middlewares"
	att "be-empower-hr/features/Attendance"
	"be-empower-hr/utils"
	"be-empower-hr/utils/encrypts"
	"be-empower-hr/utils/pdf"
	"errors"
	"fmt"
)

type attendanceService struct {
	qry               att.AQuery
	hashService       encrypts.HashInterface
	middlewareservice middlewares.MiddlewaresInterface
	accountUtility    utils.AccountUtilityInterface
	pdfUtility        pdf.PdfUtilityInterface
}

func New(ad att.AQuery, hash encrypts.HashInterface, mi middlewares.MiddlewaresInterface, au utils.AccountUtilityInterface, pu pdf.PdfUtilityInterface) att.AServices {
	return &attendanceService{
		qry: ad,
		// qryUser:           as,
		hashService:       hash,
		middlewareservice: mi,
		accountUtility:    au,
		pdfUtility:        pu,
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
func (as *attendanceService) UpdateAtt(id uint, updateAtt att.Attandance) error {
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

func (as *attendanceService) GetAttByPersonalID(personalID uint,searchBox string, limit int, offset int) ([]att.AttendanceDetail, error) {
	attendances, err := as.qry.GetAttByPersonalID(personalID, searchBox, limit, offset)
	if err != nil {
		return nil, errors.New("error retrieving attendance records")
	}
	return attendances, nil
}

func (as *attendanceService) GetAllAtt(search string, limit int, offset int) ([]att.AttendanceDetail, error) {

    // attendance, err := as.qry.GetAllAtt(limit, offset)
    attendance, err := as.qry.GetAttendanceDetails(search,limit, offset)
	if err != nil {
		return nil, errors.New("error retrieving attendance records")
	}
	return attendance, nil
}
func (as *attendanceService) GetAttByIdAtt(idAtt uint) ([]att.AttendanceDetail, error) {

    // attendance, err := as.qry.GetAllAtt(limit, offset)
    attendance, err := as.qry.GetAttByIdAtt(idAtt)
	if err != nil {
		return nil, errors.New("error retrieving attendance records")
	}
	return attendance, nil
}
func (as *attendanceService) GetAllAttbyDate(date string, limit int, offset int) ([]att.AttendanceDetail, error) {
	if date == "" {
		return nil, fmt.Errorf("silahkan isi tanggal dengan benar")
	}
	attendance, err := as.qry.GetAllAttbyDate(date, limit, offset)
	if err != nil {
		return nil, err
	}
	return attendance, nil
}

func (as *attendanceService) GetAllAttbyStatus(status string, limit int, offset int) ([]att.AttendanceDetail, error){
	if status == "" {
		return nil, fmt.Errorf("silahkan isi tanggal dengan benar")
	}
	attendance, err := as.qry.GetAllAttbyStatus(status, limit, offset)
	if err != nil {
		return nil, err
	}
	fmt.Println("Data service:",attendance)
	return attendance, nil
}

func (as *attendanceService) GetAttByPersonalIDandStatus(id uint, status string, limit int, offset int) ([]att.AttendanceDetail, error){
	if status == "" {
		return nil, fmt.Errorf("silahkan isi tanggal dengan benar")
	}
	attendance, err := as.qry.GetAllAttbyIdPersonAndStatus(id, status, limit, offset)
	if err != nil {
		return nil, err
	}
	return attendance, nil
}

func (as *attendanceService) CountAllAtt() (int64, error) {
	count, err := as.qry.GetTotalAttendancesCount()
	if err != nil {
		return 0, errors.New("terjadi kesalahan pada server saat menghitung total product")
	}
	return count, nil
}
func (as *attendanceService) CountAllAttbyDate(date string) (int64, error) {
	count, err := as.qry.GetTotalAttendancesCountbyDate(date)
	if err != nil {
		return 0, errors.New("terjadi kesalahan pada server saat menghitung total product")
	}
	return count, nil
}

func (as *attendanceService) CountAllAttbyStatus(status string) (int64, error) {
	count, err := as.qry.GetTotalAttendancesCountByStatus(status)
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
