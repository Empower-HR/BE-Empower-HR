package handler

import (
	"be-empower-hr/app/middlewares"
	companies "be-empower-hr/features/Companies"
	"be-empower-hr/utils/cloudinary"
	"be-empower-hr/utils/responses"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CompanyHandlers struct {
	srv companies.Service
	cld cloudinary.CloudinaryUtilityInterface
};


func NewCompanyHandler(s companies.Service, c cloudinary.CloudinaryUtilityInterface) companies.Handler {
	return &CompanyHandlers{
		srv: s,
		cld: c,
	}
};

func (ch *CompanyHandlers) GetCompany() echo.HandlerFunc {
	return func(c echo.Context) error {
		companyID, _ := middlewares.NewMiddlewares().ExtractCompanyID(c)
		if companyID == 0 {
			return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}
		data, err := ch.srv.GetCompany(companyID);
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "Internal server error: "+err.Error(), nil))
		}

		return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "company retrieved successfully", ToResponseCompany(data)))
	}
};


func (ch *CompanyHandlers) UpdateCompany() echo.HandlerFunc {
	return func(c echo.Context) error {
		companyID, err := middlewares.NewMiddlewares().ExtractCompanyID(c)
		if companyID == 0 {
			return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		input := CompanyInput{};
		err = c.Bind(&input);
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error binding data: "+err.Error(), nil))
		};
		
		// handle company picture
		companyPicture, err := c.FormFile("company_picture");
		if err == nil {
			src , err := companyPicture.Open();
			if err != nil {
				log.Print("Error", err.Error())
				return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "image error: "+err.Error(), nil))
			}
			defer src.Close()

			companyPictureURL, err := ch.cld.UploadCloudinary(src, companyPicture.Filename);
			if err != nil {
				log.Print("Error", err.Error())
				return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "image error: "+err.Error(), nil))
			}
			// set company picture jadi url dari response cld nya
			input.CompanyPicture = companyPictureURL;
		};

		companySignature, err := c.FormFile("signature");
		if err == nil {
			src , err := companySignature.Open();
			if err != nil {
				log.Print("Error", err.Error())
				return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "image error: "+err.Error(), nil))
			}
			defer src.Close()

			companySignatureURL, err := ch.cld.UploadCloudinary(src, companySignature.Filename);
			if err != nil {
				log.Print("Error", err.Error())
				return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "image error: "+err.Error(), nil))
			}
			// set company picture jadi url dari response cld nya
			input.Signature = companySignatureURL;
		};

		err = ch.srv.UpdateCompany(companyID, ToModelCompany(input));
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "Internal server error: "+err.Error(), nil))
		}

		return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "update company successfully", nil))
	}
}
