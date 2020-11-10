package badgerdb

import (
	"context"

	"github.com/FleekHQ/space-daemon/core/store"
	"github.com/textileio/textile/v2/model"
)

type BucketArchives struct {
	st store.Store
}

func NewBucketArchives(_ context.Context, st store.Store) (*BucketArchives, error) {
	b := &BucketArchives{st: st}
	return b, nil
}

func (b *BucketArchives) Create(ctx context.Context, bucketKey string) (*model.BucketArchive, error) {
	return nil, errNotImplemented
}

func (b *BucketArchives) Replace(ctx context.Context, ba *model.BucketArchive) error {
	return errNotImplemented
}

func (b *BucketArchives) GetOrCreate(ctx context.Context, bucketKey string) (*model.BucketArchive, error) {
	return nil, errNotImplemented
}
