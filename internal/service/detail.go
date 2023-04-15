package service

import "time"

type Detail struct {
	ID          string
	Tag         string
	ServiceID   uint
	DateCreated time.Time
}
