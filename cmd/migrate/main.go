package main

import (
	"log"

	model "github.com/rlaope/evergreen/intenral/model/kn"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "user:pass@tcp(localhost:3306)/evergreen?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB 연결 실패:", err)
	}

	if err := db.AutoMigrate(
		&model.Forest{},
		&model.Tree{},
		&model.Comment{},
	); err != nil {
		log.Fatal("마이그레이션 실패:", err)
	}

	log.Println("✅ 마이그레이션 완료!")
}
