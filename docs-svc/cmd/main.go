package cmd

import (
	"fmt"
	"github.com/lekan-pvp/gokee/docs-svc/pkg/config"
	"github.com/lekan-pvp/gokee/docs-svc/pkg/db"
	"github.com/lekan-pvp/gokee/docs-svc/pkg/pb"
	"github.com/lekan-pvp/gokee/docs-svc/pkg/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listen", err)
	}

	fmt.Println("Document Svc on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterDocumentServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
