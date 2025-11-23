package main

import (
	//"io"
	"log"
	"os"
	"fmt"
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

		
		fmt.Println("the element node is", html.ElementNode)
		fmt.Println("the n.data is", n.Data)
		fmt.Println("the attribute is", n.Attr)
		fmt.Println("the type is", n.Type)
		fmt.Println()


		if n.Data == "a"{ //need to find a links then search within those a links...
			
			var link Link
			if len(n.Attr[0].Val) != 0 {
				
				l := n.Attr[0].Val
				link.Href = l

			}
			//now search within this node?
			fmt.Println(n.Type, html.ElementNode)
			if n.Type == 3 && html.ElementNode == 3 {
				link.Text = append(link.Text, n.Data)
			}

			Links = append(Links, link)
		}	
	}

	fmt.Println(Links)

}