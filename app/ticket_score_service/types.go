package ticketscoreservice

import "time"

type TicketScoreEntity struct {
    Id int64
    Categories []TicketToCategoryEntity 
}

type TicketToCategoryEntity struct {
    Category string
    Score int64
}

type ScoreEntity struct {
    Category string
    Ratings int64
    Score int64
    ScoreDates []ScoreDate
}

type ScoreDate struct {
    Date time.Time
    Unit string
    Score int64
}
