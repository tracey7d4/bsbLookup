package main

import (
	"fmt"
	"github.com/tracey7d4/bsbLookup/cmd/config"
	"log"
)

func main() {
	configs, err := config.LoadConfig()

	if err != nil {
		log.Fatal("Error loading configs file: ", err)
	}
	port := configs.Port
	fmt.Printf("bsb-lookup started on port %v\n", port)
	//lis, err := net.Listen("tcp ", ":"+fmt.Sprintf("%v",port))
	//if err != nil {
	//	log.Fatal("failed to listen: %v", err)
	//}


}
