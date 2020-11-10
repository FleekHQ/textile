package mongodb

import (
	"github.com/textileio/textile/v2/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	key      = "pow_info"
	idKey    = "id"
	tokenKey = "token"
)

func encodePowInfo(data bson.M, powInfo *model.PowInfo) {
	if powInfo != nil {
		data[key] = bson.M{
			idKey:    powInfo.ID,
			tokenKey: powInfo.Token,
		}
	}
}

func decodePowInfo(raw primitive.M) *model.PowInfo {
	var powInfo *model.PowInfo
	if v, ok := raw[key]; ok {
		powInfo = &model.PowInfo{}
		raw := v.(bson.M)
		if v, ok := raw[idKey]; ok {
			powInfo.ID = v.(string)
		}
		if v, ok := raw[tokenKey]; ok {
			powInfo.Token = v.(string)
		}
	}
	return powInfo
}
