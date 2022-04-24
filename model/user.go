package model

type User struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
}
