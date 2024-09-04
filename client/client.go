package main

import (
	"context"
	"log"
	"time"

	pb "github.com/antomaat/softwareengineering-test-task/protos/ticket_score"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
    conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatal("failed to connect to the gRPC server")
    }

    defer conn.Close()

    c := pb.NewTicketScoresClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)

    defer cancel()

    period := &pb.PeriodEntity{StartTime: timestamppb.New(time.Date(2019, 8, 20, 0, 0, 0, 0, time.UTC)), EndTime: timestamppb.New(time.Date(2019, 8, 30, 0, 0, 0, 0, time.UTC))}

    aggregateCategoryScoresRequest := &pb.GetAggregatedCategoryScoresRequest{Period: period}

    response, err := c.GetAggregatedCategoryScores(ctx, aggregateCategoryScoresRequest)
    if err != nil {
	log.Printf("GetAggregatedCategoryScores error: ", err)
    } else {
	log.Printf("%v", response)
    }

    scoresByTicketResponse, err := c.GetScoresByTicket(ctx, &pb.GetScoresByTicketRequest{Period: period})
    if err != nil {
	log.Printf("GetScoresByTicket error: ", err)
    } else {
	log.Printf("%v", scoresByTicketResponse)
    }
    
    overallQualityScores, err := c.GetOverallQualityScores(ctx, &pb.GetOverallQualityScoresRequest{Period: period})
    if err != nil {
	log.Printf("GetOverallQualityScores error: ", err)
    } else {
	log.Printf("%v", overallQualityScores)
    }
}
