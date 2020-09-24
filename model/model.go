package model

import "time"

type Area struct {
	Id        int    `gorm:primary_key`
	Name      string `gorm:primary_key`
	ParentId  int
	Level     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
