package models

import "time"

type Job struct {
	Title    string
	Location string
	Date     time.Time
	URL      string
}
