package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Collection *mongo.Collection

func LoadTheEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
func CreateDBInstance() {
	// 取得資料庫位址
	connectionString := os.Getenv("DB_URI")
	// 取得資料庫的名稱
	dbName := os.Getenv("DB_NAME")
	// 取得資料表的名稱
	collectionComment := os.Getenv("DB_COLLECTION_COMMENT")

	// 連線至資料庫
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	fmt.Println(collectionComment)

	Collection = client.Database(dbName).Collection(collectionComment)
	fmt.Println("Collection instance created!")
}
