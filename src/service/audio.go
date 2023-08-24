package service

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func getBinary(s string, style int) ([]byte, error) { //バイナリデータをもらってくる
	str := s
	slice := []rune(str) //ルーンに変換しないと二バイト文字がバグる
	str = string(slice)

	urlParts := []string{fmt.Sprint("http://localhost:50021/audio_query?text=", url.QueryEscape(str), "&speaker=3&preset_id=", style)}
	url_query := strings.Join(urlParts, "")           //URL組み立て
	req, _ := http.NewRequest("POST", url_query, nil) //POSTでリクエスト
	req.Header.Set("accept", "application/json")      //ヘッダをセット
	client := new(http.Client)                        //クライアント生成
	resp, err := client.Do(req)                       //リクエスト
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	url_synth := fmt.Sprint("http://localhost:50021/synthesis?speaker=", style, "&enable_interrogative_upspeak=true") //音声生成用URL
	req_s, _ := http.NewRequest("POST", url_synth, resp.Body)                                                         //POSTでリクエスト
	req_s.Header.Set("accept", "audio/mp3")                                                                           //ヘッダをセット
	req_s.Header.Set("Content-Type", "application/json")                                                              //ヘッダをセット
	resp_s, err := client.Do(req_s)                                                                                   //リクエスト
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	defer resp_s.Body.Close()                             //いらなくなったら閉じる
	buff := bytes.NewBuffer(nil)                          //バッファ生成
	if _, err := io.Copy(buff, resp_s.Body); err != nil { //移す
		log.Printf("error: %v", err)
		return nil, err
	}
	return buff.Bytes(), nil //バッファを返す
}

func makeMp3File(b []byte, user string) (string, error) { //音声を生成する関数
	max := new(big.Int)                  //ファイル名の重複を避けるための乱数
	max.SetInt64(int64(10000000))        //1000万通り
	r, err := rand.Int(rand.Reader, max) //乱数生成
	if err != nil {
		log.Printf("error: %v", err)
		return "", err
	}
	file_name := fmt.Sprintf("%s_%d.mp3", user, r) //ファイル名の重複を避ける
	file, _ := os.Create(file_name)                //ファイル生成
	defer func() {
		file.Close() //終わったら閉じる
	}()
	file.Write(b)         //ファイルにデータを書き込む
	return file_name, nil //ファイルのパスを返す
}

type Dictionary struct {
	Surface       string `json:"surface"`
	Pronunciation string `json:"pronunciation"`
	AccentType    int    `json:"accent_type"`
}

func SetDictionary(r Dictionary) error {
	surface := r.Surface
	pronunciation := r.Pronunciation
	accent_type := r.AccentType
	slices := []rune(pronunciation)
	pronunciation = string(slices)
	urlParts := []string{fmt.Sprint("http://localhost:50021/user_dict_word?surface=",surface,"&pronunciation=",url.QueryEscape(pronunciation),"&accent_type=", accent_type,"&word_type=PROPER_NOUN&priority=6")}
	url_query := strings.Join(urlParts, "")           //URL組み立て
	req, _ := http.NewRequest("POST", url_query, nil) //POSTでリクエスト
	req.Header.Set("accept", "application/json")      //ヘッダをセット
	client := new(http.Client)                        //クライアント生成
	resp, err := client.Do(req)                       //リクエスト
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}
	defer resp.Body.Close()
	return nil
}
