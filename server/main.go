package main

import (
	"context"
	"grpc_demo/proto"
	"net"
  "fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{
  proto.UnimplementedAddServiceServer
}
func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}
func (s *server)Add(ctx context.Context, in *proto.Request) (*proto.Response, error) {
	a, b := in.GetA(), in.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *server) 	Multiply(ctx context.Context, in *proto.Request) (*proto.Response, error){
	a, b := in.GetA(), in.GetB()

	result := a * b

	return &proto.Response{Result: result}, nil
}

func (s *server) 	Stream(in *proto.Request, srv proto.AddService_StreamServer) ( error){
  resp := proto.Response{Result: in.GetA()}
	if err := srv.Send(&resp); err != nil {
				fmt.Printf("send error %v", err)
  }
  resp = proto.Response{Result: in.GetB()}
	if err := srv.Send(&resp); err != nil {
				fmt.Printf("send error %v", err)
  }
  fmt.Println("Send Array Done")
  return nil
}

func (s *server) 	Array(ctx context.Context, in *proto.Request) (*proto.ArrayResponse ,error){
  array:=make([]*proto.Response, 2)
  array[0] = &proto.Response{Result:in.GetA()}
  array[1] = &proto.Response{Result:in.GetB()}
  return &proto.ArrayResponse{Items:array }, nil
}
