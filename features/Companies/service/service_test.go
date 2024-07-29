package service_test

import (
	companies "be-empower-hr/features/Companies"
	"be-empower-hr/features/Companies/service"
	"be-empower-hr/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCompany(t *testing.T) {
	qry := mocks.NewQuery(t)
	srv := service.NewCompanyServices(qry)

	t.Run("Success Get Company", func(t *testing.T) {
		expectedResult := companies.CompanyDataEntity{
			CompanyPicture: "img.png", 
			CompanyName: "Pt. Konoho", 
			Email: "konoha@indo.com", 
			PhoneNumber: "08345677543",
			Npwp: 9777868,
			CompanyAddress: "jl. indo konohan 1",
			Signature: "ttd.png",
		}

		qry.On("GetCompany").Return(expectedResult, nil).Once()

		result, err := srv.GetCompany()

		qry.AssertExpectations(t)

		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Error Get Company", func(t *testing.T) {
		qry.On("GetCompany").Return(companies.CompanyDataEntity{}, errors.New("internal server error")).Once()

		result, err := srv.GetCompany()

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "internal server error")
		assert.Equal(t, companies.CompanyDataEntity{}, result)
	})
}

func TestUpdateCompany(t *testing.T) {
	qry := mocks.NewQuery(t)
	srv := service.NewCompanyServices(qry)

	ID := uint(1)
	updateData := companies.CompanyDataEntity{
		CompanyPicture: "img.png", 
		CompanyName: "Pt. Konoho", 
		Email: "konoha@indo.com", 
		PhoneNumber: "08345677543",
		Npwp: 9777868,
		CompanyAddress: "jl. indo konohan 1",
		Signature: "ttd.png",
	}

	t.Run("Success Update Company", func(t *testing.T) {
		qry.On("UpdateCompany", ID, updateData).Return(nil).Once()

		err := srv.UpdateCompany(ID, updateData)

		qry.AssertExpectations(t)

		assert.Nil(t, err)
	})

	t.Run("Error Update Company", func(t *testing.T) {
		qry.On("UpdateCompany", ID, updateData).Return(errors.New("internal server error")).Once()

		err := srv.UpdateCompany(ID, updateData)

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "internal server error")
	})
}
