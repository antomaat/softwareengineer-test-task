package ticketscoreservice

import (
    "log"
    "time"
)

func (s *ScoreService) GetPeriodOverPeriodScoreChange(fromStart time.Time, fromEnd time.Time, toStart time.Time, toEnd time.Time) (int64, error) {
    categories, err := s.db.GetRatingCategories()

    ratingsFrom, err := s.db.GetRatingsBetweenTime(fromStart, fromEnd)
    if err != nil {
	log.Printf("%v", err)
    }
    scoreFrom, err := CalculateTicketScoreByRating(ratingsFrom, categories)

    ratingsTo, err := s.db.GetRatingsBetweenTime(toStart, toEnd)
    if err != nil {
	log.Printf("%v", err)
    }

    scoreTo, err := CalculateTicketScoreByRating(ratingsTo, categories)

    log.Printf("score from %f", scoreFrom)
    log.Printf("score to %f", scoreTo)

    if scoreFrom == 0 {
	if scoreTo == 0 {
	    return 0, nil
	} else {
	    return 100, nil
	}
    }

    changePercentage := ((scoreTo - scoreFrom) / scoreFrom) * 100

    return int64(changePercentage), nil
} 
