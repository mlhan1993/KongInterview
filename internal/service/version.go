package service

import "time"

type Version struct {
	ID          string    `json:"id"`
	Tag         string    `json:"tag"`
	ServiceID   uint      `json:"serviceID"`
	DateCreated time.Time `json:"dateCreated"`
}
