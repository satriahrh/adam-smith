package catalogue_test

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/satriahrh/adam-smith/build/proto"
)

func (s *Suite) TestAddBrand() {
	s.Run("NoError", func() {
		ctx := context.TODO()
		protoBrand := &proto.Brand{Name: "Savan", Code: "savan"}
		s.sqlmock.ExpectBegin()
		s.sqlmock.ExpectExec("INSERT INTO `brands`").
			WithArgs("savan", "Savan").
			WillReturnResult(sqlmock.NewResult(123, 1))
		s.sqlmock.ExpectCommit()
		err := s.subject.AddBrand(ctx, protoBrand)
		s.NoError(err)
		s.Equal(uint64(123), protoBrand.Id)
	})
	s.Run("Error", func() {
		ctx := context.TODO()
		protoBrand := &proto.Brand{Name: "Savan", Code: "savan"}
		s.sqlmock.ExpectBegin()
		s.sqlmock.ExpectExec("INSERT INTO `brands`").
			WithArgs("savan", "Savan").
			WillReturnError(errors.New("unexpected error"))
		s.sqlmock.ExpectRollback()
		err := s.subject.AddBrand(ctx, protoBrand)
		s.EqualError(err, "insert node to table \"brands\": unexpected error")
	})
}
