package user

import "time"

type User struct {
	Id        uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	Email     string
	Password  string
	Nickname  string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
