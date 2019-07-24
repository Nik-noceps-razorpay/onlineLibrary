package Routes

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"onlineLibrary/Controllers"
)

func Router() {

	router := gin.Default()

	r1 := router.Group("/CRUD")
	{
		r1.POST("/create", Controllers.CreateBookDocument)

	}

	router.Run()
}
