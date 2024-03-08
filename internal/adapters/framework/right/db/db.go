package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {

	// Connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database with error: %v", err)
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (da Adapter) CloseDBConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("Connection to db closed: %v", err)
	}
}

func (da Adapter) AddToHistory(answer int32, operation string) error {
	queryString, args, err := sq.Insert("arith_history").Columns("date", "answer", "operation").Values(time.Now(), answer, operation).ToSql()
	if err != nil {
		return err
	}

	_, err = da.db.Exec(queryString, args...)
	if err != nil {
		return err
	}

	return nil
}
