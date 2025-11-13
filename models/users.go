package models

type User struct {
	ID           uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	FullName     string `gorm:"type:varchar(255);notNull" json:"full_name"`
	Email        string `gorm:"type:varchar(255);uniqueIndex;notNull" json:"email"`
	PasswordHash string `gorm:"type:varchar(255);notNull" json:"-"`
	Role         string `gorm:"type:varchar(50);default:user;notNull" json:"role"`
}
