package user

import "time"

type User struct {
	Id        uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	Email     string `json:"email" gorm:"column:email;required"`
	Password  string `json:"password" gorm:"column:password;required"`
	Nickname  string `json:"nickname" gorm:"column:nickname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
