package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"

	"catadev.com/stocks/models"
	"catadev.com/stocks/parser"
)

type Env struct {
	symbols models.SymbolModel
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost/stocks?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	env := &Env{
		symbols: models.SymbolModel{DB: db},
	}

	symbols, err := env.symbols.All()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	for _, s := range symbols {
		wg.Add(1)
		go fetch(s.Symbol, &wg)
	}

	wg.Wait()

}

func fetch(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	v, err := parser.FetchValue(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
}
