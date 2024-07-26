package handler

import companies "be-empower-hr/features/Companies"

type CompanyInput struct {
	CompanyPicture string `json:"company_picture"`
	CompanyName    string `json:"company_name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone"`
	Address        string `json:"address"`
	Npwp           int    `json:"npwp"`
	CompanyAddress string `json:"company_address"`
	Signature      string `json:"signature"`
}

func ToModelCompany(ci CompanyInput) companies.CompanyDataEntity{
	return companies.CompanyDataEntity{
		CompanyPicture  : ci.CompanyPicture,
		CompanyName    	: ci.CompanyName,
		Email          	: ci.Email,
		PhoneNumber     : ci.PhoneNumber,
		Address         : ci.Address,
		Npwp            : ci.Npwp,
		CompanyAddress  : ci.CompanyAddress,
		Signature       : ci.Signature,
	}
}