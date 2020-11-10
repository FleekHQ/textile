package model

import (
	"time"

	"github.com/textileio/go-threads/core/thread"
)

type User struct {
	Key              thread.PubKey
	BucketsTotalSize int64
	CreatedAt        time.Time
	PowInfo          *PowInfo
}
