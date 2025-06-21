package model

type Post struct {
	ID       uint
	Title    string
	UserID   uint
	Comments []Comment `gorm:"foreignKey:PostID"`
}
