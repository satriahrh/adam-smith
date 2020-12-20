package management_test

import (
	"errors"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/satriahrh/adam-smith/build/proto"
)

func (s *Suite) TestAddVariation() {
	s.Run("NoError", func() {
		variation := &proto.Variation{
			Type: "Warna",
			Value: "Merah",
		}
		s.sqlmock.ExpectBegin()
		s.sqlmock.ExpectExec("INSERT INTO `variations`").
			WithArgs(
				"Warna", "Merah",
			).
			WillReturnResult(sqlmock.NewResult(123, 1))
		s.sqlmock.ExpectCommit()
		err := s.subject.AddVariation(s.context, variation)
		s.NoError(err)
		s.Equal(uint64(123), variation.Id)
	})
	s.Run("DuplicatedError", func() {
		variation := &proto.Variation{
			Type: "Warna",
			Value: "Merah",
		}
		s.sqlmock.ExpectBegin()
		s.sqlmock.ExpectExec("INSERT INTO `variations`").
			WithArgs(
				"Warna", "Merah",
			).
			WillReturnError(errors.New("Error 1062: Duplicated"))
		s.sqlmock.ExpectRollback()
		err := s.subject.AddVariation(s.context, variation)
		s.EqualError(err, "these value is already created")
	})
}

