package Indexes

import (
	"context"
	"fmt"
	"onlineLibrary/DB_connection"
	"onlineLibrary/Models"
)


// make changes in mapping if Book model changes




func MakeIndex(x string) {

	ctx := context.Background()

	exists, err := DB_connection.DbElastic.IndexExists(x).Do(ctx)
	//fmt.Println("exists value is :", exists)



	if err != nil {
		// Handle error
		fmt.Println("error in index exists function")
		fmt.Println(exists)
		fmt.Println(err)

		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := DB_connection.DbElastic.CreateIndex(x).BodyString(Models.BookMapping).Do(ctx)
		fmt.Println("does not exists")
		if err != nil {
			// Handle error
			fmt.Println("error in create index: ")
			panic(err)

		}
		if !createIndex.Acknowledged {
			// Not acknowledged
			fmt.Println("Create index not acknowledged")
		}
	} else {
		fmt.Println("Index", x, "exists")
	}

	//--------------------------------- Get all index names ---------------------------------
	//names, _ := DB_connection.DbElastic.IndexNames()
	//fmt.Println("\nIndex are :")
	//for _, name := range names {
	//	fmt.Println(name)
	//}
}