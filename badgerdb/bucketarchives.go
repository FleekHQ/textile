package badgerdb

import (
	"context"
	"encoding/json"

	"github.com/FleekHQ/space-daemon/core/store"
	"github.com/dgraph-io/badger"
	"github.com/textileio/textile/v2/model"
)

type BucketArchives struct {
	st store.Store
}

const bucketKeyStoreKeyPrefix = "bucketarchives_"

func NewBucketArchives(_ context.Context, st store.Store) (*BucketArchives, error) {
	b := &BucketArchives{st: st}
	return b, nil
}

func (b *BucketArchives) Create(ctx context.Context, bucketKey string) (*model.BucketArchive, error) {
	ba := &model.BucketArchive{
		BucketKey: bucketKey,
	}

	data, err := json.Marshal(*ba)
	if err != nil {
		return nil, err
	}

	err = b.st.Set([]byte(bucketKeyStoreKeyPrefix+bucketKey), data)
	if err != nil {
		return nil, err
	}

	return ba, err
}

func (b *BucketArchives) Replace(ctx context.Context, ba *model.BucketArchive) error {
	_, err := b.st.Get([]byte(bucketKeyStoreKeyPrefix + ba.BucketKey))
	if err != nil {
		return err
	}

	// we dont do anything with res from  above
	// it was just to return an error if not found

	data, err := json.Marshal(*ba)
	if err != nil {
		return err
	}

	err = b.st.Set([]byte(bucketKeyStoreKeyPrefix+ba.BucketKey), data)
	if err != nil {
		return err
	}

	return nil
}

func (b *BucketArchives) GetOrCreate(ctx context.Context, bucketKey string) (*model.BucketArchive, error) {
	var ba *model.BucketArchive

	raw, err := b.st.Get([]byte(bucketKeyStoreKeyPrefix + bucketKey))
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return b.Create(ctx, bucketKey)
		}
		return nil, err
	}

	err = json.Unmarshal(raw, ba)
	if err != nil {
		return nil, err
	}

	return ba, nil
}
