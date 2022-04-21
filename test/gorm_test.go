package test

import (
	"fmt"
	"testing"

	"gin-ojsys.cn/config"
	"gin-ojsys.cn/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGormTest(t *testing.T) {
	db, err := gorm.Open(mysql.Open(config.DatabaseSetting.Url), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.Problem, 0)
	err = db.Find(&data).Error
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range data {
		fmt.Printf("Problom ==> %v \n", v)
	}
}
