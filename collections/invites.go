package collections

import (
	"context"

	"github.com/textileio/go-threads/core/thread"
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

func (i *Invites) Create(ctx context.Context, from thread.PubKey, org, emailTo string) (*model.Invite, error) {
	if i.hub {
		return i.m.Create(ctx, from, org, emailTo)
	} else {
		return i.b.Create(ctx, from, org, emailTo)
	}
}

func (i *Invites) Get(ctx context.Context, token string) (*model.Invite, error) {
	if i.hub {
		return i.m.Get(ctx, token)
	} else {
		return i.b.Get(ctx, token)
	}
}

func (i *Invites) ListByEmail(ctx context.Context, email string) ([]model.Invite, error) {
	if i.hub {
		return i.m.ListByEmail(ctx, email)
	} else {
		return i.b.ListByEmail(ctx, email)
	}
}

func (i *Invites) Accept(ctx context.Context, token string) error {
	if i.hub {
		return i.m.Accept(ctx, token)
	} else {
		return i.b.Accept(ctx, token)
	}
}

func (i *Invites) Delete(ctx context.Context, token string) error {
	if i.hub {
		return i.m.Delete(ctx, token)
	} else {
		return i.b.Delete(ctx, token)
	}
}

func (i *Invites) DeleteByFrom(ctx context.Context, from thread.PubKey) error {
	if i.hub {
		return i.m.DeleteByFrom(ctx, from)
	} else {
		return i.b.DeleteByFrom(ctx, from)
	}
}

func (i *Invites) DeleteByOrg(ctx context.Context, org string) error {
	if i.hub {
		return i.m.DeleteByOrg(ctx, org)
	} else {
		return i.b.DeleteByOrg(ctx, org)
	}
}

func (i *Invites) DeleteByFromAndOrg(ctx context.Context, from thread.PubKey, org string) error {
	if i.hub {
		return i.m.DeleteByFromAndOrg(ctx, from, org)
	} else {
		return i.b.DeleteByFromAndOrg(ctx, from, org)
	}
}
