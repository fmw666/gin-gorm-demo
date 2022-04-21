package models

import (
	"log"
)

func init() {
	err := DB.AutoMigrate(
		&Problem{},
	)
	if err != nil {
		log.Println("gorm Init Error : ", err)
	}
}
