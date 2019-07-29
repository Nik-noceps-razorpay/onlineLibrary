package Controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/olivere/elastic.v5"
	"io/ioutil"
	"onlineLibrary/DB_connection"
	"onlineLibrary/Models"
	"strings"
	"time"

	//"os"
	//"sync"
)





func CreateBookDocument(c *gin.Context) {
	fmt.Println("Inside book controller function CreateBookDocument")

	var book []Models.Books

	err := c.Bind(&book)
	if err != nil {
		fmt.Println("error in binding")
		fmt.Println(err)
	}

	for i := 0; i < len(book); i++ {

		Put(c, book[i])

	}

}



func GetBookDocument(c *gin.Context) {
	fmt.Println("Inside GetBookDocument function")

	get1 , err1 := DB_connection.DbElastic.Get().
		Index("library").
		Type("_doc").
		Id("xP6SKGwBUUrsKZGT7O-X").
		Do(context.TODO())

	if err1 != nil {
		fmt.Println("error in get request, ")
		fmt.Println(err1)
	}
	j,err  := json.Marshal(get1.Source)
	if err != nil {
		fmt.Println("error in json marshalling")
		fmt.Println(err)
	}
	b := string(j)
	fmt.Println(b)


}



func UploadFromFile(c *gin.Context) {
	fmt.Println("Inside uploadfromfile function")


	file, errFile := c.FormFile("file")
	if errFile != nil {
		fmt.Println("error in FormFile")
		fmt.Println(errFile)
	}
	//fmt.Println(file.Filename)
	byteValue, errReadfile := ioutil.ReadFile( file.Filename)
	if errReadfile != nil {
		fmt.Println("Error in readfile ")
		fmt.Println(c.ContentType())
		fmt.Println(errReadfile)
	}

	var data []Models.Books

	json.Unmarshal(byteValue, &data)


	for i := 0; i < len(data) ; i++ {

		Put(c, data[i])

	}

}



func SearchDocs(c *gin.Context) {


	var formDate map[string]string

	err := c.ShouldBind(&formDate)

	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("formDate", formDate)

	}



	termQuery := elastic.NewTermQuery(formDate["key"], strings.ToLower(formDate["value"]))



	searchResult, err := DB_connection.DbElastic.Search().
		Index("library").
		Query(termQuery).
		Do(context.Background())

	if err != nil {
		fmt.Println("error in search query")
		//fmt.Printf("%T\n",searchResult)
		panic(err)
	}

	fmt.Println("\n Query took", searchResult.TookInMillis, "milliseconds, and found", searchResult.TotalHits(), "results")

	//fmt.Println(searchResult)

	if searchResult.TotalHits() > 0 {
		fmt.Println("Matches found are:\n ")
		for _, hit := range searchResult.Hits.Hits {

			var t Models.Books

			err := json.Unmarshal(*hit.Source, &t)
			if err!= nil {
				fmt.Println("deserialization failed , error in json unmarshal in the SearchDoc function")
				fmt.Println(err)
			}

			fmt.Println("Book Title :", t.Title, "\nAuthor:", t.Author, "\nPublisher:", t.Publisher, "\nCreatedAt:", t.CreatedAt )


		}

	} else {
		fmt.Println("No matches found ")
	}

}




func Put(ctx *gin.Context, x Models.Books) {
	//fmt.Printf("% T\n",x)

	x.CreatedAt = time.Now()

	fmt.Println(x)

	fmt.Println("Putting json entry number: ")
	put, err  := DB_connection.DbElastic.Index().
		Index("library").
		Type("_doc").
		BodyJson(x).
		Do(ctx)

	if err != nil {
		fmt.Println("error in Indexing via the put function:")
		fmt.Println("\n\nerror in data", x, "\n")
		fmt.Println(err)
	}

	fmt.Println("Indexed the book", put.Id,"in index", put.Index, "type", put.Type)

}
