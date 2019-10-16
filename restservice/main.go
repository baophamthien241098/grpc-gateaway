package main

import (
	"context"
	hello "grpc-gateway/hello"

	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type RestServiceImp struct{}

func (r *RestServiceImp) Get(ctx context.Context, req *hello.GetAccResquest) (*hello.GetAccResponse, error) {
	ten := req.GetId()
	log.Print(ten)
	return &hello.GetAccResponse{
		Name:  "ThienBao",
		Pass:  "123123123",
		Phone: "0123456789",
	}, nil
}
func (r *RestServiceImp) Post(ctx context.Context, req *hello.PostResquest) (*hello.PostResponse, error) {
	return &hello.PostResponse{
		Id: "10",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "Localhost:5000")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, &RestServiceImp{})
	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
