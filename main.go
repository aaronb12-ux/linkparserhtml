package main

import (
	//"io"
	"log"
	"os"
	"fmt"
	"golang.org/x/net/html"
)




func main() {

	f, err := os.Open("ex1.html")

	if err != nil {
		log.Fatalf("unable to open file")
	}

	doc, err := html.Parse(f)

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "href /other-page"{
			fmt.Println(n.Data)
		}
	}

	
}