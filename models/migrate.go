package models

import (
	"log"
)

func init() {
	err := DB.AutoMigrate(
		&Problem{}, &Category{}, &ProblemCategory{},
	)
	if err != nil {
		log.Println("gorm Init Error : ", err)
	}
}
