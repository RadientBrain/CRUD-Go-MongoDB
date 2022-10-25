package main

import (
	routes "Samudai/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/", routes.CreatePost)

	// called as localhost:3000/getOne/{id}
	router.GET("getOne/:postId", routes.ReadOnePost)

	// called as localhost:3000/update/{id}
	router.PUT("/update/:postId", routes.UpdatePost)

	// called as localhost:3000/delete/{id}
	router.DELETE("/delete/:postId", routes.DeletePost)

	router.Run("localhost: 3000")
}

//get by id
//update by id
//delete by id
//date datatype
//integer
// uuid with postgres paymentid and others can be string -- dynamically generated inside the database, (created at, updated ) -> dynamic -> timestamp without timezone
//docker
//postgresql
