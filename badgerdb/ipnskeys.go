package badgerdb

import (
	"context"
	"encoding/json"
	"time"

	"github.com/FleekHQ/space-daemon/core/store"
	c "github.com/FleekHQ/textile/collections"
	badger "github.com/dgraph-io/badger"
	"github.com/textileio/go-threads/core/thread"
)

type IPNSKey struct {
	Name      string    `json:"_id"`
	Cid       string    `json:"cid"`
	ThreadID  thread.ID `json:"thread_id"`
	CreatedAt time.Time `json:"created_at"`
}

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
	id        string    `json:"_id"`
	cid       string    `json:"cid"`
	threadId  []byte    `json:"thread_id"`
	createdAt time.Time `json:"created_at"`
}

func (k *IPNSKeys) Create(ctx context.Context, name, cid string, threadID thread.ID) error {
	entry := &IPNSKeySchema{
		id:        name,
		cid:       cid,
		threadId:  threadID.Bytes(),
		createdAt: time.Now(),
	}

	data, err := json.Marshal(*entry)
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

func (k *IPNSKeys) Get(ctx context.Context, name string) (*c.IPNSKey, error) {
	var ipnsKey c.IPNSKey

	b, err := k.st.Get([]byte(nameStorePrefix + name))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, ipnsKey)
	if err != nil {
		return nil, err
	}

	return &ipnsKey, nil
}

func (k *IPNSKeys) GetByCid(ctx context.Context, cid string) (*c.IPNSKey, error) {
	var ipnsKey c.IPNSKey

	b, err := k.st.Get([]byte(cidStoreKeyPrefix + cid))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, ipnsKey)
	if err != nil {
		return nil, err
	}

	return &ipnsKey, nil
}

func (k *IPNSKeys) ListByThreadID(ctx context.Context, threadID thread.ID) ([]IPNSKey, error) {
	keys, err := k.st.KeysWithPrefix(threadIdStorePrefix + threadID.String())
	if err != nil {
		return nil, err
	}

	var docs []IPNSKey

	for _, key := range keys {
		b, err := k.st.Get([]byte(key))
		// skipping ErrKeyNotFound in case it was deleted after the
		// keyswithprefix call above
		if err != nil && err != badger.ErrKeyNotFound {
			return nil, err
		}

		var ipnsKey IPNSKey
		err = json.Unmarshal(b, ipnsKey)
		if err != nil {
			return nil, err
		}

		docs = append(docs, ipnsKey)
	}

	return docs, nil
}

func (k *IPNSKeys) Delete(ctx context.Context, name string) error {
	var ipnsKey IPNSKey

	b, err := k.st.Get([]byte(nameStorePrefix + name))
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, ipnsKey)
	if err != nil {
		return err
	}

	err = k.st.Remove([]byte(nameStorePrefix + name))
	if err != nil {
		return err
	}

	err = k.st.Remove([]byte(cidStoreKeyPrefix + ipnsKey.Cid))
	if err != nil {
		return err
	}

	err = k.st.Remove([]byte(threadIdStorePrefix + ipnsKey.ThreadID.String()))
	if err != nil {
		return err
	}

	return nil
}
