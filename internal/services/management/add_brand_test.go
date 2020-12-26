package management_test

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/satriahrh/adam-smith/build/proto"
)

func (s *Suite) TestAddBrand() {
	s.Run("NoError", func() {
		brand := &proto.Brand{Name: "Savan", Code: "savan"}
		s.sqlmock.ExpectBegin()
		s.sqlmock.ExpectExec("INSERT INTO `brands`").
			WithArgs("savan", "Savan").
			WillReturnResult(sqlmock.NewResult(123, 1))
		s.sqlmock.ExpectCommit()
		err := s.subject.AddBrand(s.context, brand)
		s.NoError(err)
		s.Equal(uint64(123), brand.Id)
	})
	s.Run("DuplicatedError", func() {
		brand := &proto.Brand{Name: "Savan", Code: "savan"}
		s.sqlmock.ExpectBegin()
		s.sqlmock.ExpectExec("INSERT INTO `brands`").
			WithArgs("savan", "Savan").
			WillReturnError(errors.New("Error 1062: Duplicated"))
		s.sqlmock.ExpectRollback()
		err := s.subject.AddBrand(s.context, brand)
		s.EqualError(err, "brand with given code is already existed")
	})
}
