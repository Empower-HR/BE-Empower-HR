package handler

import companies "be-empower-hr/features/Companies"

type CompanyResponse struct {
	ID             uint   `json:"id"`
	CompanyPicture string `json:"company_picture"`
	CompanyName    string `json:"company_name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone"`
	Address        string `json:"address"`
	Npwp           int    `json:"npwp"`
	CompanyAddress string `json:"company_address"`
	Signature      string `json:"signature"`
}

func ToResponseCompany(input companies.CompanyDataEntity) CompanyResponse{
	return CompanyResponse{
	ID       		: input.ID,    
	CompanyPicture  : input.CompanyPicture,
	CompanyName    	: input.CompanyName,
	Email          	: input.Email,
	PhoneNumber     : input.PhoneNumber,
	Address         : input.Address,
	Npwp            : input.Npwp,
	CompanyAddress  : input.CompanyAddress,
	Signature       : input.Signature,
	}
}