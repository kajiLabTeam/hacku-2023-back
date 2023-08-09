## 実行方法
### MySQLコンテナとVoicevoxコンテナの立ち上げ方
1. プロジェクトディレクトリまで移動する
2. コンテナを立ち上げる
```
docker compose up -d
```
#### コンテナが正しく動いているか確認したい場合
```
docker compose logs
```

### どうしてもコンテナがうまく立ち上がらなかった時
以下のコマンドでDocker環境を初期化してください
```
 docker stop $(docker ps -aq)
 docker rm $(docker ps -aq)
 
 docker network prune -f
 docker rmi -f $(docker images --filter dangling=true -qa)
 docker volume rm $(docker volume ls --filter dangling=true -q)
 docker rmi -f $(docker images -qa)
```
