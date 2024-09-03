package grpcservice


import (
	ticketscoreservice "github.com/antomaat/softwareengineering-test-task/app/ticket_score_service"
	pb "github.com/antomaat/softwareengineering-test-task/protos/ticket_score"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapToScores(scores []*ticketscoreservice.ScoreEntity) []*pb.ScoreEntity {
    pbScores := []*pb.ScoreEntity{}
    for _, score := range scores {
	pbScore := pb.ScoreEntity{
	    Category: score.Category,
	    Ratings: score.Ratings,
	    Score: score.Score,
	    ScoreDates: mapScoreDates(score.ScoreDates),
	}
	pbScores = append(pbScores, &pbScore)
    }
    return pbScores
}

func mapScoreDates(scoreDates []ticketscoreservice.ScoreDate) []*pb.ScoreDate {
    pbScoreDates := []*pb.ScoreDate{}
    for _, date := range scoreDates {
	pbDate := pb.ScoreDate{
	    Date: timestamppb.New(date.Date),
	    Score: date.Score,
	    Unit: date.Unit,
	}
	pbScoreDates = append(pbScoreDates, &pbDate)

    }
    return pbScoreDates 
    
}

func mapToTickets(tickets []*ticketscoreservice.TicketScoreEntity) []*pb.TicketScoreEntity {
    pbTickets := []*pb.TicketScoreEntity{}
    for _, ticket := range tickets {
	pbTicket := pb.TicketScoreEntity{
	    Id: ticket.Id,
	    Categories: mapToCategory(ticket.Categories),
	}
	pbTickets = append(pbTickets, &pbTicket)

    }
    return pbTickets 
}

func mapToCategory(categories []ticketscoreservice.TicketToCategoryEntity) []*pb.TickeToCategoryEntity {
    pbCategories := []*pb.TickeToCategoryEntity{}
    for _, ticket := range categories {
	pbCategory := pb.TickeToCategoryEntity{
	    Category: ticket.Category,
	    Score: ticket.Score,
	}
	pbCategories = append(pbCategories, &pbCategory)

    }
    return pbCategories 
}
