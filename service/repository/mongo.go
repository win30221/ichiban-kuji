package repository

import (
	"context"
	"time"

	"github.com/win30221/code/const/common"
	"github.com/win30221/core/http/catch"
	"github.com/win30221/core/storage/mongodb"
	"github.com/win30221/core/syserrno"
	"github.com/win30221/ichiban-kuji/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DB_CELLULOID_PICKET            = "celluloidPicket"
	COLLECTION_ICHIBAN_KUJI        = "ichibanKuji"
	COLLECTION_ICHIBAN_KUJI_REWARD = "ichibanKujiReward"
)

type MongoRepo struct {
	master *mongo.Client
	slave  *mongo.Client
}

func NewMongoRepo(master, slave *mongo.Client) *MongoRepo {
	return &MongoRepo{master, slave}
}

func (r *MongoRepo) GetUsersIchibanKuji(ctx context.Context, userID string) (result []domain.IchibanKuji, err error) {
	collection := r.slave.Database(DB_CELLULOID_PICKET).Collection(COLLECTION_ICHIBAN_KUJI)

	filter := bson.M{
		"userID": userID,
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		err = catch.New(syserrno.Mongo, "get ichiban kuji error", err.Error())
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		r := domain.IchibanKuji{}

		if cursor.Decode(&r) != nil {
			err = catch.New(syserrno.Mongo, "get ichiban kuji error", err.Error())
			return
		}

		result = append(result, r)
	}

	return
}

func (r *MongoRepo) CreateIchibanKuji(ctx context.Context, data domain.CreateIchibanKujiReq) (result string, err error) {

	input := domain.IchibanKuji{
		UserID:    data.UserID,
		Name:      data.Name,
		Plan:      data.Plan,
		Cost:      data.Cost,
		HashTag:   data.HashTag,
		Status:    common.ENABLE,
		SellAt:    data.SellAt,
		CreatedAt: time.Now(),
		UpatedAt:  time.Now(),
	}

	collection := r.master.Database(DB_CELLULOID_PICKET).Collection(COLLECTION_ICHIBAN_KUJI)

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		err = catch.New(syserrno.Mongo, "create ichiban kuji error", err.Error())
		return
	}

	result = res.InsertedID.(primitive.ObjectID).Hex()

	return
}

func (r *MongoRepo) GetIchibanKuji(ctx context.Context, readMaster bool, ichibanKujiID string) (result domain.IchibanKuji, err error) {

	conn := r.slave
	if readMaster {
		conn = r.master
	}
	collection := conn.Database(DB_CELLULOID_PICKET).Collection(COLLECTION_ICHIBAN_KUJI)

	filter, err := mongodb.ToObjectID(ichibanKujiID)
	if err != nil {
		err = catch.New(syserrno.Mongo, "update data error", err.Error())
		return
	}

	err = collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		err = catch.New(syserrno.Mongo, "get ichiban kuji error", err.Error())
		return
	}

	return
}

func (r *MongoRepo) UpdateIchibanKuji(ctx context.Context, ichibanKujiID string, data domain.UpdateIchibanKujiReq) (err error) {

	set, err := mongodb.ToDoc(data)
	if err != nil {
		return
	}

	filter, err := mongodb.ToObjectID(ichibanKujiID)
	if err != nil {
		err = catch.New(syserrno.Mongo, "update data error", err.Error())
		return
	}

	update := bson.M{
		"$currentDate": bson.M{
			"updatedAt": true,
		},
		"$set": set,
	}

	collection := r.master.Database(DB_CELLULOID_PICKET).Collection(COLLECTION_ICHIBAN_KUJI)

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		err = catch.New(syserrno.Mongo, "update data error", err.Error())
		return
	}

	return
}

func (r *MongoRepo) DeleteIchibanKuji(ctx context.Context, ichibanKujiID string) (err error) {

	filter, err := mongodb.ToObjectID(ichibanKujiID)
	if err != nil {
		err = catch.New(syserrno.Mongo, "update data error", err.Error())
		return
	}

	collection := r.master.Database(DB_CELLULOID_PICKET).Collection(COLLECTION_ICHIBAN_KUJI)

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		err = catch.New(syserrno.Mongo, "delete data error", err.Error())
		return
	}

	return
}

