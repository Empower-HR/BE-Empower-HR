package service_test

import (
	"bytes"
	"errors"
	"testing"

	"be-empower-hr/features/Announcement"
	"be-empower-hr/features/Announcement/service"
	"be-empower-hr/mocks"

	"github.com/stretchr/testify/assert"
)

func TestAddAnno(t *testing.T) {
	qry := mocks.NewAnnoQuery(t)
	hashService := mocks.NewHashInterface(t)
	middlewareService := mocks.NewMiddlewaresInterface(t)
	accountUtility := mocks.NewAccountUtilityInterface(t)
	cloudinaryUtility := mocks.NewCloudinaryUtilityInterface(t)

	srv := service.New(qry, hashService, middlewareService, accountUtility, cloudinaryUtility)

	newAnno := Announcement.Announcement{
		CompanyID:   1,
		Title:       "Test Title",
		Description: "Test Description",
		Attachment:  "Test Attachment",
	}

	t.Run("Success Add Announcement", func(t *testing.T) {
		qry.On("Create", newAnno).Return(nil).Once()

		err := srv.AddAnno(newAnno)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})

	t.Run("Error Add Announcement", func(t *testing.T) {
		qry.On("Create", newAnno).Return(errors.New("internal server error")).Once()

		err := srv.AddAnno(newAnno)

		qry.AssertExpectations(t)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "terjadi kesalahan pada server saat membuat pesan")
	})
}

func TestGetAll(t *testing.T) {
	qry := mocks.NewAnnoQuery(t)
	hashService := mocks.NewHashInterface(t)
	middlewareService := mocks.NewMiddlewaresInterface(t)
	accountUtility := mocks.NewAccountUtilityInterface(t)
	cloudinaryUtility := mocks.NewCloudinaryUtilityInterface(t)

	srv := service.New(qry, hashService, middlewareService, accountUtility, cloudinaryUtility)

	expectedAnnos := []Announcement.Announcement{
		{CompanyID: 1, Title: "Title 1", Description: "Description 1", Attachment: "Attachment 1"},
		{CompanyID: 2, Title: "Title 2", Description: "Description 2", Attachment: "Attachment 2"},
	}

	t.Run("Success Get All Announcements", func(t *testing.T) {
		qry.On("GetAll").Return(expectedAnnos, nil).Once()

		annos, err := srv.GetAll()

		qry.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expectedAnnos, annos)
	})

	t.Run("Error Get All Announcements", func(t *testing.T) {
		qry.On("GetAll").Return(nil, errors.New("internal server error")).Once()

		annos, err := srv.GetAll()

		qry.AssertExpectations(t)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "error retrieving attendance records")
		assert.Nil(t, annos)
	})
}

func TestGetURLAtc(t *testing.T) {
	qry := mocks.NewAnnoQuery(t)
	hashService := mocks.NewHashInterface(t)
	middlewareService := mocks.NewMiddlewaresInterface(t)
	accountUtility := mocks.NewAccountUtilityInterface(t)
	cloudinaryUtility := mocks.NewCloudinaryUtilityInterface(t)

	srv := service.New(qry, hashService, middlewareService, accountUtility, cloudinaryUtility)

	file := bytes.NewReader([]byte("file content"))
	filename := "testfile.txt"
	expectedURL := "http://cloudinary.com/testfile.txt"

	t.Run("Success Get URL Attachment", func(t *testing.T) {
		cloudinaryUtility.On("UploadCloudinary", file, filename).Return(expectedURL, nil).Once()

		url, err := srv.GetURLAtc(file, filename)

		cloudinaryUtility.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expectedURL, url)
	})

	t.Run("Error Get URL Attachment", func(t *testing.T) {
		cloudinaryUtility.On("UploadCloudinary", file, filename).Return("", errors.New("upload error")).Once()

		url, err := srv.GetURLAtc(file, filename)

		cloudinaryUtility.AssertExpectations(t)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "upload error")
		assert.Empty(t, url)
	})
}
