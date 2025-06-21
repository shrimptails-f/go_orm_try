package seeders

import (
	"business/gorm/migrations/model"
	"log"

	"gorm.io/gorm"
)

// CreateUsereToReply はメールのサンプルデータを投入する。
func CreateUsereToReply(tx *gorm.DB) error {
	user := model.User{
		UserName: "test_user",
		Posts: []model.Post{
			{
				Title: "Post 1",
				Comments: []model.Comment{
					{
						Content: "Comment A",
						Replies: []model.Reply{
							{Content: "Reply A-1"},
							{Content: "Reply A-2"},
						},
					},
					{
						Content: "Comment B",
						Replies: []model.Reply{
							{Content: "Reply B-1"},
						},
					},
				},
			},
		},
	}

	if err := tx.Create(&user).Error; err != nil {
		log.Fatal(err)
	}

	return nil
}
