package main

import (
	"database/sql"
	"fmt"
	"log"

	"catadev.com/stocks/models"
	"catadev.com/stocks/parser"

	_ "github.com/lib/pq"
)

type Env struct {
	symbols models.SymbolModel
	values  models.ValueModel
}

func main() {
	current, err := parser.FetchValue("VUSA.AS")

	fmt.Println(current)

	db, err := sql.Open("postgres", "postgres://postgres:password@localhost/stocks?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	env := &Env{
		symbols: models.SymbolModel{DB: db},
		values:  models.ValueModel{DB: db},
	}

	env.symbolsAll()

	env.valuesAll("VUSA.AS")

	env.values.Insert("NIO", 25.5)
}

func (env *Env) symbolsAll() {
	// Execute the SQL query by calling the All() method.
	symbols, err := env.symbols.All()
	if err != nil {
		log.Println(err)
		return
	}

	for _, symbol := range symbols {
		fmt.Printf("%s, %s\n", symbol.Symbol, symbol.Currency)
	}
}

func (env *Env) valuesAll(symbol string) {
	// Execute the SQL query by calling the All() method.
	values, err := env.values.All(symbol)
	if err != nil {
		log.Println(err)
		return
	}

	for _, value := range values {
		fmt.Printf("%s, %f, %s\n", value.Symbol, value.Value, value.Ts)
	}
}
