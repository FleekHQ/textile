package badgerdb

import (
	"context"
	"encoding/hex"
	"encoding/json"

	"github.com/FleekHQ/space-daemon/core/store"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/textile/v2/model"
)

type Users struct {
	st store.Store
}

const usersPubKeyPrefix = "user_pubkey_"

func NewUsers(ctx context.Context, st store.Store) (*Users, error) {
	u := &Users{st: st}
	return u, nil
}

func (u *Users) UpdatePowInfo(ctx context.Context, key thread.PubKey, powInfo *model.PowInfo) (*model.User, error) {
	return nil, errNotImplemented
}

func (u *Users) Get(ctx context.Context, key thread.PubKey) (*model.User, error) {
	id, err := key.MarshalBinary()
	if err != nil {
		return nil, err
	}

	var user model.User

	b, err := u.st.Get([]byte(usersPubKeyPrefix + hex.EncodeToString(id)))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *Users) SetBucketsTotalSize(ctx context.Context, key thread.PubKey, newTotalSize int64) error {
	id, err := key.MarshalBinary()
	if err != nil {
		return err
	}

	// Note: temporary fix for ensuring accounts don't go below zero. see #376
	if 0 > newTotalSize {
		newTotalSize = 0
	}

	var user model.User

	// TODO: txn
	b, err := u.st.Get([]byte(usersPubKeyPrefix + hex.EncodeToString(id)))
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, user)
	if err != nil {
		return err
	}

	user.BucketsTotalSize = newTotalSize

	ub, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = u.st.Set([]byte(usersPubKeyPrefix+hex.EncodeToString(id)), ub)
	if err != nil {
		return err
	}

	return nil
}
