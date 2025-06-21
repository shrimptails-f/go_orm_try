package main

import (
	"fmt"
	"log"

	"business/gorm/migrations/model"
	"business/gorm/mysql"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	connect, err := mysql.New()
	if err != nil {
		log.Fatal(err)
	}
	db := connect.DB
	silentDB := db.Session(&gorm.Session{
		Logger: db.Config.Logger.LogMode(logger.Silent),
	})

	getComments(silentDB)
}

func getComments(db *gorm.DB) {
	var users1 []model.User
	err := db.
		Preload("Posts.Comments.Replies").
		Find(&users1).
		Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("---- 条件なしパターン ----- \n")
	printDetail(users1)
	// 出力
	// User: test_user
	//   Post: Post 1
	//     Comment: Comment A
	//       replie: Comment: Reply A-1
	//       replie: Comment: Reply A-2
	//     Comment: Comment B
	//       replie: Comment: Reply B-1

	var users2 []model.User
	err = db.
		Joins("JOIN posts ON posts.user_id = users.id").
		Where("posts.title IN ?", []string{"aaa"}).
		Preload("Posts.Comments.Replies").
		Find(&users2).
		Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Printf("---- postに条件があり、postを持たないユーザー情報を取得しないパターン ----- \n")
	printDetail(users2)
	// 出力
	// User: test_user
	//   Post: Post 1
	//     Comment: Comment A
	//       replie: Comment: Reply A-1
	//       replie: Comment: Reply A-2
	//     Comment: Comment B
	//       replie: Comment: Reply B-1

	var users3 []model.User
	err = db.
		Preload("Posts", "title IN ?", []string{"aaa"}).
		Preload("Posts.Comments.Replies").
		Find(&users3).
		Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Printf("---- postに条件があり、postを持たないユーザー情報を取得するパターン ----- \n")
	printDetail(users3)
	// 出力
	// User: test_user

	var users4 []model.User
	err = db.
		Preload("Posts").
		Preload("Posts.Comments", "content IN ?", []string{"Comment A"}).
		Preload("Posts.Comments.Replies").
		Find(&users4).
		Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Printf("---- commentに条件ありパターン ----- \n")
	printDetail(users4)
	// 出力
	// User: test_user
	//   Post: Post 1
	//     Comment: Comment A
	//       replie: Comment: Reply A-1
	//       replie: Comment: Reply A-2

	var users5 []model.User
	err = db.
		Preload("Posts").
		Preload("Posts.Comments.Replies", "content IN ?", []string{"Reply A-1"}).
		Find(&users5).
		Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Printf("---- replie に条件ありパターン ----- \n")
	printDetail(users5)
	// 出力
	// User: test_user
	//   Post: Post 1
	//     Comment: Comment A
	//       replie: Comment: Reply A-1
	//     Comment: Comment B
}

func printDetail(users []model.User) {
	for _, user := range users {
		fmt.Printf("User: %s\n", user.UserName)
		for _, post := range user.Posts {
			fmt.Printf("  Post: %s\n", post.Title)
			for _, comment := range post.Comments {
				fmt.Printf("    Comment: %s\n", comment.Content)
				for _, replie := range comment.Replies {
					fmt.Printf("      replie: Comment: %s\n", replie.Content)
				}
			}
		}
	}
}
