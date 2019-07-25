package Routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"onlineLibrary/Controllers"
)

func Router() {

	router := gin.Default()

	r1 := router.Group("/CRUD")
	{
		r1.POST("/create", Controllers.CreateBookDocument)
		r1.GET("/get", Controllers.GetBookDocument)
		r1.POST("/readfile", Controllers.UploadFromFile)
		r1.POST("/searchDocs",Controllers.SearchDocs)

	}

	fmt.Println("\n\n")
	for i :=0; i < 100 ; i++ {
		fmt.Printf("**")
	}

	fmt.Println("\n    Starting Router")
	router.Run()
}
