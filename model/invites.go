package model

import (
	"time"

	"github.com/textileio/go-threads/core/thread"
)

type Invite struct {
	Token     string
	Org       string
	From      thread.PubKey
	EmailTo   string
	Accepted  bool
	ExpiresAt time.Time
}
