package main

import (
	//"io"
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

func main() {

	f, err := os.Open("ex2.html")

	if err != nil {
		log.Fatalf("unable to open file")
	}

	var Links []Link

	doc, err := html.Parse(f)

	for n := range doc.Descendants() {

		if n.Data == "a" && len(n.Attr) != 0 {

				var l Link 
				fmt.Println("THE ATTRIBUTE OF THE PARENT IS", n.Attr)
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
    

	fmt.Println(Links[0].Text)

}