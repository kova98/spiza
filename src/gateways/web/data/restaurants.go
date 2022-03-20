package data

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type Restaurant struct {
	Id   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name"`
	Menu Menu               `json:"menu"`
}

type Menu struct {
	Categories []string `json:"categories"`
	Items      []Item   `json:"items"`
}

type Item struct {
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Order       int32   `json:"order"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
}

func GetRestaurants() ([]Restaurant, error) {
	var restaurants []Restaurant
	var restaurant Restaurant
	l := log.New(os.Stdout, "gateways-web", log.LstdFlags)

	coll := RestaurantsCollection()
	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		l.Fatal(err)
		defer cursor.Close(context.TODO())
		return restaurants, err
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&restaurant)
		if err != nil {
			l.Fatal(err)
			return restaurants, err
		}
		restaurants = append(restaurants, restaurant)
	}

	return restaurants, nil
}

func GetRestaurant(id string) (Restaurant, error) {
	var restaurant Restaurant
	l := log.New(os.Stdout, "gateways-web", log.LstdFlags)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return restaurant, err
	}

	coll := RestaurantsCollection()
	err = coll.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&restaurant)
	if err != nil {
		l.Fatal(err)
	}
	return restaurant, err
}
