package dataannouncement

import (
	announcement "be-empower-hr/features/Announcement"

	"gorm.io/gorm"
)

type Announcement struct {
	gorm.Model
	CompanyID	   uint					  `gorm:"foreignKey:companyID"`
	Title	       string   		       `json:"title"`
	Description    string                  `json:"description"`
	Attachment     string                  `json:"attachment"`
}


func AnnouncementInput(input announcement.Announcement) Announcement{
	return Announcement{
	 	CompanyID: input.CompanyID,
		Title: input.Title,
		Description: input.Description,
		Attachment: input.Attachment,
	}
}