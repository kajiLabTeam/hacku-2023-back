package main

// gormを使ってDBに接続する
import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/cheggaaa/pb/v3"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// DBに接続
	db, err := gorm.Open(getDBConfig())
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}

	// DDLの実行
	create_db(db)

	// CSVファイルの中身をmysqlにimportする
	importCSV(db)

	// DBに入れられたデータの確認
	checkData(db)
}

func getDBConfig() (string, string) {
	DBMS := "mysql"
	USER := "root"
	PASS := "admin"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "hacku_nagoya"
	OPTION := "charset=utf8&parseTime=True&loc=Local"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION

	return DBMS, CONNECT
}

func create_db(db *gorm.DB) {
	// root_db.Exec("set global local_infile = 1;")
	db.Exec("CREATE DATABASE IF NOT EXISTS db;")
	db.Exec("CREATE TABLE IF NOT EXISTS db.tag_list(id INT NOT NULL AUTO_INCREMENT, tag_name VARCHAR(255) NOT NULL, PRIMARY KEY (id));")
	db.Exec("TRUNCATE TABLE db.tag_list;")
	db.Exec("CREATE USER IF NOT EXISTS 'user' IDENTIFIED BY 'password';")
	db.Exec("GRANT ALL PRIVILEGES ON db.* TO 'user';")
	db.Exec("FLUSH PRIVILEGES;")
}

func importCSV(db *gorm.DB) {
	db.Exec("use db;")
	// CSVファイルを読み込む
	csvFile, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	csvData, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// transactionを開始
	tx := db.Begin()

	// progress bar を表示
	bar := pb.StartNew(len(csvData))
	bar.SetMaxWidth(120)
	bar.SetTemplateString(`{{string . "prefix"}}{{counters . }} {{bar . }} {{percent . }} {{speed . }} {{string . "suffix"}}`)
	bar.Set("prefix", "importing...")
	bar.Set("suffix", "done")

	// CSVファイルの中身をmysqlにinsertする
	for _, data := range csvData {
		bar.Increment()
		tx.Exec("INSERT INTO tag_list (tag_name) VALUES (?);", data[0])
	}
	bar.Finish()

	// transactionを終了
	tx.Commit()

}

func checkData(db *gorm.DB) {
	// DBに入れられたデータの確認
	db.Exec("use db;")

	// 件数の取得
	var count int
	db.Table("tag_list").Count(&count)
	fmt.Println("tag count:", count)
}
