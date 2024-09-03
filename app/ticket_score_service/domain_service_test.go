package ticketscoreservice

import (
	"testing"
	"time"

	"github.com/antomaat/softwareengineering-test-task/app/db"
	"github.com/stretchr/testify/require"
)

func TestCalculateTicketScoreByRating(t *testing.T) {

    categories := map[int64]db.RatingCategory{
	1: db.RatingCategory{ Id: 1, Name: "Spelling", Weight:  1 },
	2: db.RatingCategory{ Id: 2, Name: "Grammar", Weight:  0.7 },
	3: db.RatingCategory{ Id: 3, Name: "GDPR", Weight:  1.2 },
	4: db.RatingCategory{ Id: 4, Name: "Randomness", Weight:  0 },
	
    }

    tests := []struct {
        name     string
        ratings []int
        expected int64
    }{
        {"All Ratings at Maximum", []int{5, 5, 5, 5}, 100},
        {"Score with Rounded Value", []int{4, 3, 5, 4}, 83},
        {"Mix High and Low Ratings", []int{1, 5, 2, 3}, 47},
        {"Zero Ratings", []int{0, 0, 0, 0}, 0},
        {"Missing Ratings", []int{3, 3}, 59},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {

	    ratings := []db.Rating{}
	    for i, ratingNr := range tt.ratings {
		rating := db.Rating{ Id: int64(i + 1), Rating: int64(ratingNr), Ticket_id: 0, Rating_category_id: int64(i + 1), Created_at: time.Now()}
		ratings = append(ratings, rating)
	    }

	    percent, err := CalculateTicketScoreByRating(ratings, categories)
	    require.NoError(t, err)
	    require.Equal(t, tt.expected, int64(percent))
        })
    }
}

