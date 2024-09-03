package main

import (
	"log"
	"net"

	"github.com/antomaat/softwareengineering-test-task/app/db"
	grpcservice "github.com/antomaat/softwareengineering-test-task/app/grpc_service"
	"github.com/antomaat/softwareengineering-test-task/app/ticket_score_service"
	"google.golang.org/grpc"
)

var databaseUrl = "database.go"
var port = ":9000"

func main() {
    database, err := db.NewDatabase(databaseUrl)
    if err != nil {
	log.Fatalf("Failed to initialize the Database %s", databaseUrl)
    }
    defer database.Close()

    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    } 

    grpcServer := grpc.NewServer()

    scoreService := ticketscoreservice.NewScoreService(database)

    grpcservice.NewTicketScoreServer(grpcServer, scoreService)

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve grpc:%s", err)
    }

    log.Printf("The grpc server is running on port %s", port)

}
