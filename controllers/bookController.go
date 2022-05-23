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
		bookEndpoints.GET("/", getAllBooks)
		bookEndpoints.GET("/:id", getBookById)
		bookEndpoints.DELETE("/:id", removeBookById)
	}
}

func saveBook(c *gin.Context) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad JSON Body"})
		return
	}
	book := models.NewBookBuilder().
		SetID().
		SetTitle(newBook.Title).
		SetAuthor(newBook.Author).
		SetPages(newBook.Pages).
		SetCreatedAt().
		Build()
	bk, er := repositories.AddBook(*book)
	if er != nil {
		log.Println(er)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book with ID already exists"})
		return
	}
	c.IndentedJSON(http.StatusOK, bk)
}

func getAllBooks(c *gin.Context) {
	books := repositories.GetAllBooks()
	c.IndentedJSON(http.StatusOK, books)
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

func removeBookById(c *gin.Context) {
	state, er := repositories.RemoveBook(c.Param("id"))
	if er != nil {
		log.Println(er)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, state)

}
