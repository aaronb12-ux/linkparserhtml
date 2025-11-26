package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text []string
}

func getLinks(f *os.File) []Link {

	defer f.Close() //make sure files closes when function ends
	doc, err := html.Parse(f)
	var Links []Link

	if err != nil {
		log.Fatal("error parsing file. exiting program")
	}

	for n := range doc.Descendants() {

		if n.Data == "a" && len(n.Attr) != 0 {

				var l Link 
				l.Href = n.Attr[0].Val
				
				for child := range n.Descendants() {
	
					if child.Type == html.TextNode {
						l.Text = append(l.Text, child.Data)
					}
									
				}

				strings.Join(l.Text, "")
				Links = append(Links, l)		
		}
	
	}
		return Links
}

func main() {

	f, err := os.Open("ex2.html")

	if err != nil {
		log.Fatal("unable to open file")
	}

	Links := getLinks(f)   
	fmt.Println(Links)

}