package postgres

import (
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	libsql "github.com/webcore-go/lib-sql"
	"github.com/webcore-go/webcore/infra/config"
	"github.com/webcore-go/webcore/port"
)

type PostgresLoader struct {
	name string
}

func (a *PostgresLoader) SetName(name string) {
	a.name = name
}

func (a *PostgresLoader) Name() string {
	return a.name
}

func (l *PostgresLoader) Init(args ...any) (port.Library, error) {
	config := args[1].(config.DatabaseConfig)
	dsn := libsql.BuildDSN(config)

	db := &libsql.SQLDatabase{}

	driver := pgdriver.NewConnector(pgdriver.WithDSN(dsn))
	dialect := pgdialect.New()

	// Set up Bun SQL database wrapper
	db.SetBunDB(driver, dialect)

	err := db.Install(args...)
	if err != nil {
		return nil, err
	}

	db.Connect()

	return db, nil
}
