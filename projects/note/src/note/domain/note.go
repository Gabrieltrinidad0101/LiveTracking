package domain

import (
	"time"
)

type CurrentStatus int

const (
	Pending CurrentStatus = iota
	InProgress
	Completed
	Failed
)

type Note struct {
	Name          string
	Description   string
	UserName      string
	UserId        string
	StartTime     time.Time
	EndTime       time.Time
	CurrentStatus CurrentStatus
}
