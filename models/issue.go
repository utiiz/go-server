package models

// type IssueType struct {
// 	ID        uint64 `json:"id" gorm:"primaryKey"`
// 	Type
// 	CreatedAt int    `json:"created_at"`
// 	UpdatedAt int    `json:"updated_at"`
// }

type Issue struct {
	ID        uint64 `json:"id" gorm:"primaryKey"`
	Summary   string `json:"summary"`
	UserID    int    `json:"user_id"`
	CreatedAt int    `json:"created_at" gorm:"<-:create"`
	UpdatedAt int    `json:"updated_at"`
}

func GetIssues() ([]Issue, error) {
	var issues []Issue
	tx := db.Find(&issues)
	if tx.Error != nil {
		return []Issue{}, tx.Error
	}

	return issues, nil
}

func GetIssue(id uint64) (Issue, error) {
	var issue Issue

	tx := db.Where("id = ?", id).First(&issue)
	if tx.Error != nil {
		return Issue{}, tx.Error
	}

	return issue, nil
}

func CreateIssue(issue Issue) error {
	tx := db.Create(&issue)
	return tx.Error
}

func UpdateIssue(issue Issue) error {
	tx := db.Save(&issue)
	return tx.Error
}

func DeleteIssue(id uint64) error {
	tx := db.Unscoped().Delete(&Issue{}, id)
	return tx.Error
}
