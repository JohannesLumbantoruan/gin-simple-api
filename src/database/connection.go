package database

import (
	// "database/sql"
	// _ "github.com/lib/pq"

	// "log"

	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func DB() *sql.DB {
// 	connStr := "postgresql://johndoe:johndoe@127.0.0.1/simpleapi?sslmode=disable"

// 	DB, err := sql.Open("postgres", connStr)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return DB
// }

func DB() *gorm.DB {
	dsn := "host=localhost user=johndoe password=johndoe dbname=simpleapi port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(`Gagal konek ke database`)
	}

	fmt.Println("Berhasil konek ke database")

	return db
}