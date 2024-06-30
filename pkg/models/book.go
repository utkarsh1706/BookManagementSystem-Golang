package models

import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"

    "github.com/utkarsh1706/Golang-BookManagementSystem/pkg/config"
)

var db *mongo.Client
var bookCollection *mongo.Collection

type Book struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name        string             `bson:"name" json:"name"`
    Author      string             `bson:"author" json:"author"`
    Publication string             `bson:"publication" json:"publication"`
}

func init() {
    config.Connect()
    db = config.GetDB()
    bookCollection = db.Database("bookstore").Collection("books")
}

func CreateBook(ctx context.Context, b *Book) (*mongo.InsertOneResult, error) {
    return bookCollection.InsertOne(ctx, b)
}

func GetAllBooks() ([]Book, error) {
    var books []Book
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    cursor, err := bookCollection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var book Book
        if err := cursor.Decode(&book); err != nil {
            return nil, err
        }
        books = append(books, book)
    }
    if err := cursor.Err(); err != nil {
        return nil, err
    }
    return books, nil
}

func GetBookById(Id string) (*Book, error) {
    var book Book
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    objID, err := primitive.ObjectIDFromHex(Id)
    if err != nil {
        return nil, err
    }
    err = bookCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&book)
    if err != nil {
        return nil, err
    }
    return &book, nil
}

func DeleteBook(ctx context.Context, Id string) (*mongo.DeleteResult, error) {
    objID, err := primitive.ObjectIDFromHex(Id)
    if err != nil {
        return nil, err
    }
    return bookCollection.DeleteOne(ctx, bson.M{"_id": objID})
}

func UpdateBook(ctx context.Context, Id string, book *Book) error {
    objID, err := primitive.ObjectIDFromHex(Id)
    if err != nil {
        return err
    }
    _, err = bookCollection.UpdateOne(
        ctx,
        bson.M{"_id": objID},
        bson.D{
            {Key: "$set", Value: bson.D{
                {Key: "name", Value: book.Name},
                {Key: "author", Value: book.Author},
                {Key: "publication", Value: book.Publication},
            }},
        },
    )
    return err
}
