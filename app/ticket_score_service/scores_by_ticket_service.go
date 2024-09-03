package ticketscoreservice

import (
        "time"
)

func (s *ScoreService) GetScoresByTicket(start time.Time, end time.Time) ([]*TicketScoreEntity, error) {
    ratings, err := s.db.GetRatingsBetweenTimeByTicket(start, end)
    if err != nil {
	return []*TicketScoreEntity{}, nil 
    }
    categories, err := s.db.GetRatingCategories()
    if err != nil {
	return nil, err
    }

    tickets := []*TicketScoreEntity{}

    for ratingKey, ratingValue := range ratings {
	ticket := TicketScoreEntity{
	    Id: ratingKey,
	    Categories: []TicketToCategoryEntity{},

	}
	for _, category := range categories {
	    score, err := CalculateTicketScoreByRating(ratingValue, categories)
	    if err != nil {
		ticket.Categories = append(ticket.Categories, TicketToCategoryEntity{Category: category.Name, Score: 0})
	    }
	    ticket.Categories = append(ticket.Categories, TicketToCategoryEntity{Category: category.Name, Score: int64(score)})
	}
	tickets = append(tickets, &ticket)
    }

    return tickets, nil
} 
