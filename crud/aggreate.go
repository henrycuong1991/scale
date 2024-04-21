package crud

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MgAggreate() {

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
	cur, err := colection.Aggregate(context.TODO(), mongo.Pipeline{})
	if err != nil {
		panic(err)
	}
	// display the results
	var results []bson.M
	if err = cur.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Printf("Average price of %v tea options: $%v \n", result["_id"], result["average_price"])
		fmt.Printf("Number of %v tea options: %v \n\n", result["_id"], result["type_total"])
	}
	//build match stags
	matchState := bson.D{
		{"$match", bson.D{{"toppings", "milk foam"}}}}
	unsetState := bson.D{{"$unset", bson.A{"_id", "category"}}}
	sortStage := bson.D{{"$sort", bson.D{{"price", 1}, {"toppings", 1}}}}
	limitState := bson.D{{"$limit", 2}}
	cursor, err := colection.Aggregate(context.TODO(),
		mongo.Pipeline{
			matchState,
			unsetState,
			sortStage,
			limitState})
	if err != nil {
		panic(err)
	}
	var returnValues []Tea
	if err := cursor.All(context.TODO(), &returnValues); err != nil {
		panic(err)

	}
	for _, result := range returnValues {
		fmt.Printf("Tea: %v \nToppings: %v \nPrice: $%v \n\n", result.Type, strings.Join(result.Toppings, ", "), result.Price)
	}

}

func AggregateTest(ctx context.Context, mapValue map[string]*VarFactory) ([]Shops, error) {
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
	matchState := bson.D{{"$match", bson.D{{"country", mapValue["country"].ToString()}}}}
	groupState := bson.D{
		{"$group", bson.D{
			{"_id", "null"},
			{"count", bson.D{
				{"$sum", 1},
			}},
		}},
	}
	skipState := bson.D{{"$skip", mapValue["offset"].ToInt()}}
	limitState := bson.D{{"$limit", mapValue["limit"].ToInt()}}
	var results []Shops
	cur, err := colection.Aggregate(ctx, mongo.Pipeline{
		matchState, skipState, limitState,
	}, nil)
	// var item bson.M
	// err = cur.Decode(&item)
	// if err != nil {
	// 	fmt.Println("err", err.Error())
	// }

	for cur.Next(ctx) {
		var shopItem Shops
		err = cur.Decode(&shopItem)
		if err != nil {
			fmt.Println("err", err.Error())
			continue
		}
		results = append(results, shopItem)
	}
	if err := cur.Err(); err != nil {
		fmt.Println("err", err.Error())
	}
	fmt.Println("item len ", len(results))
	// Đảm bảo tài nguyên được giải phóng
	cur.Close(ctx)
	curCount, err := colection.Aggregate(ctx, mongo.Pipeline{
		matchState,
		groupState,
	}, nil)
	for curCount.Next(context.TODO()) {
		var item bson.M
		if err := curCount.Decode(&item); err != nil {
			fmt.Println("err ", err.Error())
		}
		fmt.Println("count now ", item["count"])
	}
	defer curCount.Close(ctx)
	return results, nil
}

func MyMongoDbList(ctx context.Context, mapValue map[string]*VarFactory) ([]Shops, interface{}, error) {
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
	//build match state
	var pipiline, pipilineCount []bson.M
	matchState := bson.M{
		"$match": bson.M{
			"country": mapValue["country"].ToString(),
		},
	}
	groupState := bson.M{
		"$group": bson.M{
			"_id": "null",
			"count": bson.M{
				"$sum": 1,
			},
		},
	}
	skipState := bson.M{
		"$skip": mapValue["offset"].ToInt() * mapValue["limit"].ToInt(),
	}
	limitState := bson.M{
		"$limit": mapValue["limit"].ToInt(),
	}
	pipiline = append(pipiline, matchState, skipState, limitState)
	pipilineCount = append(pipilineCount, matchState, groupState)
	//get value
	curValue, err := colection.Aggregate(ctx, pipiline, nil)
	if err != nil {
		return nil, nil, err
	}
	var result []Shops
	defer curValue.Close(ctx)
	for curValue.Next(ctx) {
		var shopItem Shops
		if err := curValue.Decode(&shopItem); err != nil {
			return nil, nil, err
		}
		result = append(result, shopItem)
	}
	//count
	countValue, err := colection.Aggregate(ctx, pipilineCount, nil)
	if err != nil {
		return nil, nil, err
	}
	var count interface{}
	for countValue.Next(ctx) {
		var item bson.M
		if err := countValue.Decode(&item); err != nil {
			return nil, nil, err
		}
		count = item["count"]
	}
	fmt.Println("count ", count)
	return result, count, nil
}

