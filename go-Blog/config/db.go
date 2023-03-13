package config

import (
	"fmt"
	"tugas4/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDataBaseConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=123456 dbname=goblog port=5432 TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if err := db.AutoMigrate(
		entity.User{},
		entity.Blog{},
		entity.BlogComment{},
	); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	dbSQL.Close()
}
