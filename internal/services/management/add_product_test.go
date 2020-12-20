package management_test

import (
	"encoding/json"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/satriahrh/adam-smith/build/proto"
	"github.com/satriahrh/adam-smith/ent/schema"
)

func (s *Suite) TestAddProduct() {
	jsonDescriptions, _ := json.Marshal([]schema.ProductDescription{
		{Heading: "Deskripsi", Body: "Deskripsi itu ya ini"},
	})
	jsonImages, _ := json.Marshal(schema.ProductImages{
		Thumbnail: "https://savan.id/thumbnail.jpg",
		Displays:  []string{"https://sava.id/display_1.jpg"},
	})
	jsonMarketplaces, _ := json.Marshal(schema.ProductMarketplaces{
		Tokopedia: "https://tkpd.com/1",
		Shopee:    "https://shopee.com/2",
		Bukalapak: "https://bklpk.com/3",
	})
	protoDescriptions := []*proto.ProductDescription{
		{Heading: "Deskripsi", Body: "Deskripsi itu ya ini"},
	}
	protoImages := &proto.ProductImages{
		Thumbnail: "https://savan.id/thumbnail.jpg",
		Displays:  []string{"https://sava.id/display_1.jpg"},
	}
	protoMarketplaces := &proto.ProductMarketplaces{
		Tokopedia: "https://tkpd.com/1",
		Shopee:    "https://shopee.com/2",
		Bukalapak: "https://bklpk.com/3",
	}
	s.Run("NoError", func() {
		brandId := uint64(321)
		product := &proto.Product{
			Sku:          "kdsk",
			Name:         "Sleepsuit",
			Descriptions: protoDescriptions,
			Images:       protoImages,
			Marketplaces: protoMarketplaces,
		}
		s.sqlmock.ExpectBegin()
		s.sqlmock.ExpectExec("INSERT INTO `products`").
			WithArgs(
				"kdsk",
				"Sleepsuit",
				jsonDescriptions,
				jsonImages,
				jsonMarketplaces,
				brandId,
			).
			WillReturnResult(sqlmock.NewResult(123, 1))
		s.sqlmock.ExpectCommit()
		err := s.subject.AddProduct(s.context, brandId, product)
		s.NoError(err)
		s.Equal(uint64(123), product.Id)
	})
	s.Run("Error", func() {
		brandId := uint64(321)
		product := &proto.Product{
			Sku:          "kdsk",
			Name:         "Sleepsuit",
			Descriptions: protoDescriptions,
			Images:       protoImages,
			Marketplaces: protoMarketplaces,
		}
		s.sqlmock.ExpectBegin()
		s.sqlmock.ExpectExec("INSERT INTO `products`").
			WithArgs(
				"kdsk",
				"Sleepsuit",
				jsonDescriptions,
				jsonImages,
				jsonMarketplaces,
				brandId,
			).
			WillReturnError(fmt.Errorf("unexpected error"))
		s.sqlmock.ExpectRollback()
		err := s.subject.AddProduct(s.context, brandId, product)
		s.EqualError(err, "insert node to table \"products\": unexpected error")
	})
}