func (r *MongoRepo) CreateIchibanKujiReward(ctx context.Context, ichibanKujiID string, data domain.CreateIchibanKujiRewardReq) (err error) {

	var result []interface{}

	for i := range data.Name {
		r := domain.IchibanKujiReward{
			IchibanKujiID: ichibanKujiID,
			IsLast:        data.IsLast[i],
			UserID:        "",
			Name:          data.Name[i],
			Description:   data.Description[i],
			CreatedAt:     time.Now(),
			UpatedAt:      time.Now(),
		}
		result = append(result, r)
	}

	collection := r.master.Database(DB_CELLULOID_PICKET).Collection(COLLECTION_ICHIBAN_KUJI_REWARD)

	_, err = collection.InsertMany(ctx, result)

	if err != nil {
		err = catch.New(syserrno.Mongo, "create ichiban kuji rewards error", err.Error())
		return
	}

	return
}

func (r *MongoRepo) GetIchibanKujiRewards(ctx context.Context, readMaster bool, ichibanKujiID string) (result []domain.IchibanKujiReward, err error) {
	conn := r.slave
	if readMaster {
		conn = r.master
	}

	collection := conn.Database(DB_CELLULOID_PICKET).Collection(COLLECTION_ICHIBAN_KUJI_REWARD)

	filter := bson.M{
		"ichibanKujiID": ichibanKujiID,
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		err = catch.New(syserrno.Mongo, "get ichiban kuji rewards error", err.Error())
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		r := domain.IchibanKujiReward{}

		if cursor.Decode(&r) != nil {
			err = catch.New(syserrno.Mongo, "get ichiban kuji rewards error", err.Error())
			return
		}

		result = append(result, r)
	}

	return
}

func (r *MongoRepo) GetIchibanKujiReward(ctx context.Context, readMaster bool, ichibanKujiRewardID string) (result domain.IchibanKujiReward, err error) {

	conn := r.slave
	if readMaster {
		conn = r.master
	}
	collection := conn.Database(DB_CELLULOID_PICKET).Collection(COLLECTION_ICHIBAN_KUJI_REWARD)

	filter, err := mongodb.ToObjectID(ichibanKujiRewardID)
	if err != nil {
		err = catch.New(syserrno.Mongo, "update data error", err.Error())
		return
	}

	err = collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		err = catch.New(syserrno.Mongo, "get ichiban kuji reward error", err.Error())
		return
	}

	return
}

func (r *MongoRepo) DeleteIchibanKujiRewards(ctx context.Context, ichibanKujiID string) (err error) {

	filter := bson.M{
		"ichibanKujiID": ichibanKujiID,
	}

	collection := r.master.Database(DB_CELLULOID_PICKET).Collection(COLLECTION_ICHIBAN_KUJI_REWARD)

	_, err = collection.DeleteMany(ctx, filter)
	if err != nil {
		err = catch.New(syserrno.Mongo, "delete data error", err.Error())
		return
	}

	return
}

func (r *MongoRepo) UpdateIchibanKujiReward(ctx context.Context, ichibanKujiRewardID string, data domain.UpdateIchibanKujiRewardReq) (err error) {

	set, err := mongodb.ToDoc(data)
	if err != nil {
		return
	}

	filter, err := mongodb.ToObjectID(ichibanKujiRewardID)
	if err != nil {
		err = catch.New(syserrno.Mongo, "update data error", err.Error())
		return
	}

	update := bson.M{
		"$currentDate": bson.M{
			"updatedAt": true,
		},
		"$set": set,
	}

	collection := r.master.Database(DB_CELLULOID_PICKET).Collection(COLLECTION_ICHIBAN_KUJI_REWARD)

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		err = catch.New(syserrno.Mongo, "update data error", err.Error())
		return
	}

	return
}

func (r *MongoRepo) DeleteIchibanKujiReward(ctx context.Context, ichibanKujiRewardID string) (err error) {

	filter, err := mongodb.ToObjectID(ichibanKujiRewardID)
	if err != nil {
		err = catch.New(syserrno.Mongo, "update data error", err.Error())
		return
	}

	collection := r.master.Database(DB_CELLULOID_PICKET).Collection(COLLECTION_ICHIBAN_KUJI_REWARD)

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		err = catch.New(syserrno.Mongo, "delete data error", err.Error())
		return
	}

	return
}
