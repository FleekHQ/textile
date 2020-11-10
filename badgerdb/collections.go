package badgerdb

import (
	"context"
	"errors"
	"time"

	"github.com/FleekHQ/space-daemon/core/store"
)

const (
	tokenLen = 44

	DuplicateErrMsg = "E11000 duplicate key error"
)

var errNotImplemented = errors.New("Not implemented")

type ctxKey string

type Collections struct {
	st             store.Store
	IPNSKeys       *IPNSKeys
	BucketArchives *BucketArchives
	Accounts       *Accounts
	Users          *Users
}

// NewCollections gets or create store instances for active collections.
func NewCollections(ctx context.Context, storePath string, hub bool) (*Collections, error) {
	st := store.New(
		store.WithPath(storePath),
	)
	err := st.Open()
	if err != nil {
		return nil, err
	}

	c := &Collections{st: st}

	c.IPNSKeys, err = NewIPNSKeys(ctx, st)
	if err != nil {
		return nil, err
	}
	c.BucketArchives, err = NewBucketArchives(ctx, st)
	if err != nil {
		return nil, err
	}
	c.Accounts, err = NewAccounts(ctx, st)
	if err != nil {
		return nil, err
	}
	c.Users, err = NewUsers(ctx, st)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Collections) Close() error {
	_, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return c.st.Close()
}
