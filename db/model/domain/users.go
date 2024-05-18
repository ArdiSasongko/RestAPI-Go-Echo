package domain

import "time"

type User struct {
	UserID    int       `gorm:"column:user_id;primaryKey;autoIncrement"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	Orders    []*Orders `gorm:"foreignKey:user_id_fk;references:user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
