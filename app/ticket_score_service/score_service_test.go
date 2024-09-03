package ticketscoreservice_test

import (
	"testing"
	"time"

	"github.com/antomaat/softwareengineering-test-task/app/db"
	ticketscoreservice "github.com/antomaat/softwareengineering-test-task/app/ticket_score_service"
	"github.com/stretchr/testify/require"
)

var categories = map[int64]db.RatingCategory{
    1: db.RatingCategory{ Id: 1, Name: "Spelling", Weight:  1 },
    2: db.RatingCategory{ Id: 2, Name: "Grammar", Weight:  0.7 },
    3: db.RatingCategory{ Id: 3, Name: "GDPR", Weight:  1.2 },
    4: db.RatingCategory{ Id: 4, Name: "Randomness", Weight:  0 },
}

func TestGetOverallQualityScoresService(t *testing.T) {

    mockRepo := &MockTicketScoreRepo{
	CategoriesStub: func()(map[int64]db.RatingCategory, error){
	    return categories, nil
	},
	RatingsBetweenTimeStub: func(start time.Time, end time.Time) ([]db.Rating, error) {
	    return []db.Rating{
		{ Id: int64(1), Rating: 4, Ticket_id: 1, Rating_category_id: 1, Created_at: time.Now().AddDate(0, 0, -1)},
		{ Id: int64(2), Rating: 3, Ticket_id: 1, Rating_category_id: 2, Created_at: time.Now().AddDate(0, 0, -2)},
		{ Id: int64(3), Rating: 5, Ticket_id: 1, Rating_category_id: 3, Created_at: time.Now().AddDate(0, 0, -3)},
		{ Id: int64(4), Rating: 4, Ticket_id: 1, Rating_category_id: 4, Created_at: time.Now().AddDate(0, 0, -4)},
	    }, nil
	},
    }

    service := ticketscoreservice.NewScoreService(mockRepo)

    startTime := time.Now().AddDate(0, 0, -7)
    endTime := time.Now()

    score, err := service.GetOverallQualityScores(startTime, endTime)
    require.NoError(t, err)
    require.Equal(t, int64(83), score)

}

func TestGetScoresByTicketService(t *testing.T) {

    mockRepo := &MockTicketScoreRepo{
	CategoriesStub: func()(map[int64]db.RatingCategory, error){
	    return categories, nil
	},
	RatingsByTicketStub: func(start time.Time, end time.Time) (map[int64][]db.Rating, error) {
	return map[int64][]db.Rating{ 
		1 : []db.Rating{
		    { Id: int64(1), Rating: 4, Ticket_id: 1, Rating_category_id: 1, Created_at: time.Now().AddDate(0, 0, -1)},
		    { Id: int64(2), Rating: 3, Ticket_id: 1, Rating_category_id: 2, Created_at: time.Now().AddDate(0, 0, -2)},
		    { Id: int64(3), Rating: 5, Ticket_id: 1, Rating_category_id: 3, Created_at: time.Now().AddDate(0, 0, -3)},
		    { Id: int64(4), Rating: 4, Ticket_id: 1, Rating_category_id: 4, Created_at: time.Now().AddDate(0, 0, -4)},
		},
		2 : []db.Rating{
		    { Id: int64(1), Rating: 5, Ticket_id: 2, Rating_category_id: 1, Created_at: time.Now().AddDate(0, 0, -1)},
		    { Id: int64(2), Rating: 5, Ticket_id: 2, Rating_category_id: 2, Created_at: time.Now().AddDate(0, 0, -2)},
		    { Id: int64(3), Rating: 5, Ticket_id: 2, Rating_category_id: 3, Created_at: time.Now().AddDate(0, 0, -3)},
		    { Id: int64(4), Rating: 5, Ticket_id: 2, Rating_category_id: 4, Created_at: time.Now().AddDate(0, 0, -4)},
		},
	}, nil
	},
    }

    service := ticketscoreservice.NewScoreService(mockRepo)

    startTime := time.Now().AddDate(0, 0, -7)
    endTime := time.Now()

    tickets, err := service.GetScoresByTicket(startTime, endTime)

    require.NoError(t, err)
    require.Equal(t, 2, len(tickets))
    require.Equal(t, int64(83), tickets[0].Categories[0].Score)
    require.Equal(t, int64(100), tickets[1].Categories[0].Score)
}

