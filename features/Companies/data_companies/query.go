package datacompanies

import (
	companies "be-empower-hr/features/Companies"

	"gorm.io/gorm"
)

type CompanyModels struct {
	db *gorm.DB
}

func NewCompanyModels(connect *gorm.DB) companies.Query {
	return &CompanyModels{
		db: connect,
	}
}

func (cm *CompanyModels) GetCompany() (companies.CompanyDataEntity, error) {
	var result CompanyData

	err := cm.db.Find(&result).Error

	if err != nil {
		return companies.CompanyDataEntity{}, err
	}

	return result.ToCompanyEntity(), nil
}

func (cm *CompanyModels) UpdateCompany(ID uint, updateCompany companies.CompanyDataEntity) error {
	cnvCompany := ToCompanyQuery(updateCompany)

	qry := cm.db.Model(CompanyData{}).Where("id = ?", ID).Updates(&cnvCompany)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (cm *CompanyModels) GetCompanyID(ID uint) (companies.CompanyDataEntity, error) {
	var result CompanyData

	err := cm.db.Model(CompanyData{}).Where("id = ?", ID).First(&result).Error

	if err != nil {
		return companies.CompanyDataEntity{}, err
	}

	return result.ToCompanyEntity(), nil
}
