package collections

import (
	"context"

	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/textile/v2/badgerdb"
	"github.com/textileio/textile/v2/model"
	"github.com/textileio/textile/v2/mongodb"
)

type Users struct {
	hub bool
	m   mongodb.Users
	b   badgerdb.Users
}

type UsersOptions func(*Users)

func WithMongoUsersOpts(m mongodb.Users) UsersOptions {
	return func(i *Users) {
		i.m = m
	}
}

func WithBadgerUsersOpts(b badgerdb.Users) UsersOptions {
	return func(i *Users) {
		i.b = b
	}
}

func NewUsers(_ context.Context, hub bool, opts ...UsersOptions) (*Users, error) {
	u := &Users{
		hub: hub,
	}
	return u, nil
}

func (u *Users) UpdatePowInfo(ctx context.Context, key thread.PubKey, powInfo *model.PowInfo) (*model.User, error) {
	if u.hub {
		return u.m.UpdatePowInfo(ctx, key, powInfo)
	} else {
		return u.b.UpdatePowInfo(ctx, key, powInfo)
	}
}

func (u *Users) Get(ctx context.Context, key thread.PubKey) (*model.User, error) {
	if u.hub {
		return u.m.Get(ctx, key)

	} else {
		return u.b.Get(ctx, key)
	}
}

func (u *Users) SetBucketsTotalSize(ctx context.Context, key thread.PubKey, newTotalSize int64) error {
	if u.hub {
		return u.m.SetBucketsTotalSize(ctx, key, newTotalSize)
	} else {
		return u.b.SetBucketsTotalSize(ctx, key, newTotalSize)
	}
}
