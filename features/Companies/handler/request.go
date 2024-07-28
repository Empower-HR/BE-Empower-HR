package handler

import companies "be-empower-hr/features/Companies"

type CompanyInput struct {
	CompanyPicture string `json:"company_picture" form:"company_picture"`
	CompanyName    string `json:"company_name" form:"company_name"`
	Email          string `json:"email" form:"email"`
	PhoneNumber    string `json:"phone" form:"phone"`
	Npwp           int    `json:"npwp" form:"npwp"`
	CompanyAddress string `json:"company_address" form:"company_address"`
	Signature      string `json:"signature" form:"signature"`
}

func ToModelCompany(ci CompanyInput) companies.CompanyDataEntity{
	return companies.CompanyDataEntity{
		CompanyPicture  : ci.CompanyPicture,
		CompanyName    	: ci.CompanyName,
		Email          	: ci.Email,
		PhoneNumber     : ci.PhoneNumber,
		Npwp            : ci.Npwp,
		CompanyAddress  : ci.CompanyAddress,
		Signature       : ci.Signature,
	}
}