package orders

//A common test file with TestMain(..) which executes once per entire package
// TestMain(..) will be called only once per package,
// All the containers will be started and be used by any test within this package.

// Issue:
// Having TestMain(..) behaviour per-package instead of per-test file results in not having more control.

/*
var dbConnString string
var redisConnString string

func TestMain(m *testing.M) {
	ctx := context.Background()
	pgContainer, err := SetupPostgres(ctx)
	if err != nil {
		log.Fatalf("failed to setup Postgres container. Error: %v", err)
	}
	dbConnString = pgContainer.ConnectionString
	defer pgContainer.CloseFn()

	redisContainer, err := SetupRedis(ctx)
	if err != nil {
		log.Fatalf("failed to setup Redis container. Error: %v", err)
	}
	redisConnString = redisContainer.ConnectionString
	defer redisContainer.CloseFn()

	os.Exit(m.Run())
}

type PostgresContainer struct {
	Container        testcontainers.Container
	CloseFn          func()
	ConnectionString string
}

type RedisContainer struct {
	Container        testcontainers.Container
	CloseFn          func()
	ConnectionString string
}

func SetupPostgres(ctx context.Context) (*PostgresContainer, error) {
	container, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.2-alpine"),
		postgres.WithInitScripts(filepath.Join("../testdata", "initdb.sql")),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, err
	}
	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &PostgresContainer{
		Container: container,
		CloseFn: func() {
			if err := container.Terminate(ctx); err != nil {
				log.Fatalf("error terminating postgres container: %s", err)
			}
		},
		ConnectionString: connStr,
	}, nil
}

func SetupRedis(ctx context.Context) (*RedisContainer, error) {
	redisC, err := redis.RunContainer(ctx,
		testcontainers.WithImage("redis:latest"),
	)
	if err != nil {
		return nil, err
	}
	connStr, err := redisC.ConnectionString(ctx)
	if err != nil {
		return nil, err
	}
	return &RedisContainer{
		Container: redisC,
		CloseFn: func() {
			if err := redisC.Terminate(ctx); err != nil {
				log.Fatalf("error terminating postgres container: %s", err)
			}
		},
		ConnectionString: connStr,
	}, nil
}
*/
