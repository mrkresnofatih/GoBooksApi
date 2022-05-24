package models

import (
	"fmt"
	utils "mrkresnofatih/gobookapi/utils"
	"time"
)

type Book struct {
	ID        string `json:"bookId"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Pages     int    `json:"pages"`
	CreatedAt int64  `json:"createdAt"`
}

type BookBuilder struct {
	Book *Book
}

func NewBookBuilder() *BookBuilder {
	return &BookBuilder{Book: &Book{}}
}

func (b *BookBuilder) SetID(id string) *BookBuilder {
	b.Book.ID = id
	return b
}

func (b *BookBuilder) SetAutoID() *BookBuilder {
	b.Book.ID = utils.GenerateId()
	return b
}

func (b *BookBuilder) SetTitle(title string) *BookBuilder {
	b.Book.Title = title
	return b
}

func (b *BookBuilder) SetAuthor(author string) *BookBuilder {
	b.Book.Author = author
	return b
}

func (b *BookBuilder) SetPages(pages int) *BookBuilder {
	b.Book.Pages = pages
	return b
}

func (b *BookBuilder) SetCreatedAt() *BookBuilder {
	b.Book.CreatedAt = time.Now().Unix()
	return b
}

func (b *BookBuilder) Build() *Book {
	return b.Book
}

func (b *Book) PrintSelf() {
	fmt.Println(b.ID)
	fmt.Println(b.Title)
	fmt.Println(b.Author)
	fmt.Println(b.CreatedAt)
	fmt.Println(b.Pages)
}
