package model

type Post struct {
	ID      uint   `gorm:"primaryKey" gorm:"column:id"`
	Title   string `json:"title" gorm:"column:title"`
	Content string `json:"content" gorm:"column:content"`
	UserId  uint   `json:"user_id" gorm:"column:userid"`
	Tags    []Tag  `gorm:"many2many:post_tag;"`
}

func (Post) TableName() string {
	return "post"
}
