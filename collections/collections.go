package collections

import (
	"context"
	"time"

	db "github.com/FleekHQ/space-daemon/core/store"
	"github.com/textileio/go-threads/core/thread"
)

type IPNSKey struct {
	Name      string    `json:"_id"`
	Cid       string    `json:"cid"`
	ThreadID  thread.ID `json:"thread_id"`
	CreatedAt time.Time `json:"created_at"`
}

type IPNSKeys interface {
	Create(ctx context.Context, name, cid string, threadID thread.ID) error
	Get(ctx context.Context, name string) (*IPNSKey, error)
	GetByCid(ctx context.Context, cid string) (*IPNSKey, error)
	ListByThreadID(ctx context.Context, threadID thread.ID) ([]IPNSKey, error)
	Delete(ctx context.Context, name string) error
}

type BucketArchives interface {
}

type collection struct {
	st             db.Store
	IPNSKeys       *IPNSKeys
	BucketArchives *BucketArchives
}

type Collections interface {
}
