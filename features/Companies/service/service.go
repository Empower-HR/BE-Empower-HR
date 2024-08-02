package service

import (
	companies "be-empower-hr/features/Companies"
	"errors"
	"log"
)

type CompanyServices struct {
	qry companies.Query
};

func NewCompanyServices(q companies.Query) companies.Service {
	return &CompanyServices{
		qry: q,
	}
};


func (cs *CompanyServices) GetCompany(ID uint) (companies.CompanyDataEntity, error){

	result, err := cs.qry.GetCompanyID(ID);

	if err != nil {
		log.Print("get company query error", err.Error())
		return companies.CompanyDataEntity{}, errors.New("internal server error")
	};
		
	return result, nil;
};

func (cs *CompanyServices) UpdateCompany(ID uint, updateCompany companies.CompanyDataEntity) error {
	
	err := cs.qry.UpdateCompany(ID, updateCompany);

	if err != nil {
		log.Print("update company error", err.Error())
		return  errors.New("internal server error")
	};
		
	return nil;
}