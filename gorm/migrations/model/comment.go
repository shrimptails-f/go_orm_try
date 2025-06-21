package model

type Comment struct {
	ID      uint
	PostID  uint
	Content string
	Replies []Reply `gorm:"foreignKey:CommentID"`
}
