package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"catadev.com/stocks/models"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Env adds the dependencies
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

	r := mux.NewRouter()
	r.HandleFunc("/symbols", env.SymbolsHandler)
	r.HandleFunc("/values", env.ValuesHandler)
	r.Use(loggingMiddleware)
	http.Handle("/", r)

	s := &http.Server{
		Addr:         ":9999",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI, r.UserAgent())
		next.ServeHTTP(w, r)
	})
}

// SymbolsHandler handles the index
func (env *Env) SymbolsHandler(w http.ResponseWriter, r *http.Request) {
	symbols, err := env.symbols.All()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json, err := json.Marshal(symbols)
	fmt.Fprintf(w, string(json))
}

// ValuesHandler handles the index
func (env *Env) ValuesHandler(w http.ResponseWriter, r *http.Request) {
	symbols, err := env.values.All("NIO")
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json, err := json.Marshal(symbols)
	fmt.Fprintf(w, string(json))
}
