FROM python:3.11.3-alpine

RUN apk add --update &&  apk add git

# appディレクトリの作成
RUN mkdir -p /src/app
# ワーキングディレクトリの設定
WORKDIR /src/app

# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /src/app

# pipでインストールしたいモジュールをrequirements.txtに記述しておいて、
# コンテナ内でpipにインストールさせる
# requirements.txtの書き方は[pip freeze]コマンドから参考に出来る
COPY requirements.txt /src/app
# pipのアップデート
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

CMD [ "python", "app.py" ]
