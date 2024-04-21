package crud

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"

	"context"
)

func ElkQuery(ctx context.Context, mapValue map[string]*VarFactory) (interface{}, int, error) {
	esClient, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, 0, err
	}
	var (
		response map[string]interface{}
		// buf      map[string]interface{}
	)
	// param := make(map[string]interface{})
	// for k, v := range mapValue {
	// 	if k == "key" {
	// 		param[]
	// 	}
	// }
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				mapValue["key"].ToString(): mapValue["value"].ToString(),
			},
		},
	}
	body, err := json.Marshal(query)
	if err != nil {
		return nil, 0, err
	}
	res, err := esClient.Search(
		esClient.Search.WithContext(ctx),
		esClient.Search.WithIndex("shop_idx"),
		// esClient.Search.WithQuery(query),
		esClient.Search.WithBody(strings.NewReader(string(body))),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithPretty(),
		esClient.Search.WithSource("shopify_domain", "lname"),
	)

	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()
	// Check the response status
	if res.IsError() {
		log.Fatalf("Error response: %s", res.Status())
	}

	// Decode the response
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	// Print the search results
	fmt.Println(response)
	return &response, 0, nil
}

func ElkOrgQuery(ctx context.Context, mapValue map[string]*VarFactory) (interface{}, int, error) {
	esClient, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, 0, err
	}
	var (
		response map[string]interface{}
		// buf      map[string]interface{}
	)
	// param := make(map[string]interface{})
	// for k, v := range mapValue {
	// 	if k == "key" {
	// 		param[]
	// 	}
	// }
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				mapValue["key"].ToString(): mapValue["value"].ToString(),
			},
		},
	}
	body, err := json.Marshal(query)
	if err != nil {
		return nil, 0, err
	}
	res, err := esClient.Search(
		esClient.Search.WithContext(ctx),
		esClient.Search.WithIndex("new_org_index"),
		// esClient.Search.WithQuery(query),
		esClient.Search.WithBody(strings.NewReader(string(body))),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithPretty(),
		// esClient.Search.WithSource("shopify_domain", "lname"),
	)

	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()
	// Check the response status
	if res.IsError() {
		log.Fatalf("Error response: %s", res.Status())
	}

	// Decode the response
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	// Print the search results
	fmt.Println(response)
	return &response, 0, nil
}
