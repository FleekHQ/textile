package model

import (
	"time"

	"github.com/textileio/go-threads/core/thread"
)

type IPNSKey struct {
	Name      string    `json:"_id"`
	Cid       string    `json:"cid"`
	ThreadID  thread.ID `json:"thread_id"`
	CreatedAt time.Time `json:"created_at"`
}
