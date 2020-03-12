package schemas

import (
	"context"
	"fmt"
	"log"
	"strconv"
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

	update := bson.D{{"$set", bson.D{
		{"chatID", consumer.ChatID},
		{"lat", time.Now().Unix()},
		{"gender", consumer.Gender},
		{"topSizes", consumer.TopSizes},
		{"bottomSizes", consumer.BottomSizes},
		{"shoeSizes", consumer.ShoeSizes},
		{"styleConcepts", consumer.StyleConcepts},
		{"favoriteTypesOfClothes", consumer.FavoriteTypesOfClothes},
	}}}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	upd, err := db.ConsumersCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("SubmitConsumer:db.ConsumersCollection.UpdateOne: [%v]", err.Error())
	}

	fmt.Println("upd :", upd)

	return nil
}

func (db *DB) InitialSubmit(tgUser *TGUser) error {
	opts := options.Update().SetUpsert(true)

	filter := bson.D{{"chatID", tgUser.ChatID}}
	update := bson.D{{"$set", bson.D{
		{"user", tgUser.User},
		{"submitted", tgUser.Submitted},
		{"chatID", tgUser.ChatID},
	}}}

	// ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := db.JoinedCollection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return fmt.Errorf("InitialSubmit:db.JoinedCollection.UpdateOne: [%v]", err.Error())
	}

	return nil
}

func (db *DB) IsJoined(chatID string) bool {
	// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	id, err := strconv.Atoi(chatID)
	if err != nil {
		return false
	}
	filter := bson.D{{"chatID", int64(id)}}
	var result bson.M
	err = db.JoinedCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
	}
	if _, ok := result["_id"]; ok {
		return true
	}
	return false
}

func (db *DB) GetAllNotSubmitted() ([]string, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := db.JoinedCollection.Find(ctx, bson.D{{"submitted", false}})
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

var pipelineGroup = bson.D{{"$group", bson.D{
	{"_id", "$item._id"},
	{"results", bson.D{
		{"$addToSet", "$chatID"},
	}},
}}}

func (db *DB) FetchConsumersForPost(post *Post) ([]string, error) {
	pipeline := mongo.Pipeline{
		{{"$match", bson.D{
			{"gender", post.Gender},
		}}},
		{{"$match", bson.D{
			{"favoriteTypesOfClothes", post.TypeOfCloth},
		}}},
		{{"$match", bson.D{
			{"styleConcepts", post.StyleConcept},
		}}},
	}

	cursor, err := db.ConsumersCollection.Aggregate(context.Background(), pipeline)
	defer cursor.Close(context.Background())
	if err != nil {
		return nil, fmt.Errorf("FetchConsumersForPost:db.ConsumersCollection.Aggregate: [%v]", err.Error())
	}

	consumers := make(map[string]*Consumer)
	for cursor.Next(context.Background()) {
		// doc := &bson.D{}
		consumer := &Consumer{}
		err := cursor.Decode(consumer)
		if err != nil {
			return nil, fmt.Errorf("FetchConsumersForPost:cursor.Decode:err: [%v]", err.Error())
		}
		consumers[consumer.ChatID] = consumer
	}
	fmt.Println("consumers ", consumers)

	ids := []string{}
	switch post.TypeOfCloth {
	case Tops:
		for _, consumer := range consumers {
			for _, size := range post.TopSizes {
				if contains2(consumer.TopSizes, size) {
					ids = append(ids, consumer.ChatID)
				}
			}
		}
	case Bottoms:
		for _, consumer := range consumers {
			for _, size := range post.BottomSizes {
				if contains2(consumer.BottomSizes, size) {
					ids = append(ids, consumer.ChatID)
				}
			}
		}
	case Footwear:
		for _, consumer := range consumers {
			for _, size := range post.ShoeSizes {
				fmt.Println("consumer.ShoeSizes ", consumer.ShoeSizes)
				fmt.Println("size ", size)
				fmt.Println("---- ", contains2(consumer.ShoeSizes, size))
				if contains2(consumer.ShoeSizes, size) {
					ids = append(ids, consumer.ChatID)
				}
			}
		}
	case Accessories:
		for _, consumer := range consumers {
			ids = append(ids, consumer.ChatID)
		}
	}
	fmt.Println("ids ", ids)
	return ids, nil
}

func contains(s, s2 []int) bool {
	for _, a := range s {
		for _, a2 := range s2 {
			if a == a2 {
				return true
			}
		}
	}
	return false
}

func contains2(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
