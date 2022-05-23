package repositories

import (
	errors "errors"
	models "mrkresnofatih/gobookapi/models"
)

var bookDb = []models.Book{
	{ID: "sampleID1", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Pages: 134, CreatedAt: 1653299667922},
	{ID: "sampleID2", Title: "The Catcher in the Rye", Author: "J. D. Salinger", Pages: 219, CreatedAt: 1653279653918},
	{ID: "sampleID3", Title: "Invisible Man", Author: "Ralph Ellison", Pages: 198, CreatedAt: 1653229546272},
}

func AddBook(book models.Book) (*models.Book, error) {
	for _, v := range bookDb {
		if v.ID == book.ID {
			e := errors.New("ID already exists!")
			return nil, e
		}
	}

	bookDb = append(bookDb, book)
	return &book, nil
}

func GetBook(id string) (*models.Book, error) {
	for _, v := range bookDb {
		if v.ID == id {
			return &v, nil
		}
	}

	e := errors.New("Book with ID doesn't exist!")
	return nil, e
}

func GetAllBooks() []models.Book {
	return bookDb
}

func RemoveBook(bookId string) (bool, error) {
	foundIndex := -1
	for i, v := range bookDb {
		if v.ID == bookId {
			foundIndex = i
		}
	}
	if foundIndex == -1 {
		e := errors.New("Book not found")
		return false, e
	}

	bookDb = append(bookDb[:foundIndex], bookDb[foundIndex+1:]...)
	return true, nil
}
