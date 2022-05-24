package controllers

import (
	log "log"
	models "mrkresnofatih/gobookapi/models"
	repositories "mrkresnofatih/gobookapi/repositories"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

func AddBookEndpoints(routing *gin.RouterGroup) {
	bookEndpoints := routing.Group("/books")
	{
		bookEndpoints.POST("/", saveBook)
		bookEndpoints.GET("/:id", getBookById)
	}
}

func saveBook(c *gin.Context) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad JSON Body"})
		return
	}
	book := models.NewBookBuilder().
		SetAutoID().
		SetTitle(newBook.Title).
		SetAuthor(newBook.Author).
		SetPages(newBook.Pages).
		SetCreatedAt().
		Build()
	bk, er := repositories.AddBook(*book)
	if er != nil {
		log.Println(er)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
		return
	}
	c.IndentedJSON(http.StatusOK, bk)
}

func getBookById(c *gin.Context) {
	book, err := repositories.GetBook(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}
