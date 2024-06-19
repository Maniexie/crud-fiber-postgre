package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	const POSTGRESQL = "host=localhost user=postgres password=123 dbname=crud_go port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	DSN := POSTGRESQL

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		panic("koneksi tidak terhubung")
	}

	fmt.Println("Koneksi terhubung")
}
