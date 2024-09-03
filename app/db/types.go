package db

import "time"

type Ticket struct {
    ID int64
    Subject string
}

type Rating struct {
    Id int64
    Rating int64
    Ticket_id int64
    Rating_category_id int64
    Created_at time.Time
}

type RatingCategory struct {
    Id int64
    Name string
    Weight float32
}