func ChatGPTAgg() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Kết nối đến MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Chọn cơ sở dữ liệu và bảng
	collection := client.Database("db").Collection("shops")

	// Tạo một pipeline aggregation
	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{"country": "United Kingdom"}, // Match documents with category equals "electronics"
		},
		bson.M{
			"$group": bson.M{
				"_id":   nil,
				"count": bson.M{"$sum": 1}, // Count total documents
			},
		},
		bson.M{
			"$skip": 8, // Skip number of documents (change to your desired value)
		},
		bson.M{
			"$limit": 10, // Limit the number of documents returned (change to your desired value)
		},
	}

	// Thực hiện aggregation
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err)
	}

	// Đọc kết quả từ con trỏ
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		count := result["count"]
		fmt.Println("Total count:", count)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
}

// use addFields / push in aggegrate
func AggegrateWithFunc(ctx context.Context, mapValue map[string]*VarFactory) (interface{}, error) {

	type ResultStruct struct {
		ID      string   `bson:"_id" json:"_id"`
		Count   int      `bson:"count" json:"count"`
		Domains []string `bson:"domains" json:"domains"`
	}

	var results []bson.M
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
	//get all domain in special conditions
	var pipeline bson.D
	// if mapValue["country"] != nil {
	// 	pipeline = append(pipeline, bson.M{
	// 		"$match": bson.M{"country": mapValue["country"].ToString()},
	// 	})
	// }

	// if mapValue["domain"] != nil {
	// 	pipeline = append(pipeline, bson.M{
	// 		"$match": bson.M{"shopify_domain": mapValue["domain"].ToString()},
	// 	})
	// }
	//group
	pipeline = bson.D{
		{"$match",
			bson.D{
				{"payment_date",
					bson.D{
						{"$ne", ""},
						{"$ne", primitive.Null{}},
					},
				},
			},
		},
	}
	// groupState := bson.D{
	// 	{"$group",
	// 		bson.D{
	// 			{"_id",
	// 				bson.D{
	// 					{"$dateToString",
	// 						bson.D{
	// 							{"format", "%Y-%m-%d"},
	// 							{"date",
	// 								bson.D{
	// 									{"$dateFromString",
	// 										bson.D{
	// 											{"dateString",
	// 												bson.D{
	// 													{"$substr",
	// 														bson.A{"$payment_date", 0, 10},
	// 													},
	// 												},
	// 											},
	// 											{"timezone", "UTC"}, // Assuming the date is in UTC
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},

	// 			{"count",
	// 				bson.D{
	// 					{"$sum", 1},
	// 				},
	// 			},
	// 			{"domains",
	// 				bson.D{
	// 					{"$push", "$shopify_domain"},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	groupState2 := bson.D{
		{"$group", bson.D{
			{"_id", bson.D{
				{"$dateToString", bson.D{
					{"date", bson.D{
						{"$dateFromString", bson.D{
							{"dateString", bson.D{
								{"$substr", bson.A{"$payment_date", 0, 10}},
							}},
						}},
					}},
					{"format", "%Y-%m-%d"},
				}},
			}},
			{"count",
				bson.D{
					{"$sum", 1},
				},
			},
			{"domains",
				bson.D{
					{"$push", "$shopify_domain"},
				},
			},
		},
		},
	}
	skipState := bson.D{
		{"$skip", mapValue["offset"].ToInt()},
	}
	limitState := bson.D{
		{"$limit", mapValue["limit"].ToInt()},
	}
	sortState := bson.D{
		{"$sort", bson.D{
			{"_id", 1},
		}},
	}
	// pipeline = append(pipeline, bson.M{
	// 	"$addFields": bson.M{
	// 		"total_money": "$sum_value",
	// 	},
	// })
	// pipeline = append(pipeline, bson.M{
	// 	"$skip": mapValue["offset"].ToInt(),
	// })
	// pipeline = append(pipeline, bson.M{
	// 	"$limit": mapValue["limit"].ToInt(),
	// })
	cur, err := colection.Aggregate(context.TODO(), mongo.Pipeline{pipeline, groupState2, skipState, limitState, sortState}, nil)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var item bson.M
		// result := ResultStruct{}
		if err := cur.Decode(&item); err != nil {
			fmt.Println("err ", cur.Current.String())
			return nil, err
		}
		// result.ID = item["_id"].(string)
		// result.Total = item["sum"].(int32)
		// result.NameCountry = item["name_country"].(string)
		fmt.Println("value ", item["_id"], reflect.TypeOf(item["domains"]))
		// ref
		// fmt.Println("value ", item["sum"], item["_id"], item["name_country"])
		results = append(results, item)
	}
	return results, nil
}
