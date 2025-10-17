package elasticsearch

import (
	"fmt"
	"log"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	clientInstance *elasticsearch.Client
	once           sync.Once
)

// Init initializes the Elasticsearch client only once
func Init() *elasticsearch.Client {
	once.Do(func() {
		cfg := elasticsearch.Config{
			Addresses: []string{
				"http://localhost:9200",
			},
			Username: "elastic",
			Password: "changeme",
		}

		es, err := elasticsearch.NewClient(cfg)
		if err != nil {
			log.Fatalf("Error creating Elasticsearch client: %v", err)
		}

		// Check connection
		res, err := es.Info()
		if err != nil {
			log.Fatalf("Error connecting to Elasticsearch: %v", err)
		}
		defer res.Body.Close()

		fmt.Println("✅ Connected to Elasticsearch")

		clientInstance = es
	})
	return clientInstance
}

// GetClient returns the already initialized Elasticsearch client
func GetClient() *elasticsearch.Client {
	if clientInstance == nil {
		return Init()
	}
	return clientInstance
}
