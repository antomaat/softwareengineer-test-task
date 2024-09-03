package ticketscoreservice


import (
    "log"
    "time"
)

func (s *ScoreService) GetOverallQualityScores(start time.Time, end time.Time) int64 {
    ratings, err := s.db.GetRatingsBetweenTime(start, end)
    if err != nil {
	log.Printf("%v", err)
    }
    categories, err := s.db.GetRatingCategories()

    score, _ := CalculateTicketScoreByRating(ratings, categories)
    return int64(score) 
}
