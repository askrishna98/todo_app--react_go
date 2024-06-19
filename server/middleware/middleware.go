package middleware

// For handling various HTTP request and database interaction

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aswin/go-react-todo/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// collection is a reference to the MongoDB collection
var collection *mongo.Collection

// init for to load environment variables  and create a database intance
func init() {
	loadTheEnv()
	createDBInstance()
}

// loadTheEnv loads environment variables from the .env file.
func loadTheEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("ERROR LAODING the .env file")
	}
}

// createDBInstance creates a MongoDB client instance and connects to the database.
func createDBInstance() {
	// gets connectionString,DBname,Collection name from environment variable
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION_NAME")

	clientOption := options.Client().ApplyURI(connectionString)
	// Connect to MongoDB.
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	// To verify the connection Pings the MongoDB
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongo db")

	//DBinstance created - now we can query to this collection
	collection = client.Database(dbName).Collection(collName)
	fmt.Println("connection instance craeted")

}

// GetAllTasks handles the HTTP request for retrieving all tasks.
func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// function 'getAllTasks' Retrieve all tasks from the database.
	data := getAllTasks()

	// // Encode the data as JSON and write it to the response
	json.NewEncoder(w).Encode(data)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// task is an instance of TodoList struct we created before and holds fields [ID,Task,Status]
	var task models.ToDoList

	// Decode the request body into the task
	json.NewDecoder(r.Body).Decode(&task)

	// If description of task empty we return
	if task.Task == "" {
		return
	}

	// 'insertOneTask' function will add task to database
	insertOneTask(task)

	// Encode the task as JSON and write it to the response
	json.NewEncoder(w).Encode(task)
}

func TaskComplete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Get task ID from URL parameters using mux.Vars(r)
	params := mux.Vars(r)

	// 'taskComplete' function makes status of task to 'True'
	taskComplete(params["id"])

	// Encode the response as JSON and send it in the response body
	json.NewEncoder(w).Encode(params["id"])

}

func UndoTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Get task ID from URL parameters using mux.Vars(r)
	params := mux.Vars(r)

	// 'undoTask' makes the status of a task to False
	undoTask(params["id"])

	json.NewEncoder(w).Encode(params["id"])

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Get task ID from URL parameters using mux.Vars(r)
	params := mux.Vars(r)

	// 'deleteOneTask' function deletes the selected task from the collection
	deleteOneTask(params["id"])

}

func DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Call deleteAllTasks function to delete all tasks and get the count of deleted documents
	count := deleteAllTasks()
	json.NewEncoder(w).Encode(count)

}

// insertOneTask inserts a single task into the MongoDB collection
func insertOneTask(task models.ToDoList) {

	// Insert the task into the collection and capture the result and any errors
	insertresult, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted single record", insertresult)
}

// getAllTasks retrieves all tasks from the MongoDB collection.
func getAllTasks() []primitive.M {

	// bson.D{{}} is the query to get all the task from collection
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	// initializes empty slice
	var results []primitive.M

	// Iterate through the cur to decode each document and append it to results
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	// Close the cur
	cur.Close(context.Background())

	// Return the slice of BSON documents (tasks).
	return results
}

// finds the task from MongoDB and set the status field to True
func taskComplete(task string) {
	// Convert task ID string to MongoDB ObjectID
	id, _ := primitive.ObjectIDFromHex(task)

	// creates filter to find object by its ID
	filter := bson.M{"_id": id}

	// Define update to set the status field to true
	update := bson.M{"$set": bson.M{"status": true}}

	// finds result or error
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count,", result)
}

// finds the task from MongoDB and set the status field to false
func undoTask(task string) {
	// Convert task ID string to MongoDB ObjectID
	id, _ := primitive.ObjectIDFromHex(task)

	// Creates a filter to find the task by its ID
	filter := bson.M{"_id": id}

	// Define update to set the status field to false
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count : ", result.ModifiedCount)
}

func deleteOneTask(task string) {
	// Convert task ID string to MongoDB ObjectID
	id, _ := primitive.ObjectIDFromHex(task)

	// Creates a filter to find the task by its ID
	filter := bson.M{"_id": id}

	// deletes task from Collection
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted documnet", d.DeletedCount, d, filter)
}

func deleteAllTasks() int64 {

	// Perform delete operation to delete all documents from the collection
	d, err := collection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted all document", d.DeletedCount)
	return d.DeletedCount
}
