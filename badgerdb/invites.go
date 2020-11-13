package badgerdb

import (
	"context"

	"github.com/FleekHQ/space-daemon/core/store"
	"github.com/FleekHQ/textile/v2/model"
)

type Invites struct {
	st store.Store
}

func NewInvites(ctx context.Context, st store.Store) (*Invites, error) {
	k := &Invites{st: st}
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
