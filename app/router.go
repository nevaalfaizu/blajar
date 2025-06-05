package app

import (
	"tes1/controller/book"
	"tes1/controller/inventory"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	// r.GET("/hello", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello1, World!",
	// 	})
	// })

	// r.GET("/hello", hello)

	// r.GET("/login", login)
	// r.POST("/login", login)

	//Router for books
	r.GET("/books", book.GetBooks)
	r.POST("/books", book.PostBook)
	r.PUT("/books/:id", book.UpdateBook)
	r.DELETE("/books/:id", book.DeleteBook)
	r.GET("/books/:id", book.GetBookByID)

	//router for inventory
	r.GET("/inventorys", inventory.GetInventory)
	r.POST("/inventorys", inventory.PostInventory)
	r.PUT("/inventorys/:id", inventory.UpdateInventory)
	r.DELETE("/inventorys/:id", inventory.DeleteInventory)
	r.GET("/inventorys/:id", inventory.GetInventoryByID)
}

// func hello(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "Hello, World!",
// 	})
// }

// func login(c *gin.Context) {
// 	c.Params.Get("username")
// 	c.Params.Get("password")
// 	c.JSON(200, gin.H{
// 		"message":  "Login successful",
// 		"username": c.Query("username"),
// 		"password": c.Query("password"),
// 	})
// }
