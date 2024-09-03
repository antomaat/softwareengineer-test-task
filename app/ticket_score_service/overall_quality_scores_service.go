package ticketscoreservice


import (
    "time"
)

func (s *ScoreService) GetOverallQualityScores(start time.Time, end time.Time) (int64, error) {
    ratings, err := s.db.GetRatingsBetweenTime(start, end)
    if err != nil {
	return 0, err
    }
    categories, err := s.db.GetRatingCategories()

    score, _ := CalculateTicketScoreByRating(ratings, categories)
    return int64(score), nil
}
