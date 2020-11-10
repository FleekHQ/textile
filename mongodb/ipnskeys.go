package mongodb

import (
	"context"
	"time"

	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/textile/v2/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPNSKeys struct {
	col *mongo.Collection
}

func NewIPNSKeys(ctx context.Context, db *mongo.Database) (*IPNSKeys, error) {
	k := &IPNSKeys{col: db.Collection("ipnskeys")}
	_, err := k.col.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{{"cid", 1}},
		},
		{
			Keys: bson.D{{"thread_id", 1}},
		},
	})
	return k, err
}

func (k *IPNSKeys) Create(ctx context.Context, name, cid string, threadID thread.ID) error {
	_, err := k.col.InsertOne(ctx, bson.M{
		"_id":        name,
		"cid":        cid,
		"thread_id":  threadID.Bytes(),
		"created_at": time.Now(),
	})
	return err
}

func (k *IPNSKeys) Get(ctx context.Context, name string) (*model.IPNSKey, error) {
	res := k.col.FindOne(ctx, bson.M{"_id": name})
	if res.Err() != nil {
		return nil, res.Err()
	}
	var raw bson.M
	if err := res.Decode(&raw); err != nil {
		return nil, err
	}
	return decodeIPNSKey(raw)
}

func (k *IPNSKeys) GetByCid(ctx context.Context, cid string) (*model.IPNSKey, error) {
	res := k.col.FindOne(ctx, bson.M{"cid": cid})
	if res.Err() != nil {
		return nil, res.Err()
	}
	var raw bson.M
	if err := res.Decode(&raw); err != nil {
		return nil, err
	}
	return decodeIPNSKey(raw)
}

func (k *IPNSKeys) ListByThreadID(ctx context.Context, threadID thread.ID) ([]model.IPNSKey, error) {
	cursor, err := k.col.Find(ctx, bson.M{"thread_id": threadID.Bytes()})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var docs []model.IPNSKey
	for cursor.Next(ctx) {
		var raw bson.M
		if err := cursor.Decode(&raw); err != nil {
			return nil, err
		}
		doc, err := decodeIPNSKey(raw)
		if err != nil {
			return nil, err
		}
		docs = append(docs, *doc)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return docs, nil
}

func (k *IPNSKeys) Delete(ctx context.Context, name string) error {
	res, err := k.col.DeleteOne(ctx, bson.M{"_id": name})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func decodeIPNSKey(raw bson.M) (*model.IPNSKey, error) {
	threadID, err := thread.Cast(raw["thread_id"].(primitive.Binary).Data)
	if err != nil {
		return nil, err
	}
	var created time.Time
	if v, ok := raw["created_at"]; ok {
		created = v.(primitive.DateTime).Time()
	}
	return &model.IPNSKey{
		Name:      raw["_id"].(string),
		Cid:       raw["cid"].(string),
		ThreadID:  threadID,
		CreatedAt: created,
	}, nil
}
