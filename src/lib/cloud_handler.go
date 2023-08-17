package lib

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/kajiLabTeam/hacku-2023-back/conf"

	"google.golang.org/api/option"
)

var app *firebase.App

func init() {
	f := conf.GetMysqlConfig()
	conf := &firebase.Config{
		StorageBucket: f.GetString("firebase.bucket"),
	}
	opt := option.WithCredentialsFile(f.GetString("firebase.path"))
	var err error
	app, err = firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		log.Fatalln(err)
	}
}

func CloudConnect() (*storage.BucketHandle, error) {
	client, err := app.Storage(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return bucket, nil
}
