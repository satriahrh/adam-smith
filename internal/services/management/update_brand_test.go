package management_test

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/satriahrh/adam-smith/build/proto"
)

func (s *Suite) TestUpdateBrand() {
	s.Run("NoError", func() {
		brand := &proto.Brand{Id: 123, Name: "Savan", Code: "savan"}
		s.sqlmock.ExpectBegin()
		s.sqlmock.ExpectExec("UPDATE `brands` SET `name`").
			WithArgs("Savan", int64(123)).
			WillReturnResult(sqlmock.NewResult(123, 1))
		s.sqlmock.ExpectQuery("SELECT (.+) FROM `brands`").
			WithArgs(int64(123)).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "code", "name"}).
					AddRow(int64(123), "savan", "Savan"),
				)
		s.sqlmock.ExpectCommit()
		err := s.subject.UpdateBrand(s.context, brand)
		s.NoError(err)
		s.Equal(uint64(123), brand.Id)
	})
	s.Run("DuplicatedError", func() {
		brand := &proto.Brand{Id: 123, Name: "Savan", Code: "savan"}
		s.sqlmock.ExpectBegin()
		s.sqlmock.ExpectExec("UPDATE `brands` SET `name`").
			WithArgs("Savan", int64(123)).
			WillReturnError(errors.New("ent: brand not found"))
		s.sqlmock.ExpectRollback()
		err := s.subject.UpdateBrand(s.context, brand)
		s.EqualError(err, "brand with given id is not found")
	})
}
