package grpcservice

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"github.com/antomaat/softwareengineering-test-task/app/db"
	ticketscoreservice "github.com/antomaat/softwareengineering-test-task/app/ticket_score_service"
	pb "github.com/antomaat/softwareengineering-test-task/protos/ticket_score"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func newScoreService() ticketscoreservice.ScoreService {
    db, err := db.NewDatabase("../../database.db")
    if err != nil {
	log.Fatalf("database had a error %v", err)
    }
    return *ticketscoreservice.NewScoreService(db)
}

func runTestServer(ctx context.Context, scoreService ticketscoreservice.TicketScoreService) (pb.TicketScoresClient, func()) {

    bufSize := 1024 * 1024
    lis := bufconn.Listen(bufSize)

    grpcServer := grpc.NewServer()
    ticketServer := &TicketScoreServer{scoreService: scoreService}

    pb.RegisterTicketScoresServer(grpcServer, ticketServer)

    go func() {
	if err := grpcServer.Serve(lis); err != nil {
	    log.Fatalf("Error serving server: %v", err)
	}
    }()

    contextDialer := func(context.Context, string) (net.Conn, error) {
	return lis.Dial()
    }

    conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(contextDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
	log.Fatalf("Error connecting to server: %v", err)
    }

    closer := func() {
	err := lis.Close()
	if err != nil {
	    log.Fatalf("Error closing listener: %v", err)
	}
	grpcServer.Stop()
    }
    client := pb.NewTicketScoresClient(conn)
    return client, closer 
}


func TestGetAggregatedScores(t *testing.T) {
    ctx := context.Background()

    scoreService := newScoreService()
    client, closer := runTestServer(ctx, &scoreService)

    defer closer()

    start := time.Date(2010, time.March, 29, 0, 0, 0, 0, time.UTC)
    end := time.Date(2019, time.March, 30, 0, 0, 0, 0, time.UTC)


    req := &pb.GetAggregatedCategoryScoresRequest{
	Period: &pb.PeriodEntity{
	    StartTime: timestamppb.New(start),
	    EndTime: timestamppb.New(end),
	},
    }

    resp, err := client.GetAggregatedCategoryScores(ctx, req)
    require.NoError(t, err)

    require.NotEmpty(t, resp.Scores)
}

func TestGetOverallQualityScores(t *testing.T) {
    ctx := context.Background()

    scoreService := newScoreService()
    client, closer := runTestServer(ctx, &scoreService)

    defer closer()

    start := time.Date(2019, time.March, 29, 0, 0, 0, 0, time.UTC)
    end := time.Date(2019, time.March, 30, 0, 0, 0, 0, time.UTC)

    req := &pb.GetOverallQualityScoresRequest{
	Period: &pb.PeriodEntity{
	    StartTime: timestamppb.New(start),
	    EndTime: timestamppb.New(end),
	},
    }

    resp, err := client.GetOverallQualityScores(ctx, req)

    require.NoError(t, err)
    require.NotNil(t, resp.Score)
}

func TestGetScoresByTicket(t *testing.T) {
    ctx := context.Background()

    scoreService := newScoreService()
    client, closer := runTestServer(ctx, &scoreService)

    defer closer()

    start := time.Date(2019, time.March, 29, 0, 0, 0, 0, time.UTC)
    end := time.Date(2019, time.March, 30, 0, 0, 0, 0, time.UTC)


    req := &pb.GetScoresByTicketRequest{
	Period: &pb.PeriodEntity{
	    StartTime: timestamppb.New(start),
	    EndTime: timestamppb.New(end),
	},
    }

    resp, err := client.GetScoresByTicket(ctx, req)
    require.NoError(t, err)
    require.NotEmpty(t, resp)
}

func TestGetPeriodOverPeriodScoreChange(t *testing.T) {
    ctx := context.Background()

    scoreService := newScoreService()
    client, closer := runTestServer(ctx, &scoreService)

    defer closer()

    start := time.Date(2019, time.March, 29, 0, 0, 0, 0, time.UTC)
    end := time.Date(2019, time.March, 30, 0, 0, 0, 0, time.UTC)


    req := &pb.GetPeriodOverPeriodScoreChangeRequest{
	PeriodFrom: &pb.PeriodEntity{
	    StartTime: timestamppb.New(start),
	    EndTime: timestamppb.New(end),
	},

	PeriodTo: &pb.PeriodEntity{
	    StartTime: timestamppb.New(start),
	    EndTime: timestamppb.New(end),
	},
    }

    resp, err := client.GetPeriodOverPeriodScoreChange(ctx, req)

    require.NoError(t, err)
    require.NotEqual(t, 0, resp.Change)
}
