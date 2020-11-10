package collections

import (
	"context"

	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/textile/v2/badgerdb"
	"github.com/textileio/textile/v2/model"
	"github.com/textileio/textile/v2/mongodb"
)

type IPNSKeys struct {
	hub bool
	m   mongodb.IPNSKeys
	b   badgerdb.IPNSKeys
}

type IPNSKeysOptions func(*IPNSKeys)

func WithMongoIPNSKeysOpts(m mongodb.IPNSKeys) IPNSKeysOptions {
	return func(i *IPNSKeys) {
		i.m = m
	}
}

func WithBadgerIPNSKeysOpts(b badgerdb.IPNSKeys) IPNSKeysOptions {
	return func(i *IPNSKeys) {
		i.b = b
	}
}

func NewIPNSKeys(_ context.Context, hub bool, opts ...IPNSKeysOptions) (*IPNSKeys, error) {
	k := &IPNSKeys{
		hub: hub,
	}
	return k, nil
}

func (k *IPNSKeys) Create(ctx context.Context, name, cid string, threadID thread.ID) error {
	if k.hub {
		return k.m.Create(ctx, name, cid, threadID)
	} else {
		return k.b.Create(ctx, name, cid, threadID)
	}
}

func (k *IPNSKeys) Get(ctx context.Context, name string) (*model.IPNSKey, error) {
	if k.hub {
		return k.m.Get(ctx, name)

	} else {
		return k.b.Get(ctx, name)
	}
}

func (k *IPNSKeys) GetByCid(ctx context.Context, cid string) (*model.IPNSKey, error) {
	if k.hub {
		return k.m.GetByCid(ctx, cid)
	} else {
		return k.b.GetByCid(ctx, cid)
	}
}

func (k *IPNSKeys) ListByThreadID(ctx context.Context, threadID thread.ID) ([]model.IPNSKey, error) {
	if k.hub {
		return k.m.ListByThreadID(ctx, threadID)
	} else {
		return k.b.ListByThreadID(ctx, threadID)

	}
}

func (k *IPNSKeys) Delete(ctx context.Context, name string) error {
	if k.hub {
		return k.m.Delete(ctx, name)
	} else {
		return k.b.Delete(ctx, name)
	}
}
