package models

type Account struct {
	BaseModel
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Active   bool   `gorm:"not null;default:false"`
	Banned   bool   `gorm:"not null;default:false"`
}

func init() {
	db.AutoMigrate(&User{})
}

func CreateAccount(account Account) error {
	result := db.Create(&account)
	return result.Error
}
