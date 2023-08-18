package integrations

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kajiLabTeam/hacku-2023-back/lib"
)

func Storage(fn string) (string, error) {
	localFilePath := fn         // ローカルのファイル名
	remotePath := "audio/" + fn // Bucketに保存されるファイル名

	// 音声ファイルを開く
	f, err := os.Open(localFilePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	// storageにファイルをアップロード
	ctx := context.Background()
	bucket, err := lib.CloudConnect()
	if err != nil {
		return "", fmt.Errorf("error getting bucket: %v", err)
	}
	storageRef := bucket.Object(remotePath)
	wc := storageRef.NewWriter(ctx)
	if _, err := wc.Write([]byte{}); err != nil {
		return "", fmt.Errorf("error writing to storage: %v", err)
	}

	// ファイルの内容をコピー
	if _, err := f.Seek(0, 0); err != nil {
		log.Fatalf("Error seeking in file: %v\n", err)
	}
	if _, err := io.Copy(wc, f); err != nil {
		log.Fatalf("Error copying file to storage: %v\n", err)
	}

	// アップロード完了
	if err := wc.Close(); err != nil {
		log.Fatalf("Error closing storage writer: %v\n", err)
	}

	return remotePath, nil
}
