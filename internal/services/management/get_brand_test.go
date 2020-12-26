package management_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/satriahrh/adam-smith/build/proto"
)

func (s *Suite) TestGetBrand() {
	s.Run("ByID", func() {
		s.Run("Error", func() {
			s.sqlmock.ExpectQuery("SELECT (.+) FROM `brands`").
				WithArgs(123).
				WillReturnRows(sqlmock.NewRows([]string{"id", "code", "name"}))
			_, err := s.subject.GetBrand(s.context, 123, "")
			s.EqualError(err, "brand not found with given id/code")
		})
		s.Run("Succeed", func() {
			s.sqlmock.ExpectQuery("SELECT (.+) FROM `brands`").
				WithArgs("savan").
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "code", "name"}).
						AddRow(uint64(123), "savan", "Savan Baby Wear"),
				)
			protoBrand, err := s.subject.GetBrand(s.context, 0, "savan")
			s.NoError(err)
			s.EqualValues(&proto.Brand{
				Id:   123,
				Code: "savan",
				Name: "Savan Baby Wear",
			}, protoBrand)
		})
	})
	s.Run("ByCode", func() {
		s.Run("Error", func() {
			s.sqlmock.ExpectQuery("SELECT (.+) FROM `brands`").
				WithArgs("savan").
				WillReturnRows(sqlmock.NewRows([]string{"id", "code", "name"}))
			_, err := s.subject.GetBrand(s.context, 0, "savan")
			s.EqualError(err, "brand not found with given id/code")
		})
		s.Run("Succeed", func() {
			s.sqlmock.ExpectQuery("SELECT (.+) FROM `brands`").
				WithArgs("savan").
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "code", "name"}).
						AddRow(uint64(123), "savan", "Savan Baby Wear"),
				)
			protoBrand, err := s.subject.GetBrand(s.context, 0, "savan")
			s.NoError(err)
			s.EqualValues(&proto.Brand{
				Id:   123,
				Code: "savan",
				Name: "Savan Baby Wear",
			}, protoBrand)
		})
	})
}
