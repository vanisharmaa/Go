package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/vanisharmaa/go/25.mongo-api/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

const connectionString string = "mongodb+srv://vanisharma:Vani123*@cluster0.mh9etap.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName string = "netflix"
const colName string = "watchlist"

// !IMPORTANT
var collection *mongo.Collection

// connect with mongo db
func init() { // init method only runs the first time the application loads
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(clientOption)

	if err != nil {
		log.Fatal(err)
	}

	// context basically is used so that a func doesn't take forever
	// here 10s would be given to anything that uses ctx context, after that, context will be cancelled
	// context.Background creates an empty context, context.WithTimeout sets a timer to it
	// ! We always need to pass a context to any db operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // context.TODO - use when you dont know whether to use context with timeout, deadline or cancel. will let the context run without auto cancellations
	defer cancel()

	// Verify connection using Ping
	// ctx is passed here, meaning - if mongo doesn't respond in 10s, context will be cancelled
	// knowing that context is now cancelled, mongo will throw an error
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Cannot connect to MongoDB:", err)
	}
	fmt.Println("Mongo connection successful")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("colllection instance is ready")
}

// mongo db helper funcs - file

// insert one record
func insertOneMovie(movie model.Netflix) { // model -> package model

	inserted, err := collection.InsertOne(context.Background(), movie) // context.Background() without deadline/timeout/cancel works same as context.TODO()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("One Record Inserted - ", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	_id, err := bson.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{"watched": true}}

	/**

	filter := bson.E{Key: "_id", Value: _id}
	update := bson.D{bson.E{Key: "watched", Value: true}}

	*/

	res, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", res.MatchedCount)

}

func deleteOneMovie(movieId string) {
	id, err := bson.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	res, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted one movie - ", res.DeletedCount)
}

func deleteAllMovies() {
	filter := bson.M{} // bson.D{{}}
	res, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All movies deleted - ", res.DeletedCount)
}

func getAllMovies() []bson.M {
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var movies []bson.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	return movies
}

// Actual controller file

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func InsertAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	fmt.Println("inserted one movie - ", movie.ID)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode("marked " + params["id"] + " as watched")
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"] + " deleted from DB!")
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	deleteAllMovies()
	json.NewEncoder(w).Encode("Deleted all movies from DB!")
}
