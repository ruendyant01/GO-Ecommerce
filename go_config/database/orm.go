package database

import (
	"database/sql"
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type OrmDb struct {
	OrmInstance *gorm.DB
	Database    Database
}

func OpenOrmWithDatabase(database Database) (*OrmDb, error) {
	if database == nil {
		return nil, errors.New("database cannot be nil")
	}

	ormDb := OrmDb{}

	d := database.Get().(*sql.DB)
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: d,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	ormDb.OrmInstance = gormDb
	ormDb.Database = database

	return &ormDb, nil
}

func OpenOrm(host string, port int, username string, password string, dbname string) (*OrmDb, error) {
	postgresDb, err := OpenPostgresSqlDatabase(host, port, username, password, dbname)
	if err != nil {
		return nil, err
	}
	err = postgresDb.Ping()
	if err != nil {
		return nil, err
	}

	return OpenOrmWithDatabase(postgresDb)
}
