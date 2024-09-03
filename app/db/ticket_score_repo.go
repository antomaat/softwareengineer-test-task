package db


import (
	_ "modernc.org/sqlite"
        "time"
)

func (db *Database) GetRatingCategories() (map[int64]RatingCategory, error)  {
    query := "select id, name, weight from rating_categories order by id asc"
    rows, err := db.Conn.Query(query)
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    ratingCategories := make(map[int64]RatingCategory)

    for rows.Next() {
        category:= RatingCategory{}
        if err := rows.Scan(&category.Id, &category.Name, &category.Weight); err != nil {
            return nil, err
        }
        ratingCategories[category.Id] = category
    }
    return ratingCategories, nil 
}

func (db *Database) GetRatingsBetweenTimeByTicket(start time.Time, end time.Time) (map[int64][]Rating, error) {
    query := `select id, ticket_id, rating, rating_category_id from ratings where created_at between ? and ?`
    rows, err := db.Conn.Query(query, start.Format(time.RFC3339), end.Format(time.RFC3339))
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    ratings := make(map[int64][]Rating)

    for rows.Next() {
        rating := Rating{}
        if err := rows.Scan(&rating.Id, &rating.Ticket_id, &rating.Rating, &rating.Rating_category_id); err != nil {
            return nil, err
        }
        ratings[rating.Ticket_id] = append(ratings[rating.Ticket_id], rating)
    }
    return ratings, nil 
}

func (db *Database) GetRatingsBetweenTime(start time.Time, end time.Time) ([]Rating, error) {
    query := `select id, ticket_id, rating, rating_category_id from ratings where created_at between ? and ?`
    rows, err := db.Conn.Query(query, start.Format(time.RFC3339), end.Format(time.RFC3339))
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    ratings := []Rating{}

    for rows.Next() {
        rating := Rating{}
        if err := rows.Scan(&rating.Id, &rating.Ticket_id, &rating.Rating, &rating.Rating_category_id); err != nil {
            return nil, err
        }
        ratings = append(ratings, rating)
    }
    return ratings, nil 
}

func (db *Database) GetRatingsBetweenTimeByCategory(start time.Time, end time.Time, categories map[int64]RatingCategory) (map[int64][]Rating, error) {
    query := `select id, ticket_id, rating, rating_category_id, created_at from ratings where created_at between ? and ?`
    rows, err := db.Conn.Query(query, start.Format(time.RFC3339), end.Format(time.RFC3339))
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    ratingsByCategories := make(map[int64][]Rating)

    for rows.Next() {
        rating := Rating{}
        if err := rows.Scan(&rating.Id, &rating.Ticket_id, &rating.Rating, &rating.Rating_category_id, &rating.Created_at); err != nil {
            return nil, err
        }
        ratingsByCategories[rating.Rating_category_id] = append(ratingsByCategories[rating.Rating_category_id], rating)
    }
    return ratingsByCategories, nil 
}