func TestGetPeriodOverPeriodScoreService(t *testing.T) {
    mockRepo := &MockTicketScoreRepo{
	CategoriesStub: func()(map[int64]db.RatingCategory, error){
	    return categories, nil
	},
	RatingsBetweenTimeStub: func(start time.Time, end time.Time) ([]db.Rating, error) {

	    if (start.Before(time.Now().AddDate(0, 0, -8))) {
		return []db.Rating{
		    { Id: int64(1), Rating: 3, Ticket_id: 2, Rating_category_id: 1, Created_at: time.Now().AddDate(0, 0, -14)},
		    { Id: int64(2), Rating: 3, Ticket_id: 2, Rating_category_id: 2, Created_at: time.Now().AddDate(0, 0, -15)},
		    { Id: int64(3), Rating: 4, Ticket_id: 2, Rating_category_id: 3, Created_at: time.Now().AddDate(0, 0, -16)},
		}, nil
	    }

	    return []db.Rating{
		{ Id: int64(1), Rating: 4, Ticket_id: 1, Rating_category_id: 1, Created_at: time.Now().AddDate(0, 0, -1)},
		{ Id: int64(2), Rating: 3, Ticket_id: 1, Rating_category_id: 2, Created_at: time.Now().AddDate(0, 0, -2)},
		{ Id: int64(3), Rating: 5, Ticket_id: 1, Rating_category_id: 3, Created_at: time.Now().AddDate(0, 0, -3)},
	    }, nil
	},
    }

    service := ticketscoreservice.NewScoreService(mockRepo)

    fromStart := time.Now().AddDate(0, 0, -14)
    fromEnd := time.Now().AddDate(0, 0, -7)

    toStart := time.Now().AddDate(0, 0, -7)
    toEnd := time.Now() 

    score, err := service.GetPeriodOverPeriodScoreChange(fromStart, fromEnd, toStart, toEnd)
    require.NoError(t, err)
    require.Equal(t, int64(22), score)
}

type MockTicketScoreRepo struct {
    CategoriesStub func() (map[int64]db.RatingCategory, error)
    RatingsByTicketStub func(start time.Time, end time.Time) (map[int64][]db.Rating, error)
    RatingsBetweenTimeStub func(start time.Time, end time.Time) ([]db.Rating, error)
    RatingsByCategoryStub func(start time.Time, end time.Time, categories map[int64]db.RatingCategory) (map[int64][]db.Rating, error)
}

func (m *MockTicketScoreRepo) GetRatingCategories() (map[int64]db.RatingCategory, error) {
    if m.CategoriesStub != nil {
	    return m.CategoriesStub()
    }
    return nil, nil
}

func (m *MockTicketScoreRepo) GetRatingsBetweenTimeByTicket(start time.Time, end time.Time) (map[int64][]db.Rating, error) {
    if m.RatingsByTicketStub != nil {
	    return m.RatingsByTicketStub(start, end)
    }
    return nil, nil
}

func (m *MockTicketScoreRepo) GetRatingsBetweenTime(start time.Time, end time.Time) ([]db.Rating, error) {
    if m.RatingsBetweenTimeStub != nil {
	    return m.RatingsBetweenTimeStub(start, end)
    }
    return nil, nil
}

func (m *MockTicketScoreRepo) GetRatingsBetweenTimeByCategory(start time.Time, end time.Time, categories map[int64]db.RatingCategory) (map[int64][]db.Rating, error) {
    if m.RatingsByCategoryStub != nil {
	    return m.RatingsByCategoryStub(start, end, categories)
    }
    return nil, nil
}

