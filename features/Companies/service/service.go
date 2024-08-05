package service

import (
	companies "be-empower-hr/features/Companies"
	"be-empower-hr/utils/cloudinary"
	"errors"
	"log"
	"mime/multipart"
)

type CompanyServices struct {
	qry companies.Query
	cld cloudinary.CloudinaryUtilityInterface
};

func NewCompanyServices(q companies.Query, c cloudinary.CloudinaryUtilityInterface) companies.Service {
	return &CompanyServices{
		qry: q,
		cld: c,
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

func (cs *CompanyServices) UpdateCompany(ID uint, companyPicture, companySignature *multipart.FileHeader, updateCompany companies.CompanyDataEntity) error {
	
	// upload company picture 
	if companyPicture != nil {
		src, err := companyPicture.Open();
		if err != nil {
			log.Print("Error opening company picture", err.Error())
			return errors.New("image error: " + err.Error())
		}
		defer src.Close();

		companyPictureURL, err := cs.cld.UploadCloudinary(src, companyPicture.Filename);
		if err != nil {
			log.Print("Error uploading company picture", err.Error())
			return errors.New("image error: " + err.Error())
		}
		updateCompany.CompanyPicture = companyPictureURL
	}

	// upload company picture 
	if companySignature != nil {
		src, err := companySignature.Open();
		if err != nil {
			log.Print("Error opening company signature", err.Error())
			return errors.New("image error: " + err.Error())
		}
		defer src.Close();

		companySignatureURL, err := cs.cld.UploadCloudinary(src, companyPicture.Filename);
		if err != nil {
			log.Print("Error uploading company signature", err.Error())
			return errors.New("image error: " + err.Error())
		}
		updateCompany.Signature = companySignatureURL
	}

	err := cs.qry.UpdateCompany(ID, updateCompany);
	if err != nil {
		log.Print("update company error", err.Error())
		return  errors.New("internal server error")
	};
		
	return nil;
}