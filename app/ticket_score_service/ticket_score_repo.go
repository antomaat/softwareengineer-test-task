package ticketscoreservice

import (
    "time"
    "github.com/antomaat/softwareengineering-test-task/app/db"
)

type TicketScoreRepo interface {
    GetRatingCategories() (map[int64]db.RatingCategory, error)
    GetRatingsBetweenTimeByTicket(start time.Time, end time.Time) (map[int64][]db.Rating, error)
    GetRatingsBetweenTime(start time.Time, end time.Time) ([]db.Rating, error)
    GetRatingsBetweenTimeByCategory(start time.Time, end time.Time, categories map[int64]db.RatingCategory) (map[int64][]db.Rating, error)
}
