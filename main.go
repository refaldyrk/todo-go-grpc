package main

import (
	"net"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/refaldyrk/todo-go-grpc/adapter"
	"github.com/refaldyrk/todo-go-grpc/connector/todo"
	"github.com/refaldyrk/todo-go-grpc/service"
	"google.golang.org/grpc"
)

func main() {

	//Cache Service
	caches := cache.New(24*time.Hour, 24*time.Hour)
	serviceCache := service.NewCacheService(caches)
	adapters := adapter.NewAdapter(serviceCache)

	list, err := net.Listen("tcp", ":8001")
	if err != nil {
		panic(err.Error())
	}

	serverGrpc := grpc.NewServer()

	todo.RegisterTodoServiceServer(serverGrpc, adapters)

	err = serverGrpc.Serve(list)
	if err != nil {
		panic(err.Error())
	}
}
