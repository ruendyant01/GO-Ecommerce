package database

import (
	"database/sql"
	"fmt"
	"go_config/config"
	"time"
)

type Postgresql struct {
	db *sql.DB
}

func OpenPostgresSqlDatabase(host string, port int, user string, password string, database string) (*Postgresql, error) {
	postgresDb := &Postgresql{
		nil,
	}

	connMaxLifeTime := config.Default().GetInt("db.postgressql.connMaxLifeTime")
	maxIdleConn := config.Default().GetInt("db.postgressql.maxIdleConn")
	maxOpenConn := config.Default().GetInt("db.postgressql.maxOpenConn")
	param := config.Default().GetString("db.postgressql.param")
	postgresDb.Open(Options{
		Host:             host,
		Port:             port,
		Username:         user,
		Password:         password,
		DatabaseName:     database,
		Param:            param,
		ConnMaxIdleConns: maxIdleConn,
		ConnMaxLifetime:  time.Duration(connMaxLifeTime),
		ConnMaxOpenConns: maxOpenConn,
	})

	return postgresDb, nil
}

func (m *Postgresql) Open(options Options) {
	dbs, err := BuildDNS(options)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Opening PostgreSQL database with options: username %s, password %s, database %s", options.Username, options.Password, options.DatabaseName)
	db, err := sql.Open("postgres", dbs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully opened PostgreSQL database with options: username %s, password %s, database %s", options.Username, options.Password, options.DatabaseName)
	db.SetConnMaxIdleTime(options.ConnMaxLifetime)
	db.SetMaxIdleConns(options.ConnMaxIdleConns)
	db.SetMaxOpenConns(options.ConnMaxOpenConns)
	m.db = db
}

func (m *Postgresql) Close() error {
	if m.db != nil {
		err := m.db.Close()
		if err != nil {
			fmt.Printf("Failed to close PostgreSQL database: %v", err)
			return err
		}
	}
	fmt.Printf("Successfully closed PostgreSQL database")
	return nil
}

func (m *Postgresql) Get() interface{} {
	if m.db == nil {
		panic("PostgreSQL database is not open")
	}
	return m.db
}

func (m *Postgresql) Ping() error {
	if m.db == nil {
		panic("PostgreSQL database is not initialized")
	}
	err := m.db.Ping()
	if err != nil {
		return err
	}
	fmt.Println("Successfully pinged PostgreSQL database")
	return nil
}
