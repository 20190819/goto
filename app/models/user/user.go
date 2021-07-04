package user

import "time"

type User struct {
	Id        uint64    `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	Email     string    `json:"email" gorm:"column:email;required" validate:"required||email"`
	Password  string    `json:"password" gorm:"column:password;required" validate:"required"`
	Nickname  string    `json:"nickname" gorm:"column:nickname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
