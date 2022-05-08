package connecttomongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connecttomongodb() {
	os.Setenv("MONGODB_USERNAME", "cesar")
	os.Setenv("MONGODB_PASSWORD", "cmarcial1")
	os.Setenv("MONGODB_sDBNAME", "practice")

	uri := os.ExpandEnv("mongodb+srv://${MONGODB_USERNAME}:${MONGODB_PASSWORD}@cluster0.tufuq.mongodb.net/${MONGODB_DBNAME?retryWrites=true&w=majority")
	if uri == "" {
		log.Fatalln("MONGODB_URI NEEDS to be set before starting the application")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("sample_mflix").Collection("movies")
	title := "Back to the Future"

	var result bson.M

	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}

	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
