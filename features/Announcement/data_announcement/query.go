package dataannouncement

import (
	announcement "be-empower-hr/features/Announcement"

	"gorm.io/gorm"
)

type AnnoModel struct {
	db *gorm.DB
}

func NewModelAnnouncement(connection *gorm.DB) announcement.AnnoQuery {
	return &AnnoModel{
		db: connection,
	}
}

// Create Att
func (anmo *AnnoModel) Create(newAnn announcement.Announcement) error {
	cnv := AnnouncementInput(newAnn)
	return anmo.db.Create(&cnv).Error
}
func (anmo *AnnoModel) GetAll() ([]announcement.Announcement, error) {
	var announcement []announcement.Announcement
	err := anmo.db.Where("deleted_at IS NULL").Order("created_at DESC").Find(&announcement).Error
	if err != nil {
		return nil, err
	}
	return announcement, nil
}




