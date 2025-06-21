package main

import (
	"business/gorm/mysql"
	"business/gorm/seeder/seeders"
	"errors"
	"fmt"
	"os"

	"gorm.io/gorm"
)

func main() {
	// コマンドラインのバリデーション
	err := CheckArgs()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	var conn *mysql.MySQL
	if os.Args[1] == "dev" {
		conn, err = mysql.New()
	} else if os.Args[1] == "test" {
		conn, err = mysql.NewTest()
	}
	if err != nil {
		panic(err)
	}

	// connがnilでないことを確認
	if conn == nil || conn.DB == nil {
		panic("データベース接続が初期化されていません。")
	}

	tx, cleanUP := mysql.Transactional(conn.DB)
	defer cleanUP()

	err = Seed(tx)
	if err != nil {
		tx.Error = err
		fmt.Printf("データ投入中にエラーが発生しました。\n")
		return
	}

	fmt.Printf("正常に終了しました。\n")
}

// CheckArgs はコマンドライン引数を確認する。
func CheckArgs() error {
	if len(os.Args) != 2 {
		return errors.New("期待している引数は1つです。引数を確認してください。")
	}

	if os.Args[1] != "dev" && os.Args[1] != "test" {
		return errors.New("第一引数が期待している語群は以下の通りです。\n1:dev\n2:test")
	}

	return nil
}

// Seed　はサンプルデータを投入する。
func Seed(tx *gorm.DB) error {
	if err := seeders.CreateUsereToReply(tx); err != nil {
		return err
	}

	return nil
}
