package models

type User struct {
	ID        uint64 `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
}

func GetUsers() ([]User, error) {
	var users []User

	// Transaction
	tx := db.Find(&users)
	if tx.Error != nil {
		return []User{}, tx.Error
	}

	return users, nil
}
