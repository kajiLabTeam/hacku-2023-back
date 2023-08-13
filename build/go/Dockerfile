# goバージョン
FROM golang:1.19.3-alpine
# アップデートとgitのインストール
RUN apk add --update &&  apk add git
# appディレクトリの作成
RUN mkdir /src/app
# ワーキングディレクトリの設定
WORKDIR /src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /src/app

# CMD ["go", "run", "./main.go"]