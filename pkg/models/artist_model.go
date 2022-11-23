package models

type Artist struct {
	BaseModel
	Name        string `gorm:"not null" json:"name"`
	Image       string `gorm:"not null" json:"image"`
	Description string `gorm:"not null" json:"description"`
	BirthDate   string `gorm:"not null" json:"birthDate"`
	Songs       []Song `gorm:"many2many:songs_artists"`
}

func init() {
	db.AutoMigrate(&Artist{})
}
