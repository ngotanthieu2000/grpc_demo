package mongo

import (
	"context"
	pb "elasticsearch/example/grpc"
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

type Product struct {
	Name     string  `bson:"name"`
	Price    float64 `bson:"price"`
	Quantity int32   `bson:"quantity"`
}
type Connect struct {
}
type Query struct {
}
type Database interface {
	Connect() interface{}
	Query() interface{}
}
type MongoDB struct {
}

func (mg *MongoDB) Connect() interface{} {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	return client
}
func (mg *MongoDB) Query() interface{} {
	return Query{}
}

type Elastic struct {
}

func (el *Elastic) Connect() interface{} {
	return Connect{}
}
func (el *Elastic) Query() interface{} {
	return Query{}
}

// DatabaseFactory("mongodb").Connect()
func DatabaseFactory(dbType string) (Database, error) {
	switch dbType {
	case "mongodb":
		return &MongoDB{}, nil
	case "elastic":
		return &Elastic{}, nil
	default:
		return nil, fmt.Errorf("Invalid database type")
	}
}
func connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	return client
}

func CreateProduct(p *pb.Product) {
	client := connect()
	collection = client.Database("mydb").Collection("product")
	res, err := collection.InsertOne(ctx, p)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)
	fmt.Println("Inserted a single document", res.InsertedID)
}

func GetProduct(idP string) *pb.Product {
	client := connect()
	collection = client.Database("mydb").Collection("product")
	defer client.Disconnect(ctx)
	id, _ := primitive.ObjectIDFromHex(idP)
	filter := bson.D{{"_id", id}}
	fmt.Println("Filter data db:", filter)
	var result pb.Product
	fmt.Println("result empty:", result)

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return &result
}
func CreateIndex(databaseName string, collectionName string) {
	client := connect()
	defer client.Disconnect(ctx)
	db := client.Database(databaseName)
	collection = db.Collection(collectionName)
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"productname": 1, // 1 cho index tăng dần, -1 cho index giảm dần
		},
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		panic(err)
	}
}
func UpdateProduct(idP string, updateProduct *pb.Product) {
	// connected db
	client := connect()
	collection = client.Database("mydb").Collection("product")
	// create ObjectID by string id
	id, _ := primitive.ObjectIDFromHex(idP)
	// set fillter
	filter := bson.D{{"_id", id}}
	fmt.Println("Filter : ", filter)
	// update data by filter and update data

	update := bson.D{}

	val := reflect.ValueOf(updateProduct)
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		key := field.Tag.Get("bson")
		if key == "" {
			key = field.Name
		}
		update = append(update, bson.E{Key: key, Value: val.Field(i).Interface()})
	}
	fmt.Println("update : ", update)

	result, err := collection.UpdateOne(ctx, filter, bson.D{{"$set", update}})
	if err != nil {
		panic(err)
	}
	// disconnect db
	fmt.Println(result)
	defer client.Disconnect(ctx)

}
