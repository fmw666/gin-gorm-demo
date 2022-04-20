package models

func init() {
	DB.AutoMigrate(&Problem{})
}
