package db

import (
	"context"
	"fmt"

	"github.com/Ath3r/hotel-backend/internal/config"
	"github.com/Ath3r/hotel-backend/internal/types"
	"github.com/Ath3r/hotel-backend/pkg/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserStore interface {
	GetById(ctx context.Context, id string) (*types.User, error)
	GetAll(ctx context.Context) ([]*types.User, error)
	Create(ctx context.Context, user *types.User) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client, ) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll: client.Database(config.AppConfig.DatabaseName).Collection("users"),
	}
}

func ConnectMongo() (*mongo.Client, error) {
	credential := options.Credential{
		Username: config.AppConfig.DatabaseUser,
		Password: config.AppConfig.DatabasePassword,
	}
	connString := fmt.Sprintf(
		"mongodb://%v:%v/%v",
		config.AppConfig.DatabaseHost,
		config.AppConfig.DatabasePort,
		config.AppConfig.DatabaseName,
	)
	fmt.Println(connString)
	clientOptions := options.Client().ApplyURI(connString).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB")
	return client, nil
}

func (s *MongoUserStore) GetById(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	objId, err := helpers.ToObjectId(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}
	cur := s.coll.FindOne(ctx, bson.M{"_id": objId})
	if err := cur.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *MongoUserStore) GetAll(ctx context.Context) ([]*types.User, error) {
	cur, err := s.coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var users []*types.User
	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (s *MongoUserStore) Create(ctx context.Context, user *types.User) (*types.User, error) {
	_, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

