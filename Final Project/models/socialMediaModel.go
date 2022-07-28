package models

type SocialMedia struct {
	Model
	Name           string `gorm:"not null"`
	SocialMediaUrl string `gorm:"not null"`
	UserID         uint

	User *User
}
