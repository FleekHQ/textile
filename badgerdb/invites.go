package badgerdb

import (
	"context"

	"github.com/FleekHQ/space-daemon/core/store"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/textile/v2/model"
)

type Invites struct {
	st store.Store
}

func NewInvites(ctx context.Context, st store.Store) (*Invites, error) {
	k := &Invites{st: st}
	return k, nil
}

func (i *Invites) Create(ctx context.Context, from thread.PubKey, org, emailTo string) (*model.Invite, error) {
	return nil, errNotImplemented
}

func (i *Invites) Get(ctx context.Context, token string) (*model.Invite, error) {
	return nil, errNotImplemented
}

func (i *Invites) ListByEmail(ctx context.Context, email string) ([]model.Invite, error) {
	return []model.Invite{}, errNotImplemented
}

func (i *Invites) Accept(ctx context.Context, token string) error {
	return errNotImplemented
}

func (i *Invites) Delete(ctx context.Context, token string) error {
	return errNotImplemented
}

func (i *Invites) DeleteByFrom(ctx context.Context, from thread.PubKey) error {
	return errNotImplemented
}

func (i *Invites) DeleteByOrg(ctx context.Context, org string) error {
	return errNotImplemented
}

func (i *Invites) DeleteByFromAndOrg(ctx context.Context, from thread.PubKey, org string) error {
	return errNotImplemented
}
