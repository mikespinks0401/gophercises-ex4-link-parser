package main

import (

	"golang.org/x/net/html"
	"fmt"
	"os"
)

/*
Turn This
<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>

Into This
Link{
  Href: "/dog",
  Text: "Something in a span Text not in a span Bold text!",
}
*/

type Link struct{
	Href string
	Text string
}



type Links []Link

func main(){
	var links Links
	file, err  := os.Open("examples/ex2.html")
	if err != nil{
		fmt.Println(err.Error())
	}
	defer file.Close()
	parsed, _ := html.Parse(file)
	var f func(*html.Node)
	// f = func(n *html.Node) {
	// 	if n.Type == html.ElementNode && n.Data == "a" {
	// 		for _, a := range n.Attr{

	// 			if a.Key == "href" {
	// 				var link Link
	// 				link.Href = a.Val
	// 				link.Text = strings.TrimSpace(n.FirstChild.Data)
	// 				links = append(links, link)
	// 			}
	// 		}
	// 	}
	// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 		f(c)
	// 	}
	// }
	f = func(n *html.Node){
		if n == nil {
			return
		}	
		if n.Type == html.ElementNode && n.Data == "a"{
			for _,val := range n.Attr{
				var link Link
				if val.Key == "href"{
					link.Href = val.Val
					link.Text = n.FirstChild.Data
					links = append(links, link)
				}
				

			}
		}
		b := n.FirstChild
		for b != nil{
			f(b)
			b = b.NextSibling

		}
	}
	f(parsed)
	for _,val := range links{
		fmt.Printf("Href:%s\n", val.Href)
		fmt.Printf("Text:%s\n", val.Text)
		
	}
}