package order

import (
	"time"
)

type Order struct {
	ID         string    `gorm:"primaryKey"`
	CustomerID string    `gorm:"not null"`
	ReceivedAt time.Time `gorm:"autoCreateTime"`
	Status     string
}
