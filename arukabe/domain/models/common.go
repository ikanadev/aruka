package models

import "time"

type PaginationData struct {
	Page       uint32
	Limit      uint32
	TotalItems uint32
	TotalPages uint32
	HasNext    bool
	HasPrev    bool
}

type TimeData struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ArchivedAt time.Time
	DeletedAt time.Time
}
