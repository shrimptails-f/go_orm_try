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
## gorm使用
また、並列にテストを実行できるように、スキーマ名にnanoidを付与して作成する実装例も含んでいます。<br>
・スキーマ名にnanoidを含める実装例<br>
gorm/mysql/mysql.goのCreateNewTestDB関数

### テーブル追加手順
1. gorm/migrations/modelに構造体を追加
### サンプルデータ追加手順
1. /gorm/seeder/seeders/にseederの関数を追加
2. /data/gorm/seeder/main.goに2の関数を追加
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
### 懸念点
1~5はActiveRecord的なライブラリで共通して言えることではある。<br>
6は致命的。だがgolang-migrate/migrateを併用すれば解決できる<br>
1. Find First Takeなどの取得関数の挙動について理解が必要。どのライブラリでも言えることではあるが。。。
2. エラーハンドリングでコツがいる。gorm.ErrRecordNotFoundは許容するなど実装パターンを掴むひつようがある。
3. 複雑なクエリは結局SQLベタ書きの実装方法に逃げざるを得ない。
4. 1ドメインの関連テーブルが多い場合VIEWを用意・利用することを検討しないといけない
5. 開発メンバーが多い場合コンフリクトが発生しやすい。テーブル定義の構造体はあらかじめ取り込むなどの運用が必要になる。
6. gormだけだとマイグレーション管理ができない。Laravel Railsのように既存に追加する仕組みが必要になる。
7. 6はgolang-migrate/migrateを併用すれば解決できるが構造体と二重管理になる。<br>だが、クリーンアーキテクチャならdomainとDB定義は分けるので二重管理にでも問題ない。
### 利点
1. ネットに情報が転がっている
2. ActiveRecord的でとっつきやすい（新人に教えやすい）
3. gormだけで大体完結する。
4. テーブルの追加が容易
5. テストでのテーブル作成が簡単
6. データ作成時のリレーション(ID)管理が容易。
7. 開発スピードが速い（初期構築が容易）
## sqlc golang-migrate/migrateを使用
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