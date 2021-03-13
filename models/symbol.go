package models

import "database/sql"

type Symbol struct {
	Id       int
	Symbol   string
	Currency string
}

type SymbolModel struct {
	DB *sql.DB
}

func (m SymbolModel) All() ([]Symbol, error) {
	rows, err := m.DB.Query("SELECT * FROM symbols")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var symbols []Symbol

	for rows.Next() {
		var symbol Symbol

		err := rows.Scan(&symbol.Id, &symbol.Symbol, &symbol.Currency)
		if err != nil {
			return nil, err
		}

		symbols = append(symbols, symbol)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return symbols, nil
}
