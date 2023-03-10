package db

import (
	"database/sql"
	"fmt"
	"github.com/OybekAbduvosiqov/movie/config"
	_ "github.com/lib/pq"
)

func NewConnectPostgres(cfg config.Config) (*sql.DB, error) {

	connect := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	)

	conn, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return conn, nil
}
