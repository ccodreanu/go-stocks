package main

import (
	"fmt"
	"log"

	"github.com/antchfx/htmlquery"
)

func main() {
	doc, err := htmlquery.LoadURL("https://finance.yahoo.com/quote/VUSA.AS")
	if err != nil {
		log.Fatalln("error getting stock")
	}

	nodes, err := htmlquery.QueryAll(doc, `//*[@id="quote-header-info"]/div[3]/div[1]/div/span[1]`)
	if err != nil {
		panic(`not a valid XPath expression.`)
	}

	fmt.Println(htmlquery.InnerText(nodes[0]))
}
