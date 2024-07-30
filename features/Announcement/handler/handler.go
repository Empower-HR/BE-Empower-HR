package handler

import (
	"be-empower-hr/app/middlewares"
	announcement "be-empower-hr/features/Announcement"
	"be-empower-hr/utils/responses"
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
)

type AnnHandler struct {
	srv announcement.AnnoServices
}

func New(anoserv announcement.AnnoServices) announcement.AnnoHandler {
	return &AnnHandler{
		srv: anoserv,
	}
}


// add Data Absen
func (ah *AnnHandler) AddAnnouncement(c echo.Context) error {
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "unauthorized", nil))
	}

	annRequest := AnnRequest{}
	if errBind := c.Bind(&annRequest); errBind != nil {
		log.Printf("Add Attendances: Error binding data: %v", errBind)
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error binding data: "+errBind.Error(), nil))
	}

	// Handle profile picture upload to Cloudinary
	var attachmentURL string
	attachment, err := c.FormFile("attachment")
	if err == nil {
		src, err := attachment.Open()
		if err != nil {
			log.Printf("attachment: Error opening file: %v", err)
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error opening file: "+err.Error(), nil))
		}
		defer src.Close()
		attachmentURL, err = ah.srv.GetURLAtc(src, attachment.Filename)
		if err != nil {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error binding data: "+err.Error(), nil))
		}
	}


	dataAnn := announcement.Announcement{
		CompanyID         : annRequest.CompanyID,
		Title		   	  : annRequest.Title,
		Description       : annRequest.Description,
		Attachment		  : attachmentURL,
	}
	// panggil fungsi addAtt pada service
	err = ah.srv.AddAnno(dataAnn)
	if err != nil {
		log.Printf("Add Attandance: Error add Attendance: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error add data: "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Add Attchment successfully", nil))
}

func (ah *AnnHandler) GetAnno(c echo.Context) error {
	personalID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if personalID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "unauthorized", nil))
	}
	var anno []announcement.Announcement
	var response []AnnoResponse
	
	anno, err := ah.srv.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "failed", "Failed to retrieve announcement data", nil))
	}

	for _, ann := range anno {
		response = append(response, ToResponseAnno(ann))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK,"success", "Successfully retrieved all announcement data", response))

}