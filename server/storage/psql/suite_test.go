package psql

import (
	"advdiploma/server/storage/psql/migrations"
	"advdiploma/server/storage/psql/testtool"
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

const (
	dbName = "tst_00"
)

type TestSuite struct {
	suite.Suite

	container *testtool.PostgreSQLContainer
	storage   *Storage
	//fixtures  fixtures.Fixtures

	ctx context.Context
}

func (s *TestSuite) SetupSuite() {
	//ctx, ctxCancel := context.WithTimeout(context.Background(), 90*time.Second)
	//defer ctxCancel()

	ctx := context.Background()

	c, err := testtool.NewPostgreSQLContainer(ctx, testtool.WithPostgreSQLDatabaseName(dbName))
	s.Require().NoError(err)

	s.Require().NoError(migrations.RunMigrations(c.GetDSN(), dbName))

	st, err := NewStorage(c.GetDSN())
	s.Require().NoError(err)

	//fixtures, err := fixtures.LoadFixtures(ctx, st.db)
	//s.Require().NoError(err)

	s.ctx = context.TODO()
	s.container = c
	s.storage = st
	//	s.fixtures = fixtures
}

func (s *TestSuite) TearDownSuite() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()

	s.storage.Close()
	s.Require().NoError(s.container.Terminate(ctx))
}

func TestSuite_PostgreSQLStorage(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
