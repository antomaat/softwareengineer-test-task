package grpcservice

import (
	"context"
	"log"

	ticketscoreservice "github.com/antomaat/softwareengineering-test-task/app/ticket_score_service"
	pb "github.com/antomaat/softwareengineering-test-task/protos/ticket_score"
	"google.golang.org/grpc"
)

type TicketScoreServer struct {
    pb.UnimplementedTicketScoresServer
    scoreService ticketscoreservice.TicketScoreService 
}

func NewTicketScoreServer(grpcServer grpc.ServiceRegistrar, scoreService ticketscoreservice.TicketScoreService) *TicketScoreServer {
    server := &TicketScoreServer{scoreService: scoreService}
    pb.RegisterTicketScoresServer(grpcServer, server)
    return server
}


func (s *TicketScoreServer) GetAggregatedCategoryScores(ctx context.Context, request *pb.GetAggregatedCategoryScoresRequest) (*pb.GetAggregatedCategoryScoresResponse, error) {
    start := request.Period.StartTime.AsTime()
    end := request.Period.EndTime.AsTime()
    
    scores, err := s.scoreService.GetAggregatedCategoryScores(start, end)
    if err != nil {
	log.Printf("%s", err.Error())
	return nil, err
    }

    mappedScores := mapToScores(scores)

    aggregatedCategoryScores := pb.GetAggregatedCategoryScoresResponse{
	Scores: mappedScores,
    }

    return &aggregatedCategoryScores, nil
}

func (s *TicketScoreServer) GetOverallQualityScores(ctx context.Context, request *pb.GetOverallQualityScoresRequest) (*pb.GetOverallQualityScoresResponse, error) {
    start := request.Period.StartTime.AsTime()
    end := request.Period.EndTime.AsTime()


    score, err := s.scoreService.GetOverallQualityScores(start, end)
    if err != nil {
	log.Printf("%s", err.Error())
	return nil, err
    }

    return &pb.GetOverallQualityScoresResponse{
	Score: score,
    }, nil
}

func (s *TicketScoreServer) GetScoresByTicket(ctx context.Context, request *pb.GetScoresByTicketRequest) (*pb.GetScoresByTicketResponse, error) {
    start := request.Period.StartTime.AsTime()
    end := request.Period.EndTime.AsTime()

    response := pb.GetScoresByTicketResponse{
	Tickets: []*pb.TicketScoreEntity{},
    }
    tickets, err := s.scoreService.GetScoresByTicket(start, end)
    if err != nil {
	log.Printf("%s", err.Error())
	return nil, err
    }
    mappedTickets := mapToTickets(tickets)
    response.Tickets = mappedTickets 
    return &response, nil
} 

func (s *TicketScoreServer) GetPeriodOverPeriodScoreChange(ctx context.Context, request *pb.GetPeriodOverPeriodScoreChangeRequest) (*pb.GetPeriodOverPeriodScoreChangeResponse, error) {
    fromStart := request.PeriodFrom.StartTime.AsTime()
    fromEnd := request.PeriodFrom.EndTime.AsTime()

    toStart := request.PeriodTo.StartTime.AsTime()
    toEnd := request.PeriodTo.EndTime.AsTime()

    change, err := s.scoreService.GetPeriodOverPeriodScoreChange(fromStart, fromEnd, toStart, toEnd)
    if err != nil {
	log.Printf("%s", err.Error())
	return nil, err
    }

    return &pb.GetPeriodOverPeriodScoreChangeResponse{
	Change: int64(change),
    }, nil
} 

