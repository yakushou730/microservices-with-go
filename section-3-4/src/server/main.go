package main

import (
	"log"
	"net"

	pb "github.com/yakushou730/microservices-with-go/section-3-4/src/proto"
	"github.com/yakushou730/microservices-with-go/section-3-4/src/server/WTAServer"
	"github.com/yakushou730/microservices-with-go/section-3-4/src/server/repository"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	myS := &WTAServer.MyWTAServer{
		Repo: repository.NewWTARepository(),
	}
	defer myS.Repo.CloseWTARepository()

	pb.RegisterWTAServer(s, myS)

	reflection.Register(s)

	log.Println("Starting Server.")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
