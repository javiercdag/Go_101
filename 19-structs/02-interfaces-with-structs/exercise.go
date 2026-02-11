package structsinterfaces

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

import "fmt"

type Printable interface {
	Info() string
	PageNum() int
}

type Book struct {
	Author string
	Title  string
	Pages  int
}

type Magazine struct {
	Title string
	Issue string
	Pages int
}

func NewBook(author, title string, pages int) *Book {
	return &Book{
		Author: author,
		Title:  title,
		Pages:  pages,
	}
}

func NewMagazine(title, issue string, pages int) *Magazine {
	return &Magazine{
		Title: title,
		Issue: issue,
		Pages: pages,
	}
}

func (b *Book) Info() string {
	return fmt.Sprintf("%s by %s", b.Title, b.Author)
}

func (b *Book) PageNum() int {
	return b.Pages
}

func (m *Magazine) Info() string {
	return fmt.Sprintf("%s, Issue: %s", m.Title, m.Issue)
}

func (m *Magazine) PageNum() int {
	return m.Pages
}