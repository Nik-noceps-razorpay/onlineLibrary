package main

import (
	"fmt"
	"onlineLibrary/DB_connection"
	"onlineLibrary/Indexes"
	"onlineLibrary/Routes"
)

func init () {
	// context needed for using any elastic service

	// Initializing Redis Database
	// --------------------------------------------- Uncomment for redis connection -------------------------------


	//DB_connection.InitRedisDB()

	// ------------------------------------------------------------------------------------------------------------
	// Initializing Elastic Database
	DB_connection.InitElasticDB()




	// Creating an index Library in elasticsearch database

	Indexes.MakeIndex("library")

}

func main() {



	fmt.Println("In the main function")
	fmt.Println("database created")

	Routes.Router()


	// opening json file


	//defer DB_connection.DbRedis.Close()
	defer DB_connection.DbElastic.Stop()

	//err := DB_connection.Db.Do("HMSET", "book1", "Title", "The subtle art of not giving a F*ck", "Author", "Mark Manson").Err()
	//
	//if err != nil {
	//	fmt.Println("error in HMSET")
	//	fmt.Println(err)
	//}
	//
	//d := DB_connection.Db.Do("HGET","book1", "Author").Val()
	//
	//fmt.Println(d)


}

