package ticketscoreservice

import (
	"log"
	"time"

	"github.com/antomaat/softwareengineering-test-task/app/db"
)

func (s *ScoreService) GetAggregatedCategoryScores(startTime time.Time, endTime time.Time) ([]*ScoreEntity, error) {

    categories, err := s.db.GetRatingCategories()
    ratingsByCategories, err := s.db.GetRatingsBetweenTimeByCategory(startTime, endTime, categories)
    if err != nil {
	log.Printf("%v", err)
    }

    isUnitInWeeks := getRatingsByUnit(startTime, endTime)
    
    scores := []*ScoreEntity{}

    for category_id, ratings := range ratingsByCategories {
	category := categories[category_id]

	scoreEntity := ScoreEntity{}
	scoreEntity.Category = category.Name 

	if len(ratings) > 0 {
	    scoreEntity.Score = calculateTotalScoreOfCategory(ratings, categories)

	    ratingsByDay, timeUnit := mapRatingsByUnit(ratings, isUnitInWeeks) 

	    for singleDayRatingKey, singleDayRatingValue := range ratingsByDay {
		singleDayScore, _ := CalculateTicketScoreByRating(singleDayRatingValue, categories)
		scoreDate := ScoreDate{Date: singleDayRatingKey, Score: int64(singleDayScore), Unit: timeUnit }
		scoreEntity.ScoreDates = append(scoreEntity.ScoreDates, scoreDate)
	    } 

	}
	scores = append(scores, &scoreEntity)
    }


    return scores, nil
}

func mapRatingsByUnit(ratings []db.Rating, isUnitInWeeks bool) (map[time.Time][]db.Rating, string)  {
    if isUnitInWeeks {
	return mapRatingsByFunc(ratings, getStartOfWeek), "weeks"
    }	
    return mapRatingsByFunc(ratings, getStartOfDay), "days"
}

func getRatingsByUnit(start time.Time, end time.Time) bool {
    return end.AddDate(0, -1, 0).Before(start)
}

func calculateTotalScoreOfCategory(ratings []db.Rating, categories map[int64]db.RatingCategory) int64 {
    score, err := CalculateTicketScoreByRating(ratings, categories)
    if err != nil {
	return 0
    }
    return int64(score) 
}

func mapRatingsByFunc(ratings []db.Rating, categorisingFunc func(time.Time) time.Time) map[time.Time][]db.Rating {
    ratingsByDay := make(map[time.Time][]db.Rating)
    
    for _, rating := range ratings {
	ratingsByDay[categorisingFunc(rating.Created_at)] = append(ratingsByDay[categorisingFunc(rating.Created_at)], rating)
    }

    return ratingsByDay
}

func getStartOfDay(t time.Time) time.Time {
    return t.Truncate(24 * time.Hour)
}


func getStartOfWeek(t time.Time) time.Time {
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 6
	} else {
		weekday--
	}
	return t.AddDate(0, 0, -weekday).Truncate(24 * time.Hour)
}

