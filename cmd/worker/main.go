package main

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"

	"catadev.com/stocks/models"
	"catadev.com/stocks/parser"
)

type Env struct {
	symbols models.SymbolModel
	values  models.ValueModel
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost/stocks?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	env := &Env{
		symbols: models.SymbolModel{DB: db},
		values:  models.ValueModel{DB: db},
	}

	symbols, err := env.symbols.All()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	for _, s := range symbols {
		wg.Add(1)
		go fetch(s.Symbol, env, &wg)
	}

	wg.Wait()

}

func fetch(s string, env *Env, wg *sync.WaitGroup) {
	defer wg.Done()

	v, err := parser.FetchValue(s)
	if err != nil {
		log.Printf("cannot fetch %s", s)
		return
	}

	env.values.Insert(s, v.Value)
}
