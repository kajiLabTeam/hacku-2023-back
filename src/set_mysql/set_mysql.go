package main

import (
	// "encoding/csv" // 使っていなかったのでコメントアウト
	"fmt"
	// "os" // 使っていなかったのでコメントアウト
	"strings"

	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// タグの候補をDBから読み込む
	tagCandidates := readTagCandidatesFromDB()

	// 分析する文章
	// 大文字　小文字を区別できていない
	userInput := `ドメイン駆動設計とは、その名の通り "ドメイン" の知識にフォーカスした設計手法です。

	ここで言う "ドメイン" とは、「ソフトウェアを使って問題解決しようとしている領域」や「プログラムを適用する対象となる業務領域」のなどを指します。
	具体的には、会計システムにおける「金銭」や「振込処理」、SNSにおける「投稿」や「ユーザー」などが該当します。
	
	これらのドメインを含め、システムが扱う業務仕様やビジネスルールを軸に設計を行い、
	最適な業務実現・課題解決をしていこうという手法をドメイン駆動設計と呼びます。
	
	ざっくり言うと、良いシステムを構築するための設計のベストプラクティスに近い話です。
	業務の要件やビジネスルールに重きをおいた設計をすることで、よりユーザーのニーズを満たした開発を行うことができます。
	`

	// 入力からタグの候補を抽出
	extractedTagCandidates := extractTagCandidates(userInput, tagCandidates)

	// 重複を削除
	extractedTagCandidates = removeDuplicates(extractedTagCandidates)

	// 抽出されたタグの候補を表示
	fmt.Println("Extracted tag candidates:", extractedTagCandidates)
	//すべてのタグ候補を表示
	// fmt.Println("Tag candidates:", tagCandidates)
}

func db_connect() *sql.DB {
	jst, err := time.LoadLocation("Asia/Tokyo")

	c := mysql.Config{
		DBName: "db",
		User:   "user",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "localhost:3306",
		Loc:    jst,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		panic(err.Error())
	}
	return db
}

func readTagCandidatesFromDB() []string {
	// DBに接続
	db := db_connect()
	defer db.Close() // DBのCloseが関数の最後に呼び出されるようにする

	// タグの候補を格納するスライス
	var tagCandidates []string

	// タグの候補をDBから読み込む
	rows, err := db.Query("SELECT tag_name FROM tag_list")
	if err != nil {
		panic(err.Error())
	}

	// タグの候補をスライスに格納
	for rows.Next() {
		var tag string
		err := rows.Scan(&tag) // カラム名を指定
		if err != nil {        // エラー処理
			panic(err.Error()) // エラーが発生したら処理を中断
		}
		tagCandidates = append(tagCandidates, tag) // スライスに格納
	}
	return tagCandidates
}

func extractTagCandidates(text string, tagCandidates []string) []string {
	var extractedTagCandidates []string

	// 入力を単語に分割
	words := strings.Fields(text)

	// 単語がタグの候補と一致するかチェック
	for _, word := range words {
		for _, tagCandidate := range tagCandidates {
			if strings.Contains(strings.ToLower(word), strings.ToLower(tagCandidate)) {
				extractedTagCandidates = append(extractedTagCandidates, tagCandidate)
				// break // 重複を避けるためにbreakを追加
			}
		}
	}

	return extractedTagCandidates
}

func removeDuplicates(tags []string) []string {
	// 重複を削除するためのマップ
	seen := make(map[string]bool)
	result := []string{}

	for _, tag := range tags {
		if !seen[tag] {
			seen[tag] = true
			result = append(result, tag)
		}
	}

	return result
}
