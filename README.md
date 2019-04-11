# gRPC-sample
Template project for microservices using gRPC in Golang. Copy this code to create a new microservice.

## Usage
Run server
```
docker-compose build
# サーバーを立ち上げる
docker-compose up application
# サーバーを叩く
docker-compose up client
```

To add a function, create a `.proto` file and create go codes from it using gRPC.
```
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go

cd proto

# ... write protocolbuffer

protoc --go_out=plugins=grpc:../pb todo.proto
```
Once you have an interface using gRPC, write implementations in the `controllers` and` usecases` layers. The dependency injection is implemented in `depin / injector.go`.

## Layers

### controllers
This layer maps the received request to a structure that the `usecase` layer allows, and calls the methods of the` usecase` layer. Then, the structure returned by the `usecase` layer is mapped to the response and returned. Where to do something other than business logic.

### usecases
This layer executes business logic while calling the `entity` layer based on the structure received from controllers, and returns a value. Only logic is condensed here to make it easy to test.


### entities
This layer defines and manipulates domain models. All the data manipulation and persistence logic using the OR mapper is condensed here, and this layer is called only from the `usecase` layer.

