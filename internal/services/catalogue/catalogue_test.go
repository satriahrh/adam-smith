package catalogue_test

import (
	"github.com/facebook/ent/dialect/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/satriahrh/adam-smith/build/ent"
	"github.com/satriahrh/adam-smith/internal/services/catalogue"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	subject catalogue.Catalogue
	sqlmock sqlmock.Sqlmock
}

func (s *Suite) SetupSuite() {
	db, mock, _ := sqlmock.New()

	s.subject = catalogue.New(
		ent.NewClient([]ent.Option{
			ent.Driver(
				sql.OpenDB("mysql", db),
			),
		}...),
	)
	s.sqlmock = mock
}

func TestServicesCatalogue(t *testing.T) {
	suite.Run(t, new(Suite))
}
