package Controllers


import(
	"fmt"
	"github.com/gin-gonic/gin"
	"onlineLibrary/DB_connection"
	"onlineLibrary/Models"
	"sync"
)

var wg sync.WaitGroup




func CreateBookDocument(c *gin.Context) {
	fmt.Println("Inside book controller function CreateBookDocument")

	var book []Models.Books

	err := c.Bind(&book)
	if err != nil {
		fmt.Println("error in binding")
		fmt.Println(err)
	}


	wg.Add(len(book))


	for i := 0; i < len(book); i++ {

		go Put(c, book[i])

	}

	wg.Wait()

}

func Put(ctx *gin.Context, x Models.Books) {
	defer wg.Done()

	put, err  := DB_connection.DbElastic.Index().
		Index("library").
		BodyJson(x).
		Do(ctx)

	if err != nil {
		fmt.Println("error in Indexing via the put function")
		panic(err)
	}

	fmt.Println("Indexed the book", put.Id,"in index", put.Index, "type", put.Type)

}
