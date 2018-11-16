package main

import (
	"fmt"

	"github.com/olivere/elastic"
	"github.com/prometheus/common/log"
	"context"
	"encoding/json"
)

type SupervisorTweet struct {
	Timestamp string `json:"@timestamp"`
	Hostname  string `json:"hostname"`
	Level     string `json:"level"`
	Line      string `json:"line"`
	Name      string `json:"name"`
	State     string `json:"state"`
	Status    string `json:"status"`
}

func CreateElasticSearchClient(url string) (*elastic.Client, error) {
	ctx := context.Background()

	c, err := elastic.NewSimpleClient(elastic.SetURL(url))
	if err != nil {
		log.Errorf("Create Elastic client error: %s NewSimpleClient(%s)", err, url)
		return nil, err
	}

	info, code, err := c.Ping(url).Do(ctx)
	if err != nil {
		log.Errorf("Elastic client can not ping %s, error: %s", url, err)
		return nil, err
	}
	log.Infof("Elastic returned with code %d and version %s", code, info.Version.Number)
	return c, nil
}

func SearchData(client *elastic.Client) {
	termQuery := elastic.NewQueryStringQuery("name:alertmanager")

	searchResult, err := client.Search().Query(termQuery).Do(context.Background())

	if err != nil {
		log.Errorf("Elastic query term error: %s", err)
	}

	if searchResult.Hits.TotalHits > 0 {
		log.Infof("Total: %d", searchResult.Hits.TotalHits)

		for _, hit := range searchResult.Hits.Hits {
			var t SupervisorTweet
			err := json.Unmarshal(*hit.Source, &t)

			if err != nil {
				log.Errorln("Deserialization failed")
			}
			log.Infof("time: %s, hostname: %s, name: %s, status: %s", t.Timestamp, t.Hostname, t.Name, t.Status)
		}
	} else {
		log.Errorln("Found no tweets")
	}
}

func main() {
	fmt.Println("Elastic Demo...")

	url := "http://192.168.1.78:9200"
	c, err := CreateElasticSearchClient(url)
	if err != nil {
		panic("Create ElasticSearch client failed!")
	}

	SearchData(c)
}
