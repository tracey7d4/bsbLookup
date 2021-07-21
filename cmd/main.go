package main

import (
	"fmt"
	"github.com/tracey7d4/bsbLookup/cmd/config"
	"github.com/tracey7d4/bsbLookup/proto"
	"github.com/tracey7d4/bsbLookup/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	configs, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configs file: ", err)
	}
	port := configs.Port

	lis, err := net.Listen("tcp", ":"+fmt.Sprintf("%v",port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api := &service.LookupAPI{}
	if err := api.UpdateCache(); err != nil {
		fmt.Println("error reading api cache")
		return
	}
	proto.RegisterBsbLookupServer(s, api)

	fmt.Printf("bsb-lookup started on port %v\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
