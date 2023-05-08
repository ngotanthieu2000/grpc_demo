package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
)

func Connect(cloudId string, apiKey string) *elasticsearch.Client {
	fmt.Println("Connect to elastic")
	cfg := elasticsearch.Config{
		CloudID: cloudId,
		APIKey:  apiKey,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	return es
}
func CheckConnected(es *elasticsearch.Client) {
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)
}
func AddDocument(indexName string, dataDocument string, client *elasticsearch.Client) {
	// data, err := json.Marshal(dataDocument{ID: "id1",
	// 	Price:       1222,
	// 	Productname: "Apple",
	// 	Quantity:    2})
	// if err != nil {
	// 	log.Fatalf("Error endcode data: %s", err)
	// }
	req := esapi.IndexRequest{
		Index:   indexName,
		Body:    strings.NewReader(dataDocument),
		Refresh: "true",
	}
	res, errReq := req.Do(context.Background(), client)
	if errReq != nil {
		log.Fatalf("Error getting response: %s", errReq)
	}
	fmt.Println("response: %s", res)
	defer res.Body.Close()
}
func DeleteDocument(indexName string, documentId string, client *elasticsearch.Client) {
	req := esapi.DeleteRequest{
		Index:      indexName,
		DocumentID: documentId,
	}
	res, err := req.Do(context.Background(), client)
	if err != nil {
		fmt.Println("Search error: ", err)
	}
	fmt.Println("Delete result: ", res)
}
func AggregationsQuery(indexName string, body string, client *elasticsearch.Client) {
	req := esapi.SearchRequest{
		Index: []string{indexName},
		Body:  strings.NewReader(body),
	}
	res, err := req.Do(context.Background(), client)
	if err != nil {
		fmt.Println("Search error: ", err)
	}
	printAggregationsResult(res)
}
func printAggregationsResult(res *esapi.Response) error {
	defer res.Body.Close()
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		fmt.Println("Search error: ", err)
		return err
	}
	aggregations := result["aggregations"].(map[string]interface{})
	fmt.Println("Search aggregations: ", aggregations)
	return nil
}
func SearchDocument(indexName string, body string, client *elasticsearch.Client) {
	req := esapi.SearchRequest{
		Index: []string{indexName},
		Body:  strings.NewReader(body),
	}
	res, err := req.Do(context.Background(), client)
	if err != nil {
		fmt.Println("Search error: ", err)
	}
	fmt.Println("Search result: ", res)
	printResponseResult(res)
}
func CreateNewIndex(indexName string, client *elasticsearch.Client) {
	fmt.Println("Create new index: %s ", indexName)
	req := esapi.IndicesCreateRequest{
		Index: indexName,
		Body: strings.NewReader(`{
			"mappings": {
				"properties": {
				}
			}	
		}`),
	}
	res, err := req.Do(context.Background(), client)
	if err != nil {
		fmt.Println("Error getting response: ", err)
	}
	defer res.Body.Close()

	// Kiểm tra response
	if res.IsError() {
		fmt.Printf("Error: %s", res.String())
	} else {
		fmt.Println("Index created: %s", res)
	}
}
func printResponseResult(res *esapi.Response) error {
	defer res.Body.Close()
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return err
	}
	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"]
		fmt.Println(source)
	}
	return nil
}
