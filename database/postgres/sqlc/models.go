// Code generated by sqlc. DO NOT EDIT.

package postgres

import (
	"time"
)

type Todo struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Complate  bool      `json:"complate"`
	CreatedAt time.Time `json:"createdAt"`
}
