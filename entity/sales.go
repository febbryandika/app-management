package entity

import "time"

type Sale struct {
	ID        int
	ProductID int
	Quantity  int
	SalesDate time.Time
}
