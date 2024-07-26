package repository_adapters

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection struct {
	DB *gorm.DB
}

func (DBConnection *DBConnection) Connect() error {
	var err error

	if DBConnection.DB == nil {
		DBConnection.DB, err = gorm.Open(postgres.Open("host=localhost user=andres password=123 dbname=people_crud port=5432 sslmode=disable"))
	}

	return err
}
