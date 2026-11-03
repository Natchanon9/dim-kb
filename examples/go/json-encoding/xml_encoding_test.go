package jsonencoding

import (
	"encoding/xml"
	"fmt"
)

type Book struct {
	XMLName xml.Name `xml:"book"`
	ID      string   `xml:"id,attr"`
	Title   string   `xml:"title"`
}

func Example_xml() {
	b, _ := xml.Marshal(Book{ID: "1", Title: "Go"})
	fmt.Println(string(b))

	var back Book
	xml.Unmarshal(b, &back)
	fmt.Println(back.Title)
	// Output:
	// <book id="1"><title>Go</title></book>
	// Go
}
