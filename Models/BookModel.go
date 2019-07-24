package Models



type Books struct {
	Title  		string 		`json:"Title"`
	Author 		string 		`json:"Author"`
	Publisher 	string 		`json:"publisher"`
}

var BookMapping = `{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"properties":{
			"Title": {"type": "keyword"},
			"Author": {"type": "keyword"},
			"Publisher": {"type": "keyword"}
		}
	}
}`
