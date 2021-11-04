package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGO_URI")

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	coll := client.Database("test").Collection("newusers")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	r := gin.Default()

	r.GET("post/", func(c *gin.Context) {

		filter := bson.D{{}}
		cursor, err := coll.Find(context.TODO(), filter)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(404, err)
			}
			panic(err)
			return
		}

		var results []Blog
		if err = cursor.All(context.TODO(), &results); err != nil {
			panic(err)
		}

		c.JSON(200, results)
	})
	r.GET("/post/:email", func(c *gin.Context) {
		var result Blog
		err := coll.FindOne(context.TODO(), bson.D{{"username", c.Param("email")}}).Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(404, gin.H{
					"errorMessage": c.Param("email") + " Not found",
				})
				return
			}
		} else {
			c.JSON(200, result)
		}

	})

	r.POST("post/", func(c *gin.Context) {
		var kalifa Blog
		if err := c.BindJSON(&kalifa); err != nil {
			panic(err)
			return
		}
		result, err := coll.InsertOne(context.TODO(), kalifa)
		if err != nil {
			panic(err)
			return
		}
		fmt.Println(result)
		c.JSON(201, kalifa)
	})

	r.DELETE("/post/:email", func(c *gin.Context) {
		filter := bson.D{{"username", c.Param("email")}}
		result, err := coll.DeleteOne(context.TODO(), filter)
		if err != nil {
			panic(err)
			return
		}
		c.JSON(204, result)
	})

	r.DELETE("/post", func(c *gin.Context) {
		filter := bson.D{{}}
		result, err := coll.DeleteMany(context.TODO(), filter)
		if err != nil {
			panic(err)
			return
		}
		c.JSON(204, result)
	})
	r.Run()
}

type Blog struct {
	ID       string `json:"_id"`
	EMAIL    string `json:"email"`
	PASSWORD string `json:"password"`
	HASH     string `json:"hash"`
	SALT     string `json:"salt"`
	USERNAME string `json:"username"`
}
