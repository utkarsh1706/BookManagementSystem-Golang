package config

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func Connect() {
    clientOptions := options.Client().ApplyURI("mongodb+srv://utkarsh:sonamutkarsh@cluster0.fjbuelq.mongodb.net/{0}?retryWrites=true&w=majority")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    if err := client.Ping(ctx, nil); err != nil {
        log.Fatal(err)
    }

    db = client
}

func GetDB() *mongo.Client {
    return db
}
