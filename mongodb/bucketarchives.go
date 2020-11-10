package mongodb

import (
	"context"

	"github.com/textileio/textile/v2/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BucketArchives struct {
	col *mongo.Collection
}

func NewBucketArchives(_ context.Context, db *mongo.Database) (*BucketArchives, error) {
	s := &BucketArchives{col: db.Collection("bucketarchives")}
	return s, nil
}

func (k *BucketArchives) Create(ctx context.Context, bucketKey string) (*model.BucketArchive, error) {
	ba := &model.BucketArchive{
		BucketKey: bucketKey,
	}
	_, err := k.col.InsertOne(ctx, ba)
	return ba, err
}

func (k *BucketArchives) Replace(ctx context.Context, ba *model.BucketArchive) error {
	res, err := k.col.ReplaceOne(ctx, bson.M{"_id": ba.BucketKey}, ba)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (k *BucketArchives) GetOrCreate(ctx context.Context, bucketKey string) (*model.BucketArchive, error) {
	res := k.col.FindOne(ctx, bson.M{"_id": bucketKey})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return k.Create(ctx, bucketKey)
		} else {
			return nil, res.Err()
		}
	}
	var raw model.BucketArchive
	if err := res.Decode(&raw); err != nil {
		return nil, err
	}
	return &raw, nil
}
