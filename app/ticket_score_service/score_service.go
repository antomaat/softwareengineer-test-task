package ticketscoreservice 

import (
	"time"
)

type TicketScoreService interface {
    GetAggregatedCategoryScores(startTime time.Time, endTime time.Time) ([]*ScoreEntity, error)
    GetOverallQualityScores(start time.Time, end time.Time) (int64, error )
    GetScoresByTicket(start time.Time, end time.Time) ([]*TicketScoreEntity, error)
    GetPeriodOverPeriodScoreChange(fromStart time.Time, fromEnd time.Time, toStart time.Time, toEnd time.Time) (int64, error)
}

type ScoreService struct {
    db TicketScoreRepo
}

func NewScoreService(db TicketScoreRepo) *ScoreService {
    return &ScoreService{db: db}
}

