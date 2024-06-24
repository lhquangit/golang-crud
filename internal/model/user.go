package model

type User struct {
	ID    uint   `gorm:"primaryKey" gorm:"column:id"`
	Name  string `json:"name" gorm:"column:name"`
	Posts []Post `gorm:"foreignKey:UserId"`
}

func (User) TableName() string {
	return "user"
}
