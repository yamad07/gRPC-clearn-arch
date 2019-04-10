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

cd proto

# ... write protocolbuffer

protoc --go_out=plugins=grpc:../pb todo.proto
```
インターフェースができたら、`controllers`レイヤーと`usecases`レイヤーに実装を書く。依存性の注入は`depin/injector.go`で行う。

## Layers

### controllers
このレイヤーは、受け取ったリクエストを`usecase`レイヤーが許容する構造体にマッピングし、`usecase`レイヤーのメソッドを呼び出す。そして、`usecase`レイヤーが返してくれた構造体を、レスポンスにマッピングして返す。ビジネスロジック以外のところをやるところ。

### usecases
このレイヤーは、controllersから受け取った構造体をもとに`entity`レイヤーを呼び出しながらビジネスロジックを遂行し、値を返す。テストがしやすいように、ロジックだけがここに凝縮される。


### entities
このレイヤーは、ドメインモデルの定義や操作を行う。ORマッパーを用いたデータ操作、永続化のロジックは全てここに凝縮され、このレイヤーは`usecase`レイヤーのみから呼ばれる。

