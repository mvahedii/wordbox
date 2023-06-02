package dto

import "time"

type Word struct {
	ID        uint
	Title     string
	Def       string
	CreatedAt time.Time
}
