package Controllers


import(
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"onlineLibrary/DB_connection"
	"onlineLibrary/Models"
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
		Id("DrXDImwBVvzGiteB0JPM").
		Do(context.TODO())

	if err1 != nil {
		fmt.Println("error in get request, ")
		fmt.Println(err1)
	}
	fmt.Println(get1)

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

func Put(ctx *gin.Context, x Models.Books) {
	//fmt.Printf("% T\n",x)
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
