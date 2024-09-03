package ticketscoreservice

import (
    "time"
)

func (s *ScoreService) GetPeriodOverPeriodScoreChange(fromStart time.Time, fromEnd time.Time, toStart time.Time, toEnd time.Time) (int64, error) {
    categories, err := s.db.GetRatingCategories()
    if err != nil {
	return 0, err
    }

    ratingsFrom, err := s.db.GetRatingsBetweenTime(fromStart, fromEnd)
    if err != nil {
	return 0, err
    }
    scoreFrom, err := CalculateTicketScoreByRating(ratingsFrom, categories)

    ratingsTo, err := s.db.GetRatingsBetweenTime(toStart, toEnd)
    if err != nil {
	return 0, err
    }

    scoreTo, err := CalculateTicketScoreByRating(ratingsTo, categories)

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
