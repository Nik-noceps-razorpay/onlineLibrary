package Models

import (
	"time"
)




type Books struct {
	Title     string `json:"Title"`
	Author    string `json:"Author"`
	Publisher string `json:"Publisher"`
	CreatedAt time.Time
}

var BookMapping = `{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"_doc":{
			"properties":{
				"Title": {"type": "text"},
				"Author": {"type": "text"},
				"Publisher": {"type": "keyword"},
				"CreatedAt": {
					"type": "date"
					"format": "yyyy-MM-dd HH:mm:ss"
				}
			}
		}
	}
}`
