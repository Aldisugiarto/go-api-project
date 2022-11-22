package cake_model

import (
	"time"
)

type CAKE struct {
	ID          int32     `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Rating      float64   `json:"rating" binding:"required"`
	Image       string    `json:"image" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
