package mongodb

//https://blog.logrocket.com/how-to-use-mongodb-with-go/
import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect() (bool, error) {

	//get the context
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	//get the client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)

	}
	//ping database or get collection instance
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	} else {
		fmt.Println("Connected")
		usersCollection := client.Database("catalina").Collection("catalina")

		user := bson.D{{"fullName", "User 1"}, {"age", 30}}
		// insert the bson object using InsertOne()
		result, err := usersCollection.InsertOne(context.TODO(), user)
		// check for errors in the insertion
		if err != nil {
			panic(err)
		}
		// display the id of the newly inserted object
		fmt.Println(result.InsertedID)
	}
	return true, nil
}
