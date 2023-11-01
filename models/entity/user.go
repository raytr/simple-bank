package entity

type User struct {
	BaseModel
	FullName     string `gorm:"type:varchar(255);not null"`
	Username     string `gorm:"type:varchar(255);unique;not null"`
	HashPassword string `gorm:"type:varchar(255);not null"`
	Salt         string `gorm:"type:varchar(255);not null"`
}
