package settings

import (
	"crud-golang/configuration"
	"crud-golang/infra/db/entities"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Conn struct {
	DB *gorm.DB
}

func Connect(
	dbConfig configuration.Database) (*Conn, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	conn := new(Conn)
	conn.DB = db

	conn.DB.AutoMigrate(
		&entities.User{},
		&entities.UserAddres{},
	)

	return conn, nil
}
