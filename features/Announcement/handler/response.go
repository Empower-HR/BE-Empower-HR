package handler

import (
	ann "be-empower-hr/features/Announcement"
)
type AnnoResponse struct {
	ID       uint  		 `json:"id"`
	CompanyID uint		 `json:"company_id"`
	Title  string 		 `json:"title"`
	Description string   `json:"description"`
	Attachment   string  `json:"attchment"`
}

func ToResponseAnno(announcement ann.Announcement) AnnoResponse {
	return AnnoResponse{
		ID:       announcement.ID,
		CompanyID: announcement.CompanyID,
		Title: announcement.Title,
		Description: announcement.Description,
		Attachment: announcement.Attachment,
	}
}