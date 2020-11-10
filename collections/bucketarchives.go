package collections

import (
	"context"

	"github.com/textileio/textile/v2/badgerdb"
	"github.com/textileio/textile/v2/model"
	"github.com/textileio/textile/v2/mongodb"
)

type BucketArchives struct {
	hub bool
	m   mongodb.BucketArchives
	b   badgerdb.BucketArchives
}

type BucketArchiveOptions func(*BucketArchives)

func WithMongoBAOpts(m mongodb.BucketArchives) BucketArchiveOptions {
	return func(i *BucketArchives) {
		i.m = m
	}
}

func WithBadgerBAOpts(b badgerdb.BucketArchives) BucketArchiveOptions {
	return func(i *BucketArchives) {
		i.b = b
	}
}

func NewBucketArchives(_ context.Context, hub bool, opts ...BucketArchiveOptions) (*BucketArchives, error) {
	b := &BucketArchives{hub: hub}
	return b, nil
}

func (b *BucketArchives) Create(ctx context.Context, bucketKey string) (*model.BucketArchive, error) {
	if b.hub {
		return b.m.Create(ctx, bucketKey)
	} else {
		return b.b.Create(ctx, bucketKey)
	}
}

func (b *BucketArchives) Replace(ctx context.Context, ba *model.BucketArchive) error {
	if b.hub {
		return b.m.Replace(ctx, ba)
	} else {
		return b.b.Replace(ctx, ba)
	}
}

func (b *BucketArchives) GetOrCreate(ctx context.Context, bucketKey string) (*model.BucketArchive, error) {
	if b.hub {
		return b.m.GetOrCreate(ctx, bucketKey)
	} else {
		return b.b.GetOrCreate(ctx, bucketKey)
	}
}
