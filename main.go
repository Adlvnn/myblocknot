package main

import (
	"fmt"

	"strconv"
)

type Blocknot struct {
	Head string
	Text string
	Len  int
}

type Books struct {
	books [][]string
}

func (b Blocknot) String() string {
	return fmt.Sprintf("Ваша книга %s удачно добавлена в архив.", b.Head)
}

func (b *Books) AddBook(book Blocknot) {
	lengthStr := strconv.Itoa(book.Len)
	b.books = append(b.books, []string{book.Head, book.Text, lengthStr})

}

func main() {

	book := Blocknot{
		Head: "Война и мир",
		Text: "Самая короткая книга в мире",
		Len:  26,
	}
	arhive := Books{}
	arhive.AddBook(book)

	fmt.Println(book)
	fmt.Println(arhive)
}
