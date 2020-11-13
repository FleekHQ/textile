package badgerdb

import (
	"context"
	"encoding/json"
	"time"

	"github.com/FleekHQ/space-daemon/core/store"
	badger "github.com/dgraph-io/badger"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/textile/v2/model"
)

type IPNSKeys struct {
	st store.Store
}

const cidStoreKeyPrefix = "cid_"
const threadIdStorePrefix = "thread_id_"
const nameStorePrefix = "name_"

func NewIPNSKeys(ctx context.Context, st store.Store) (*IPNSKeys, error) {
	k := &IPNSKeys{st: st}
	return k, nil
}

type IPNSKeySchema struct {
	Id        string    `json:"_id"`
	Cid       string    `json:"cid"`
	ThreadID  []byte    `json:"thread_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (k *IPNSKeys) Create(ctx context.Context, name, cid string, threadID thread.ID) error {
	entry := &IPNSKeySchema{
		Id:        name,
		Cid:       cid,
		ThreadID:  threadID.Bytes(),
		CreatedAt: time.Now(),
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	err = k.st.Set([]byte(cidStoreKeyPrefix+cid), data)
	if err != nil {
		return err
	}

	err = k.st.Set([]byte(threadIdStorePrefix+threadID.String()+name), data)
	if err != nil {
		return err
	}

	err = k.st.Set([]byte(nameStorePrefix+name), data)
	if err != nil {
		return err
	}

	return nil
}

func (k *IPNSKeys) Get(ctx context.Context, name string) (*model.IPNSKey, error) {
	var i IPNSKeySchema

	b, err := k.st.Get([]byte(nameStorePrefix + name))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &i)
	if err != nil {
		return nil, err
	}

	tid, err := thread.Cast(i.ThreadID)
	if err != nil {
		return nil, err
	}

	return &model.IPNSKey{
		Name:      i.Id,
		Cid:       i.Cid,
		ThreadID:  tid,
		CreatedAt: i.CreatedAt,
	}, nil
}

func (k *IPNSKeys) GetByCid(ctx context.Context, cid string) (*model.IPNSKey, error) {
	var i IPNSKeySchema

	b, err := k.st.Get([]byte(cidStoreKeyPrefix + cid))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &i)
	if err != nil {
		return nil, err
	}

	tid, err := thread.Cast(i.ThreadID)
	if err != nil {
		return nil, err
	}

	return &model.IPNSKey{
		Name:      i.Id,
		Cid:       i.Cid,
		ThreadID:  tid,
		CreatedAt: i.CreatedAt,
	}, nil
}

func (k *IPNSKeys) ListByThreadID(ctx context.Context, threadID thread.ID) ([]model.IPNSKey, error) {
	keys, err := k.st.KeysWithPrefix(threadIdStorePrefix + threadID.String())
	if err != nil {
		return nil, err
	}

	var docs []model.IPNSKey

	for _, key := range keys {
		b, err := k.st.Get([]byte(key))
		// skipping ErrKeyNotFound in case it was deleted after the
		// keyswithprefix call above
		if err != nil && err != badger.ErrKeyNotFound {
			return nil, err
		}

		var i IPNSKeySchema
		err = json.Unmarshal(b, i)
		if err != nil {
			return nil, err
		}

		tid, err := thread.Cast(i.ThreadID)
		if err != nil {
			return nil, err
		}

		ipnsKey := &model.IPNSKey{
			Name:      i.Id,
			Cid:       i.Cid,
			ThreadID:  tid,
			CreatedAt: i.CreatedAt,
		}

		docs = append(docs, *ipnsKey)
	}

	return docs, nil
}

func (k *IPNSKeys) Delete(ctx context.Context, name string) error {
	var i IPNSKeySchema

	b, err := k.st.Get([]byte(nameStorePrefix + name))
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &i)
	if err != nil {
		return err
	}

	err = k.st.Remove([]byte(nameStorePrefix + name))
	if err != nil {
		return err
	}

	err = k.st.Remove([]byte(cidStoreKeyPrefix + i.Cid))
	if err != nil {
		return err
	}

	err = k.st.Remove([]byte(threadIdStorePrefix + string(i.ThreadID)))
	if err != nil {
		return err
	}

	return nil
}
