# gRPC-sample
GolangでgRPCを使ったマイクロサービスのためのテンプレートプロジェクト。ここにコードを書き足してマイクロサービスができていくと嬉しい。

## Usage
とりあえず動かす。
```
docker-compose build
# サーバーを立ち上げる
docker-compose up application
# サーバーを叩く
docker-compose up client
```

機能の追加は、gRPCを用いて`.proto`ファイルを作成し、それをもとに作成する。
```
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```
