package collections

import (
	"context"

	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/textile/v2/badgerdb"
	"github.com/textileio/textile/v2/model"
	"github.com/textileio/textile/v2/mongodb"
)

type Accounts struct {
	hub bool
	m   mongodb.Accounts
	b   badgerdb.Accounts
}

type AccountsOptions func(*Accounts)

func WithMongoAccountsOpts(m mongodb.Accounts) AccountsOptions {
	return func(i *Accounts) {
		i.m = m
	}
}

func WithBadgerAccountsOpts(b badgerdb.Accounts) AccountsOptions {
	return func(i *Accounts) {
		i.b = b
	}
}

func NewAccounts(_ context.Context, hub bool, opts ...AccountsOptions) (*Accounts, error) {
	a := &Accounts{
		hub: hub,
	}

	for _, opt := range opts {
		opt(a)
	}

	return a, nil
}

func (a *Accounts) UpdatePowInfo(ctx context.Context, key thread.PubKey, powInfo *model.PowInfo) (*model.Account, error) {
	if a.hub {
		return a.m.UpdatePowInfo(ctx, key, powInfo)
	} else {
		return a.b.UpdatePowInfo(ctx, key, powInfo)
	}
}

func (a *Accounts) Get(ctx context.Context, key thread.PubKey) (*model.Account, error) {
	if a.hub {
		return a.m.Get(ctx, key)

	} else {
		return a.b.Get(ctx, key)
	}
}

// func (a *Accounts) SetBucketsTotalSize(ctx context.Context, key thread.PubKey, newTotalSize int64) error {
// 	if a.hub {
// 		return a.m.SetBucketsTotalSize(ctx, key, newTotalSize)
// 	} else {
// 		return a.b.SetBucketsTotalSize(ctx, key, newTotalSize)
// 	}
// }

func (a *Accounts) GetByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*model.Account, error) {
	if a.hub {
		return a.m.GetByUsernameOrEmail(ctx, usernameOrEmail)
	} else {
		return a.b.GetByUsernameOrEmail(ctx, usernameOrEmail)
	}
}

func (a *Accounts) AddMember(ctx context.Context, username string, member model.Member) error {
	if a.hub {
		return a.m.AddMember(ctx, username, member)
	} else {
		return a.b.AddMember(ctx, username, member)
	}
}
