package Controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/olivere/elastic.v5"
	"io/ioutil"
	"net/http"
	"onlineLibrary/DB_connection"
	"onlineLibrary/Models"
	"strconv"
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

	if _, exists := formDate["time1"]; exists {


		//time1, _ := time.Parse(time., formDate["time1"])
		//time2, _ := time.Parse(time.Unix(), formDate["time2"])

		time1, err1 := strconv.ParseInt(formDate["time1"], 10, 64)
		if err1 != nil {
			c.JSON(http.StatusBadRequest, "Incorrect time format in time1")
			panic(err1)
		}
		time2, err2 := strconv.ParseInt(formDate["time2"], 10, 64)
		if err2 != nil {
			c.JSON(http.StatusBadRequest, "Incorrect time format in time2")
			panic(err2)
		}

		query := elastic.NewRangeQuery("CreatedAt").
				Gte(time1).
				Lte(time2)

		size, _ := strconv.ParseInt(formDate["rpp"],10,0)
		from, _ := strconv.ParseInt(formDate["pageNo"],10,0)

		from *= size

		searchResult, err := DB_connection.DbElastic.Search().
			Index("library").
			Query(query).
			From(int(from)).
			Size(int(size)).
			Do(context.Background())

		if err != nil {
			fmt.Println("error in search query")
			//fmt.Printf("%T\n",searchResult)
			panic(err)
		}

		fmt.Printf("%T\n", searchResult)

		PrintQues(c, searchResult)




	} else {

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

		PrintQues(c,searchResult)

		//fmt.Println(searchResult)
	}


}

//func simpleSearch(c *gin.Context)



func Put(ctx *gin.Context, x Models.Books) {
	//fmt.Printf("% T\n",x)

	x.CreatedAt = time.Now().UnixNano()

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
	ctx.JSON(http.StatusOK, x)

}

func PrintQues(c *gin.Context, searchResult *elastic.SearchResult) {
	var s string

	fmt.Println("number of matches are :", searchResult.TotalHits())

	if searchResult.TotalHits() > 0 {
		fmt.Println("Matches found are:\n ")
		for _, hit := range searchResult.Hits.Hits {

			var t Models.Books

			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				fmt.Println("deserialization failed , error in json unmarshal in the SearchDoc function")
				fmt.Println(err)
			}

			fmt.Println("Book Title :", t.Title, "\nAuthor:", t.Author, "\nPublisher:", t.Publisher, "\nCreatedAt:", t.CreatedAt )
			c.JSON(http.StatusOK, t)
		}

	} else {
		s = fmt.Sprintf("No matches found ")
		c.JSON(http.StatusNotFound, s)
	}
}
