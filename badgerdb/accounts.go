package badgerdb

import (
	"context"
	"encoding/hex"
	"encoding/json"

	"github.com/FleekHQ/space-daemon/core/store"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/textile/v2/model"
)

type Accounts struct {
	st store.Store
}

const accountPubKeyPrefix = "account_pubkey_"

func NewAccounts(ctx context.Context, st store.Store) (*Accounts, error) {
	a := &Accounts{st: st}
	return a, nil
}

func (a *Accounts) UpdatePowInfo(ctx context.Context, key thread.PubKey, powInfo *model.PowInfo) (*model.Account, error) {
	return nil, errNotImplemented
}

func (a *Accounts) Get(ctx context.Context, key thread.PubKey) (*model.Account, error) {
	id, err := key.MarshalBinary()
	if err != nil {
		return nil, err
	}

	var account model.Account

	b, err := a.st.Get([]byte(accountPubKeyPrefix + hex.EncodeToString(id)))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (a *Accounts) SetBucketsTotalSize(ctx context.Context, key thread.PubKey, newTotalSize int64) error {
	id, err := key.MarshalBinary()
	if err != nil {
		return err
	}

	// Note: temporary fix for ensuring accounts don't go below zero. see #376
	if 0 > newTotalSize {
		newTotalSize = 0
	}

	var account model.Account

	// TODO: txn
	b, err := a.st.Get([]byte(accountPubKeyPrefix + hex.EncodeToString(id)))
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, account)
	if err != nil {
		return err
	}

	account.BucketsTotalSize = newTotalSize

	ub, err := json.Marshal(account)
	if err != nil {
		return err
	}

	err = a.st.Set([]byte(accountPubKeyPrefix+hex.EncodeToString(id)), ub)
	if err != nil {
		return err
	}

	return nil
}
