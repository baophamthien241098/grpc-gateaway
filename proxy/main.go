package main

import (
	"context"
	hello "grpc-gateway/hello"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	//RegisterHelloServiceHandlerFromEndpoint dùng để chuyển request từ rest sang grpc
	err := hello.RegisterHelloServiceHandlerFromEndpoint(
		ctx, mux, "localhost:5000",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatal(err)
	}
	http.ListenAndServe(":8080", mux)
}
