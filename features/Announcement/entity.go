package Announcement

import (
	"io"

	"github.com/labstack/echo/v4"
)

type Announcement struct {
	ID              uint
 	CompanyID	 	uint
	Title	        string
	Description	    string
	Attachment	    string
}

type AnnoHandler interface {
	AddAnnouncement(c echo.Context) error
	GetAnno(c echo.Context) error
}

type AnnoServices interface {
	AddAnno(newAnno Announcement) error
	GetAll() ([]Announcement, error)
	GetURLAtc(file io.Reader, filename string) (string, error)
}

type AnnoQuery interface {
	Create(newAnn Announcement) error
    GetAll() ([]Announcement, error)

}
