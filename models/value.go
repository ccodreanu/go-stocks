package models

import (
	"database/sql"
	"time"
)

// Value represents a value for a symbol.
type Value struct {
	Symbol    string
	Timestamp time.Time
	Value     float32
}

// ValueModel is the wrapper for the DB.
type ValueModel struct {
	DB *sql.DB
}

// All fetches all the values over time for a symbol.
func (m ValueModel) All(symbol string) ([]Value, error) {
	rows, err := m.DB.Query("SELECT * FROM values WHERE symbol = $1", symbol)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var values []Value

	for rows.Next() {
		var value Value

		err := rows.Scan(&value.Symbol, &value.Value, &value.Timestamp)
		if err != nil {
			return nil, err
		}

		values = append(values, value)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return values, nil
}

func (m ValueModel) Insert(symbol string, value float32) error {
	rows, err := m.DB.Query(`INSERT INTO values ("symbol", "value", "timestamp") values ($1, $2, CURRENT_TIMESTAMP)`, symbol, value)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
