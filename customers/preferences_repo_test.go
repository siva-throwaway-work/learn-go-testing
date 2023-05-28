package customers

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"testing"
	"time"

	"github.com/sivaprasadreddy/learn-go-testing/testsupport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PreferenceRepoTestSuite struct {
	suite.Suite
	PgContainer *postgres.PostgresContainer
}

func (suite *PreferenceRepoTestSuite) SetupSuite() {
	fmt.Println("-----------SetupSuite()-----------")
	container, err := postgres.RunContainer(context.Background(),
		testcontainers.WithImage("postgres:15.2-alpine"),
		postgres.WithInitScripts(filepath.Join("../testdata", "initdb.sql")),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatal(err)
	}
	suite.PgContainer = container
}

func (suite *PreferenceRepoTestSuite) SetupTest() {
	fmt.Println("-----------SetupTest()-----------")
}

func (suite *PreferenceRepoTestSuite) TearDownTest() {
	fmt.Println("-----------TearDownTest()-----------")
}

func (suite *PreferenceRepoTestSuite) TearDownSuite() {
	fmt.Println("-----------TearDownSuite()-----------")
	defer func() {
		if err := suite.PgContainer.Terminate(context.Background()); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()
}

func (suite *PreferenceRepoTestSuite) TestGetPreferenceByCustomerId() {
	fmt.Println("-----------TestGetPreferenceByCustomerId()-----------")
	ctx := context.Background()
	t := suite.T()
	connStr, err := suite.PgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)

	conn := testsupport.GetDb(connStr)
	preferencesRepo := NewPreferenceRepo(conn)

	preference, err := preferencesRepo.GetByCustId(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, preference)
	assert.Equal(t, true, preference.Subscribe)
}

func (suite *PreferenceRepoTestSuite) TestCreateCustomerPreference() {
	fmt.Println("-----------PreferenceRepoTestSuite()-----------")
	ctx := context.Background()
	t := suite.T()

	connStr, err := suite.PgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)

	conn := testsupport.GetDb(connStr)
	preferencesRepo := NewPreferenceRepo(conn)

	preference, err := preferencesRepo.CreatePreference(ctx, Preference{
		Id:        9,
		CustId:    2,
		Subscribe: false,
	})
	if err != nil {
		return
	}
	assert.NoError(t, err)

	preference, err = preferencesRepo.GetByCustId(ctx, preference.CustId)
	assert.NoError(t, err)
	assert.NotNil(t, preference)
	assert.Equal(t, false, preference.Subscribe)
}

func TestPreferenceRepoTestSuite(t *testing.T) {
	suite.Run(t, new(PreferenceRepoTestSuite))
}
