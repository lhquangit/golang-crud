package model

type Tag struct {
	ID uint `gorm:"primaryKey gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Posts []Post `gorm:"many2many:post_tags;"`
}

func (Tag) TableName() string {
    return "tag"
}