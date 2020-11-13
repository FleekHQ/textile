package collections

import (
	"context"

	"github.com/textileio/textile/v2/badgerdb"
	"github.com/textileio/textile/v2/model"
	"github.com/textileio/textile/v2/mongodb"
)

type Invites struct {
	hub bool
	m   mongodb.Invites
	b   badgerdb.Invites
}

type InvitesOptions func(*Invites)

func WithMongoInvitesOpts(m mongodb.Invites) InvitesOptions {
	return func(i *Invites) {
		i.m = m
	}
}

func WithBadgerInvitesOpts(b badgerdb.Invites) InvitesOptions {
	return func(i *Invites) {
		i.b = b
	}
}

func NewInvites(_ context.Context, hub bool, opts ...InvitesOptions) (*Invites, error) {
	k := &Invites{
		hub: hub,
	}

	for _, opt := range opts {
		opt(k)
	}

	return k, nil
}

func (i *Invites) Get(ctx context.Context, token string) (*model.Invite, error) {
	return nil, errNotImplemented
}

func (i *Invites) Accept(ctx context.Context, token string) error {
	return errNotImplemented
}

func (i *Invites) Delete(ctx context.Context, token string) error {
	return errNotImplemented
}
