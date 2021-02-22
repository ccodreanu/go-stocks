package main

import (
	"database/sql"
	"fmt"
	"log"

	"catadev.com/stocks/models"
	"github.com/antchfx/htmlquery"

	_ "github.com/lib/pq"
)

type Env struct {
	symbols models.SymbolModel
}

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

	db, err := sql.Open("postgres", "postgres://postgres:password@localhost/stocks?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	env := &Env{
		symbols: models.SymbolModel{DB: db},
	}

	env.symbolsAll()
}

func (env *Env) symbolsAll() {
	// Execute the SQL query by calling the All() method.
	symbols, err := env.symbols.All()
	if err != nil {
		log.Println(err)
		return
	}

	for _, symbol := range symbols {
		fmt.Printf("%s, %s", symbol.Symbol, symbol.Currency)
	}
}
