package schemas

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Client              *mongo.Client
	ConsumersCollection *mongo.Collection
	JoinedCollection    *mongo.Collection
	MongoURL            string

	//TODO: mongo creds
}

func (db *DB) InitDB() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(db.MongoURL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Init:mongo.Connect:err [%v]", err.Error())
	}

	db.Client = client
	db.ConsumersCollection = client.Database(DBName).Collection(ConsumersCollectionName)
	db.JoinedCollection = client.Database(DBName).Collection(JoinedCollectionName)
}

func (db *DB) SubmitConsumer(consumer *Consumer) error {

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"chatID", consumer.ChatID}}
	update := bson.D{
		{"chatID", consumer.ChatID},
		{"lat", consumer.LAT},
		{"gender", consumer.Gender},
		{"topSizes", consumer.TopSizes},
		{"bottomSizes", consumer.BottomSizes},
		{"shoeSizes", consumer.ShoeSizes},
		{"styleConcepts", consumer.StyleConcepts},
		{"favoriteTypesOfClothes", consumer.FavoriteTypesOfClothes},
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := db.ConsumersCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("SubmitConsumer:db.ConsumersCollection.UpdateOne: [%v]", err.Error())
	}
	return nil
}

func (db *DB) InitialSubmit(tgUser *TGUser) error {
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"chatID", tgUser.ChatID}}
	update := bson.D{
		{"user", tgUser.User},
		{"submitted", tgUser.Submitted},
		{"chatID", tgUser.ChatID},
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := db.JoinedCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("InitialSubmit:db.JoinedCollection.UpdateOne: [%v]", err.Error())
	}

	return nil
}

func (db *DB) UpdateSubmitted(chatID string) error {
	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{"chatID", "chatID"}}
	update := bson.D{{"$set", bson.D{{
		Key:   "submitted",
		Value: true,
	}}}}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return db.JoinedCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(nil)
}

func (db *DB) IsJoined(chatID string) bool {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.D{{"chatID", chatID}}
	err := db.JoinedCollection.FindOne(ctx, filter).Decode(nil)
	if err != nil {
		return false
	}
	return true
}

func (db *DB) GetAllNotSubmitted() ([]string, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := db.JoinedCollection.Find(ctx, bson.D{{"submitted", "false"}})
	if err != nil {
		return nil, fmt.Errorf("GetAllNotSubmitted:db.JoinedCollection.Find: [%v]", err.Error())
	}

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}

	return nil, nil
}
