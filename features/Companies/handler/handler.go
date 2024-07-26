package handler

import (
	companies "be-empower-hr/features/Companies"
	"be-empower-hr/utils/cloudinary"
	"be-empower-hr/utils/responses"
	"log"
	"net/http"
	"strconv"

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

		CompanyID, err := strconv.Atoi(c.Param("id"));
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error company_id data: "+err.Error(), nil))
		}

		data, err := ch.srv.GetCompany(uint(CompanyID));
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "Internal server error: "+err.Error(), nil))
		}

		return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "company retrieved successfully", ToResponseCompany(data)))
	}
};


func (ch *CompanyHandlers) UpdateCompany() echo.HandlerFunc {
	return func(c echo.Context) error {
		CompanyID, err := strconv.Atoi(c.Param("id"));
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error company_id data: "+err.Error(), nil))
		};

		var input CompanyInput;

		err = c.Bind(&input);
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error binding data: "+err.Error(), nil))
		};

		file, err := c.FormFile("company_picture");

		if err == nil {
			src , err := file.Open();
			if err != nil {
				log.Print("Error", err.Error())
				return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "image error: "+err.Error(), nil))
			}
			defer src.Close()

			urlImage, err := ch.cld.UploadCloudinary(src, file.Filename);
			if err != nil {
				log.Print("Error", err.Error())
				return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "image error: "+err.Error(), nil))
			}
			// set company picture jadi url dari response cld nya
			input.CompanyPicture = urlImage;
		};

		err = ch.srv.UpdateCompany(uint(CompanyID), ToModelCompany(input));
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "Internal server error: "+err.Error(), nil))
		}

		return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "update company successfully", nil))
	}
}
