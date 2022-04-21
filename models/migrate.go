package models

import (
	"log"
)

func init() {
	err := DB.AutoMigrate(
		&Problem{}, &Category{},
	)
	if err != nil {
		log.Println("gorm Init Error : ", err)
	}
}
