### ポイント
- 起動した事を知りたいコンテナにヘルスチェックを定義する
- そのヘルスチェックを depends_on で監視して、起動完了後にメッセージを出力するコンテナを追加する

### 初期設定
```
$ docker-compose build
```

### 通常の Detach モードでのサーバー起動
```
$ docker-compose up -d server
Starting launch_message_server_1 ... done
```
すぐに `Starting launch_message_server_1 ... done` のメッセージが出力されるが  
コンテナ内でサーバーが起動するまでは疎通できない。
```
$ curl http://localhost:8080/healthcheck
curl: (56) Recv failure: Connection reset by peer
```
しばらく待つと疎通できるようになる。
```
$ curl http://localhost:8080/healthcheck
{"message":"server is ready"}
```

### サーバーの起動を待つ Detach モード起動
```
$ docker-compose run server_watcher
Creating network "launch_message_default" with the default driver
Creating launch_message_server_1 ... done
```
`server_watcher` 自体は foreground で起動させる。  
`server_watcher` を `server` のヘルスチェック完了に depends_on させているため  
サーバー起動完了後に `SERVER LAUNCHED!` が出力される。
```
Creating launch_message_server_watcher_run ... done
SERVER LAUNCHED!
```
従ってこのメッセージが出力された後 (= コンソールが返ってきた後) は確実に疎通できる。
```
$ curl http://localhost:8080/healthcheck
{"message":"server is ready"}
```
この時 `server` はバックグラウンドで実行されている。  
リファレンスに記載は見当たらないが、恐らく depends_on で起動されたサービスは  
常に Detach モードになるのだと思う。
