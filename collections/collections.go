package collections

import (
	"context"

	"github.com/textileio/textile/v2/badgerdb"
	"github.com/textileio/textile/v2/mongodb"
)

type Collections struct {
	hub        bool
	mongoname  string
	mongouri   string
	badgerpath string
	mdb        *mongodb.Collections
	bdb        *badgerdb.Collections

	//Sessions *Sessions
	Accounts *Accounts
	//Invites  *Invites

	//Threads         *Threads
	//APIKeys         *APIKeys
	IPNSKeys       *IPNSKeys
	BucketArchives *BucketArchives
	//ArchiveTracking *ArchiveTracking

	Users *Users
}

type CollectionsOptions func(*Collections)

func WithMongoCollectionOpts(uri, mongoname string) CollectionsOptions {
	return func(c *Collections) {
		c.mongoname = mongoname
		c.mongouri = uri
	}
}

func WithBadgerCollectionOpts(storepath string) CollectionsOptions {
	return func(c *Collections) {
		c.badgerpath = storepath
	}
}

func NewCollections(ctx context.Context, hub bool, opts ...CollectionsOptions) (*Collections, error) {
	c := &Collections{
		hub: hub,
	}

	for _, opt := range opts {
		opt(c)
	}

	var err error

	if c.hub {
		c.mdb, err = mongodb.NewCollections(ctx, c.mongouri, c.mongoname, c.hub)
		if err != nil {
			return nil, err
		}

		c.IPNSKeys, err = NewIPNSKeys(ctx, hub, WithMongoIPNSKeysOpts(*c.mdb.IPNSKeys))
		if err != nil {
			return nil, err
		}

		c.BucketArchives, err = NewBucketArchives(ctx, hub, WithMongoBAOpts(*c.mdb.BucketArchives))
		if err != nil {
			return nil, err
		}

		c.Accounts, err = NewAccounts(ctx, hub, WithMongoAccountsOpts(*c.mdb.Accounts))
		if err != nil {
			return nil, err
		}

		c.Users, err = NewUsers(ctx, hub, WithMongoUsersOpts(*c.mdb.Users))
		if err != nil {
			return nil, err
		}
	} else {
		c.bdb, err = badgerdb.NewCollections(ctx, c.badgerpath, c.hub)
		if err != nil {
			return nil, err
		}
		c.IPNSKeys, err = NewIPNSKeys(ctx, hub, WithBadgerIPNSKeysOpts(*c.bdb.IPNSKeys))
		if err != nil {
			return nil, err
		}

		c.BucketArchives, err = NewBucketArchives(ctx, hub, WithBadgerBAOpts(*c.bdb.BucketArchives))
		if err != nil {
			return nil, err
		}

		c.Accounts, err = NewAccounts(ctx, hub, WithBadgerAccountsOpts(*c.bdb.Accounts))
		if err != nil {
			return nil, err
		}

		c.Users, err = NewUsers(ctx, hub, WithBadgerUsersOpts(*c.bdb.Users))
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Collections) Close() error {
	err := c.bdb.Close()
	if err != nil {
		return err
	}

	err = c.mdb.Close()
	if err != nil {
		return err
	}

	return nil
}
