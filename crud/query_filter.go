package crud

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func QueryFilter() {
	uri := "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.0"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	fmt.Println("hello mongodb")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	colection := client.Database("db").Collection("shops")
	var shop []Shops
	filter := bson.M{
		"country": "United Kingdom",
		"payment_price": bson.M{
			"$gte": 500,
		},
	}
	// projection := bson.M{
	// 	"shopify_domain":  1,
	// 	"user_id":         1,
	// 	"payment_type":    1,
	// 	"created_at":      1,
	// 	"profile.country": 1,
	// }
	opts := new(options.FindOptions)
	opts.SetSkip(1)
	opts.SetLimit(1)
	cur, err := colection.Find(context.TODO(), filter, opts)
	if err != nil {
		panic(err)
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var itemShop Shops
		if err := cur.Decode(&itemShop); err != nil {
			panic(err)
		}
		shop = append(shop, itemShop)
	}
	count, _ := colection.CountDocuments(context.TODO(), filter)
	fmt.Println(shop)
	fmt.Println(count)
}

func UpdateTime() {
	uri := "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.0"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	fmt.Println("hello mongodb")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	filter := bson.M{}
	// updated := bson.M{
	// 	"$set": bson.M{
	// 		"created_at":
	// 	}
	// }
	cur, _ := client.Database("db").Collection("shops").Find(context.TODO(), filter)
	for cur.Next(context.TODO()) {
		var doc bson.M
		if err := cur.Decode(&doc); err != nil {
			continue
		}
		if createdAtStr, ok := doc["created_at"].(string); ok {
			createdAtInt64, err := strconv.ParseInt(createdAtStr, 10, 64)
			if err != nil {
				log.Printf("Error parsing created_at: %v\n", err)
				continue
			}
			doc["created_at"] = createdAtInt64
		}
		newFilter := bson.M{"_id": doc["_id"]}
		updated := bson.M{"$set": doc}
		_, err = client.Database("db").Collection("shops").UpdateOne(context.TODO(),
			newFilter, updated)
		if err != nil {
			fmt.Println("err ", err)
		}
	}
}

func GetShopM(ctx context.Context, mapValue map[string]*VarFactory, count *int) ([]Shops, error) {
	var (
		results       []Shops
		limit, offset int64
		noNeedOpt     bool
		optFind       *options.FindOptions
	)
	uri := "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.0"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	fmt.Println("hello mongodb")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	filter := bson.M{}
	if mapValue["country"] != nil {
		filter["country"] = mapValue["country"].ToString()
	}
	if mapValue["domain"] != nil {
		filter["shopify_domain"] = mapValue["domain"].ToString()
	}
	if mapValue["paid"] != nil {
		filter["paid"] = mapValue["paid"].ToString()
	}
	if mapValue["name"] != nil {
		filter["profile.name"] = mapValue["name"].ToString()
	}
	// if mapValue["domain"] != nil {
	// 	filter["shopify_domain"] = mapValue["domain"].ToString()

	// }
	filter["shopify_domain"] = "$ne:soccermondial-academynew.myshopify.com"
	if mapValue["offset"] != nil {
		offset = int64(mapValue["offset"].ToInt())
	}
	if mapValue["limit"] != nil {
		limit = int64(mapValue["limit"].ToInt())
	}
	countDoc, _ := client.Database("db").Collection("shops").CountDocuments(context.TODO(),
		filter)
	count1 := int(countDoc)
	if countDoc == 1 {
		noNeedOpt = true
	}
	if !noNeedOpt {
		optFind = options.Find()
		optFind.Skip = &offset
		optFind.Limit = &limit
	}

	cur, err := client.Database("db").Collection("shops").Find(context.TODO(),
		filter, optFind)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var item Shops
		if err := cur.Decode(&item); err != nil {
			fmt.Println("err ", err.Error())
			continue
		}
		results = append(results, item)
	}
	*count = count1
	return results, nil
}

func GetShopD(ctx context.Context, mapValue map[string]*VarFactory, count *int) ([]Shops, error) {
	var (
		results       []Shops
		limit, offset int64
		noNeedOpt     bool
		optFind       *options.FindOptions
	)
	uri := "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.0"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	fmt.Println("hello mongodb")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	filter := bson.D{}
	if mapValue["country"] != nil {
		// filter["country"] = mapValue["country"].ToString()
		filter = append(filter, primitive.E{"country", mapValue["country"].ToString()})
	}
	if mapValue["domain"] != nil {
		// filter["shopify_domain"] = mapValue["domain"].ToString()
		filter = append(filter, primitive.E{"shopify_domain", mapValue["domain"].ToString()})

	}
	if mapValue["paid"] != nil {
		// filter["paid"] = mapValue["paid"].ToString()
		filter = append(filter, primitive.E{"paid", mapValue["paid"].ToString()})

	}
	if mapValue["name"] != nil {
		// filter["profile.name"] = mapValue["name"].ToString()
		filter = append(filter, primitive.E{"profile.name", mapValue["name"].ToString()})

	}
	if mapValue["domain"] != nil {
		// filter["shopify_domain"] = mapValue["domain"].ToString()
		filter = append(filter, primitive.E{"shopify_domain", mapValue["domain"].ToString()})

	}
	if mapValue["payment_price"] != nil {
		filter = append(filter, bson.E{"payment_price", bson.D{{"$eq", mapValue["payment_price"].ToFloat()}}})
	}
	if mapValue["offset"] != nil {
		offset = int64(mapValue["offset"].ToInt())
	}
	if mapValue["limit"] != nil {
		limit = int64(mapValue["limit"].ToInt())
	}
	countDoc, _ := client.Database("db").Collection("shops").CountDocuments(context.TODO(),
		filter)
	count1 := int(countDoc)
	if countDoc == 1 {
		noNeedOpt = true
	}
	if !noNeedOpt {
		optFind = options.Find()
		optFind.Skip = &offset
		optFind.Limit = &limit
	}

	cur, err := client.Database("db").Collection("shops").Find(context.TODO(),
		filter, optFind)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var item Shops
		if err := cur.Decode(&item); err != nil {
			fmt.Println("err ", err.Error())
			continue
		}
		results = append(results, item)
	}
	*count = count1
	return results, nil
}
