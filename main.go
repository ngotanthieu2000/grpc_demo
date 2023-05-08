package main

// "fmt"
// "log"
// "github.com/elastic/go-elasticsearch/v8"
import (
	"elasticsearch/example/mongo"
	"math/rand"
	"strings"
	"time"
)

type Order struct {
	OrderName string
	Quantity  int32
	TotalCost int32
}

// random text
var sentences = []string{
	"The quick brown fox jumps over the lazy dog.",
	"She sells seashells by the seashore.",
	"Peter Piper picked a peck of pickled peppers.",
	"How much wood would a woodchuck chuck if a woodchuck could chuck wood?",
	"I scream, you scream, we all scream for ice cream!",
}

func randomSentence() string {
	rand.Seed(time.Now().UnixNano())
	return sentences[rand.Intn(len(sentences))]
}

func randomString(n int) string {
	var result []string

	for i := 0; i < n; i++ {
		result = append(result, randomSentence())
	}

	return strings.Join(result, " ")
}

func main() {
	// fmt.Println("Heelo Elasticsearch")
	// elasticCloudId := "elastic_golang_intern:dXMtY2VudHJhbDEuZ2NwLmNsb3VkLmVzLmlvOjQ0MyRmMmRkNmRiNDJhNjE0MTcwODIwZDZiNjM1OWU5NDA4MiQzYzc2OWE2M2JiZTk0NDgyYjVjYWI5M2Q1OWM3NGY2OQ=="
	// elasticApiKey := "OE5hQnhZY0IwTXl3SG5HNEw4WjM6S2RkVERJVVJUbkc0SC1qeDdJcTBiUQ=="
	// data := Order{
	// 	OrderName: "Oppo",
	// 	Quantity:  4,
	// 	TotalCost: 9999,
	// }
	// jsonBytes, err := json.Marshal(data)
	// if err != nil {
	// 	// xử lý lỗi
	// }
	// jsonString := string(jsonBytes)
	// fmt.Println("jsonString:", jsonString)
	// connectES := elastic.Connect(elasticCloudId, elasticApiKey)
	// elastic.CheckConnected(connectES)
	// elastic.AddDocument("orders", jsonString, connectES)
	// SearchDocument(indexName string, searchValue string, client *elasticsearch.Client)
	// body := `{
	// 	"query": {
	// 	  "query_string": {
	// 		"query": "(new york city) OR (big apple)",
	// 		"default_field": "OrderName"
	// 	  }
	// 	}
	//   }`
	// body :=
	// 	`{
	// 		"aggs": {
	// 		  "avg_grade": { "avg": { "field": "TotalCost" } }
	// 		}
	// 	  }`
	// elastic.SearchDocument("orders", body, connectES)
	// elastic.AggregationsQuery("orders", body, connectES)
	// DeleteDocument(indexName string, documentId string, client *elasticsearch.Client)
	// elastic.DeleteDocument("orders", "i6Z35YcB1xNvbEd4BeAv", connectES)
	// log.Println(es.Info())

	// mongo.GetProduct("644b8d25e68fe558a81a93c0")
	// mongo.UpdateProduct("644b92492d5f09ed3d58444e", myProduct)
	mongo.CreateIndex("mydb", "product")
}
