package entity

import "time"

type Activities struct {
	Id           string    `json:"id"`
	CustomerId   string    `json:"customerId"`
	Activity     string    `json:"activity"`
	ActivityTime time.Time `json:"activityTime"`
}
