package ticketscoreservice

import "github.com/antomaat/softwareengineering-test-task/app/db"

var maxRating float32 = 5.0

func CalculateTicketScoreByRating(ratings []db.Rating, categories map[int64]db.RatingCategory) (float32, error) {
	var weightScore float32
	var weight float32

	for _, rating := range ratings {
		category, ok := categories[rating.Rating_category_id]
		if !ok {
			continue
		}

		weightedScore := float32(rating.Rating) * category.Weight
		weightScore += weightedScore
		weight += category.Weight
	}

	if weightScore == 0 {
		return 0, nil
	}

	normalizedScore := weightScore / weight
	score := (normalizedScore / maxRating) * 100

	return score, nil
}

