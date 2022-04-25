package models

type User struct {
	ID        uint64 `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
}

func GetUsers() ([]User, error) {
	var users []User

	tx := db.Find(&users)
	if tx.Error != nil {
		return []User{}, tx.Error
	}

	return users, nil
}

func GetUser(id uint64) (User, error) {
	var user User

	tx := db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return User{}, tx.Error
	}

	return user, nil
}

func CreateUser(user User) error {
	tx := db.Create(&user)
	return tx.Error
}

func UpdateUser(user User) error {
	tx := db.Save(&user)
	return tx.Error
}

func DeleteUser(id uint64) error {
	tx := db.Unscoped().Delete(&User{}, id)
	return tx.Error
}
