package model

type User struct {
	ID       uint
	UserName string
	Posts    []Post `gorm:"foreignKey:UserID"`
}
