package catalogue_test

import (
	"fmt"
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
	s.Run("Error", func() {
		variation := &proto.Variation{
			Type: "Warna",
			Value: "Merah",
		}
		s.sqlmock.ExpectBegin()
		s.sqlmock.ExpectExec("INSERT INTO `variations`").
			WithArgs(
				"Warna", "Merah",
			).
			WillReturnError(fmt.Errorf("unexpected error"))
		s.sqlmock.ExpectRollback()
		err := s.subject.AddVariation(s.context, variation)
		s.EqualError(err, "insert node to table \"variations\": unexpected error")
	})
}

