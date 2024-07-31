package service

import (
	"be-empower-hr/app/middlewares"
	att "be-empower-hr/features/Attendance"
	"be-empower-hr/utils"
	"be-empower-hr/utils/encrypts"
	"be-empower-hr/utils/maps"
	"be-empower-hr/utils/pdf"
	"errors"
	"fmt"
	"strconv"
)

type attendanceService struct {
	qry               att.AQuery
	hashService       encrypts.HashInterface
	middlewareservice middlewares.MiddlewaresInterface
	accountUtility    utils.AccountUtilityInterface
	pdfUtility        pdf.PdfUtilityInterface
	mapsUtility		  maps.MapsUtilityInterface
}

func New(ad att.AQuery, hash encrypts.HashInterface, mi middlewares.MiddlewaresInterface, au utils.AccountUtilityInterface, pu pdf.PdfUtilityInterface, mu maps.MapsUtilityInterface) att.AServices {
	return &attendanceService{
		qry: ad,
		// qryUser:           as,
		hashService:       hash,
		middlewareservice: mi,
		accountUtility:    au,
		pdfUtility:        pu,
		mapsUtility: mu,
	}

}

// AddAtt menambahkan catatan absensi baru
func (as *attendanceService) AddAtt(newAtt att.Attandance) error {
	// Periksa apakah catatan sudah ada untuk personalID dan tanggal yang diberikan
	exists, err := as.qry.IsDateExists(newAtt.PersonalDataID, newAtt.Date)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("attendance record already exists for this date")
	}

	// Ambil data perusahaan berdasarkan personalDataID
	company, err := as.qry.GetCompany(newAtt.PersonalDataID)
	if err != nil {
		return err
	}

	// Ambil alamat perusahaan dan lakukan geocoding untuk mendapatkan latitude dan longitude
	companyLat, companyLng, err := as.mapsUtility.GeoCode(company[0].CompanyAddress)
	if err != nil {
		return err
	}

	// Parsing latitude dan longitude dari newAtt
	attLat, err := strconv.ParseFloat(newAtt.Lat, 64)
	if err != nil {
		return errors.New("invalid latitude format")
	}

	attLng, err := strconv.ParseFloat(newAtt.Long, 64)
	if err != nil {
		return errors.New("invalid longitude format")
	}

	// Hitung jarak antara lokasi absensi dan lokasi perusahaan
	minDistance := 100.0 // Jarak minimum dalam meter
	distance := as.mapsUtility.Haversine(attLat, attLng, companyLat, companyLng)
	if distance > minDistance {
		return errors.New("absensi ditolak karena lokasi anda terpantau jauh dari kantor")
		 // Anda bisa mengembalikan nil karena absensi ditolak, bukan error
	}
	// Jika jarak dalam batas yang diperbolehkan, simpan catatan absensi
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
