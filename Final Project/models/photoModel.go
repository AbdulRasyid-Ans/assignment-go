package models

type Photo struct {
	Model
	Title    string `gorm:"not null"`
	Caption  string
	PhotoUrl string `gorm:"not null"`
	UserID   uint
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	User *User
}
