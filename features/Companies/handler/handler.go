package handler

import (
	"be-empower-hr/app/middlewares"
	companies "be-empower-hr/features/Companies"
	"be-empower-hr/utils/responses"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CompanyHandlers struct {
	srv companies.Service
};


func NewCompanyHandler(s companies.Service) companies.Handler {
	return &CompanyHandlers{
		srv: s,
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
		if err != nil {
			companyPicture = nil;
		}

		companySignature, err := c.FormFile("signature");
		if err != nil {
			companySignature = nil;
		}

		err = ch.srv.UpdateCompany(companyID, companyPicture, companySignature,  ToModelCompany(input));
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "Internal server error: "+err.Error(), nil))
		}

		return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "update company successfully", nil))
	}
}
