package service

import (
	"be-empower-hr/app/middlewares"
	ann "be-empower-hr/features/Announcement"
	"be-empower-hr/utils"
	"be-empower-hr/utils/encrypts"
	"errors"
)

type announcementService struct {
	qry               ann.AnnoQuery
	hashService       encrypts.HashInterface
	middlewareservice middlewares.MiddlewaresInterface
	accountUtility    utils.AccountUtilityInterface

}

func New(ad ann.AnnoQuery, hash encrypts.HashInterface, mi middlewares.MiddlewaresInterface, au utils.AccountUtilityInterface) ann.AnnoServices {
	return &announcementService{
		qry:    			ad,
		hashService:       hash,
		middlewareservice: mi,
		accountUtility:    au,
	}

}


func (as *announcementService) AddAnno(newAnno ann.Announcement) error {
	
	err := as.qry.Create(newAnno)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat membuat pesan")
	}
	return nil
}

func (as *announcementService) GetAll() ([]ann.Announcement, error) {

    announcement, err := as.qry.GetAll()
	if err != nil {
		return nil, errors.New("error retrieving attendance records")
	}
	return announcement, nil
}