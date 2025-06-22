![Version](https://img.shields.io/badge/Version-1.0.0-green)
# プロジェクトの概要説明
Go ORMの比較用です
# 環境構築手順
## [VsCode使用者向け](./docs/VsCodeDevContainer.md)
## 言語
* Go1.24
## DB
* MySQL8.0
# ORMごとの実装パターン
前提
- マイグレーションは別で実装するものとする。
## gorm使用
並列にテストを実行できるように、スキーマ名にnanoidを付与して作成する実装例も含んでいます。<br>
・スキーマ名にnanoidを含める実装例<br>
gorm/mysql/mysql.goのCreateNewTestDB関数

### テーブル追加手順
1. gorm/migrations/modelに構造体を追加
### サンプルデータ追加手順
1. gorm/seeder/seeders/にseederの関数を追加
2. gorm/seeder/main.goに2の関数を追加
### テーブル作成
```bash
task gorm-migration-create
```
### テーブル削除
```bash
task gorm-migration-drop
```
### テーブル作成・削除
```bash
task gorm-migration-fresh
```
### サンプルデータ投入
```bash
task gorm-seed
```
### データ取得処理実装例
以下を参照してください。
gorm/main.go
```
go run gorm/main.go
```
### テスト時のデータ投入例(gorm)

<details>
<summary>実装例</summary>

``` go
// CreateData はテストデータを作成します。
func CreateData(t *testing.T, conn *pkgmysql.MySQL) {
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
}
```
</details>

### テスト時のデータ投入例(fixture)
gorm/fixture/fixture.goで定義しておけばこんな感じでリレーションIDやNameなどテストで意識したくない情報を無視してできます。<br>
だいぶ前に実装したっきりなのでちょっと間違ってるかも。
<details>
<summary>実装例</summary>

``` go
// CreateData はテストデータを作成します。
func CreateData(t *testing.T, conn *pkgmysql.MySQL) {
	f := fixture.Build(t,
		user1.
		Connect(chatRoom).
		Connect(
			fixture.ChatMessage(func(cm *model.ChatMessage) {
				cm.Message = "テスト1"
				cm.CreatedAt = timeDate
			}),
		),
		Connect(chatRoom2).
		Connect(
			fixture.ChatMessage(func(cm *model.ChatMessage) {
				cm.Message = "テスト4"
				cm.CreatedAt = timeDate
			}),
		),
	)

	f.Setup(t, conn)
}
```
</details>

### 懸念点
1~5はActiveRecord的なライブラリで共通して言えることではある。<br>
1. Find First Takeなどの取得関数の挙動について理解が必要。どのライブラリでも言えることではあるが。。。
2. エラーハンドリングでコツがいる。gorm.ErrRecordNotFoundは許容するなど実装パターンを掴むひつようがある。
3. 複雑なクエリは結局SQLベタ書きの実装方法に逃げざるを得ない。
4. 1ドメインの関連テーブルが多い場合VIEWを用意・利用することを検討しないといけない
5. 開発メンバーが多い場合コンフリクトが発生しやすい。テーブル定義の構造体はあらかじめ取り込むなどの運用が必要になる。
### 利点
1. ネットに情報が転がっている
2. ActiveRecord的でとっつきやすい（新人に教えやすい）
3. gormだけで大体完結する。
4. テーブルの追加が容易
5. テストでのテーブル作成が簡単
6. データ作成時のリレーション(ID)管理が容易。
7. 開発スピードが速い（初期構築が容易）
8. 型安全性：タグから自動でマッピングしてくれる
## sqlc
練習を兼ねてgolang-migrate/migrateを使用しています。
### テーブル追加手順
1. sqlc/migrationsにテーブル作成クエリを追加
2. sqlc/queriesにデータ取得クエリを追加
3. sqlc generate
### テーブル作成
```bash
task sqlc-migration-create
```
### テーブル削除
```bash
task sqlc-migration-drop
```
### テーブル作成・削除
```bash
task sqlc-migration-fresh
```
### データ取得処理実装例
以下を参照してください。
sqlc/main.go
```
go run sqlc/main.go
```
### 懸念点
1. gormのようにネストしてサンプルデータの作成ができないため、サンプルデータのリレーション管理が煩雑になる
2. 開発メンバーが多い場合似たSQLが大量生産される。SQLの取得内容と命名が別途管理が必要になるかも?
3. 開発メンバーが多い場合コンフリクトが発生しやすい。基本自動生成ファイルなので、そこまで困らないはずだが。。。
4. クエリの結果がそのままなのでネストされた構造が欲しい場合手動で詰め替えなければならない。
5. 扱いが独特でとっつきにくい。sqlc.yamlの書き方、構造体の自動生成ルールの理解など。
### 利点
1. ネットに情報が転がっている
2. 型安全性：Goでの型チェックがSQLベースで保証される
3. SQLをベースにデータを格納する構造体を自動で作成できる。
4. テーブル正規化を行うと自然と複雑なクエリとなりやすいが、クエリ前提なのでつらさがない
## 今のところの結論
gorm golang-migrate/migrateの併用<br>
- マイグレーションはgolang-migrate/migrateで行う
- 処理ではgormを使用
- 関連テーブルが多い場合はViewを使用して1つのドメイン(構造体)として扱う。
