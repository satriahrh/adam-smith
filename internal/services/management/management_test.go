package management_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/facebook/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/satriahrh/adam-smith/build/ent"
	"github.com/satriahrh/adam-smith/internal/services/management"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type Suite struct {
	suite.Suite
	subject management.Catalogue
	sqlmock sqlmock.Sqlmock
	context context.Context
}

func (s *Suite) SetupSuite() {
	db, mock, _ := sqlmock.New()

	logger, _ := zap.NewDevelopment([]zap.Option{}...)
	s.subject = management.New(
		ent.NewClient([]ent.Option{
			ent.Driver(
				sql.OpenDB("mysql", db),
			),
		}...),
		logger,
	)
	s.sqlmock = mock
	s.context = context.WithValue(context.TODO(), "ActionID", strconv.Itoa(int(time.Now().Unix())))
}

func TestServicesCatalogue(t *testing.T) {
	suite.Run(t, new(Suite))
}
