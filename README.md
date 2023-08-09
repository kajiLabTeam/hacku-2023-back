## MySQLコンテナとVoicevoxコンテナの立ち上げ方
1. プロジェクトディレクトリまで移動する
2. コンテナを立ち上げる
```
docker compose up -d
```
### コンテナが正しく動いているか確認したい場合
```
docker compose logs
```

## どうしてもコンテナがうまく立ち上がらなかった時
以下のコマンドでDocker環境を初期化してください
```
 docker stop $(docker ps -aq)
 docker rm $(docker ps -aq)
 
 docker network prune -f
 docker rmi -f $(docker images --filter dangling=true -qa)
 docker volume rm $(docker volume ls --filter dangling=true -q)
 docker rmi -f $(docker images -qa)
```

## GoとMySQLの接続時に注意すること
**mysql.ymlを変更してください**
### ローカルで実行したGoからDockerコンテナ内のMySQLに接続する時

```mysql.yml
mysql:
  user: root
  password: admin
  protocol: tcp(localhost:3306)
  dbname: data_sets
```

### DockerコンテナのGoからDockerコンテナ内のMySQLに接続する時
```mysql.yml
mysql:
  user: root
  password: admin
  protocol: tcp(HackU2023_Nagoya_DB:3306)
  dbname: data_sets
```
