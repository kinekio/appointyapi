package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}

type Post struct {
	Id              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption         string             `json:"caption" bson:"caption"`
	ImageURL        string             `json:"imageurl" bson:"imageurl"`
	PostedTimestamp string             `json:"postedtimestamp" bson:"postedtimestamp"`
}

var client *mongo.Client

func createuser(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		response.Header().Add("content-type", "application/json")
		var user User
		json.NewDecoder(request.Body).Decode(&user)
		collection := client.Database("instagram").Collection("users")
		ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
		result, _ := collection.InsertOne(ctx, user)
		json.NewEncoder(response).Encode(result)
	}
}

func getuser(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		response.Header().Add("content-type", "application/json")
		geturl := request.URL
		path := geturl.Path
		parts := strings.Split(path, "/")
		param := parts[1]
		id, _ := primitive.ObjectIDFromHex(param)
		var user User
		collection := client.Database("instagram").Collection("users")
		ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
		result := collection.FindOne(ctx, User{Id: id}).Decode(&user)
		json.NewEncoder(response).Encode(result)
	}

}
func createpost(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		response.Header().Add("content-type", "application/json")
		var post Post
		json.NewDecoder(request.Body).Decode(&post)
		collection := client.Database("instagram").Collection("posts")
		ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
		result, _ := collection.InsertOne(ctx, post)
		json.NewEncoder(response).Encode(result)
	}
}

func getpost(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		response.Header().Add("content-type", "application/json")
		geturl := request.URL
		path := geturl.Path
		parts := strings.Split(path, "/")
		param := parts[1]
		id, _ := primitive.ObjectIDFromHex(param)
		var post Post
		collection := client.Database("instagram").Collection("posts")
		ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
		result := collection.FindOne(ctx, Post{Id: id}).Decode(&post)
		json.NewEncoder(response).Encode(result)
	}

}
func getallpost(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		response.Header().Add("content-type", "application/json")
		var post []Post
		json.NewDecoder(request.Body).Decode(&post)
		collection := client.Database("instagram").Collection("posts")
		ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
		result, _ := collection.Find(ctx, bson.M{})
		json.NewEncoder(response).Encode(result)
		defer result.Close(ctx)
		for result.Next(ctx) {
			var postdummy Post
			result.Decode(&postdummy)
			post = append(post, postdummy)
		}
	}

}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(client)
	router := http.NewServeMux()
	router.HandleFunc("/users", createuser)
	router.HandleFunc("/users/{id}", getuser)
	router.HandleFunc("/posts", createpost)
	router.HandleFunc("/posts/{id}", getpost)
	router.HandleFunc("/posts/users/{id}", getallpost)

	log.Fatal(http.ListenAndServe(":9000", router))
}
