package management_test

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/satriahrh/adam-smith/build/proto"
)

func (s *Suite) TestAddVariant() {
	images := []string{"https://savan.id/1.jpg"}
	jsonImages, _ := json.Marshal(images)
	stock := uint32(12)
	price := uint32(12000)
	productID, variationID := uint64(1), uint64(2)
	insertIntoArgsExpectation := []driver.Value{
		jsonImages, stock, price, variationID, productID,
	}
	sampleVariant := func() *proto.Variant {
		return &proto.Variant{
			Images: images,
			Stock:  stock,
			Price:  price,
		}
	}
	s.Run("NoError", func() {
		variant := sampleVariant()
		s.sqlmock.ExpectBegin()
		s.sqlmock.ExpectExec("INSERT INTO `variants`").
			WithArgs(insertIntoArgsExpectation...).
			WillReturnResult(sqlmock.NewResult(123, 1))
		s.sqlmock.ExpectCommit()
		err := s.subject.AddVariant(s.context, productID, variationID, variant)
		s.NoError(err)
		s.Equal(uint64(123), variant.Id)
	})
	s.Run("Error", func() {
		variant := sampleVariant()
		s.sqlmock.ExpectBegin()
		s.sqlmock.ExpectExec("INSERT INTO `variants`").
			WithArgs(insertIntoArgsExpectation...).
			WillReturnError(fmt.Errorf("unexpected error"))
		s.sqlmock.ExpectRollback()
		err := s.subject.AddVariant(s.context, productID, variationID, variant)
		s.EqualError(err, "insert node to table \"variants\": unexpected error")
	})
}
