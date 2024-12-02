package pg_conn

import (
	"basicProjectLayout/config"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewPsqlDB(cfg *config.Config) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.Postgres.PostgresqlHost,
		cfg.Postgres.PostgresqlPort,
		cfg.Postgres.PostgresqlUser,
		cfg.Postgres.PostgresqlDbname,
		cfg.Postgres.PostgresqlPassword,
	)

	db, err := sqlx.Connect("pgx", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
