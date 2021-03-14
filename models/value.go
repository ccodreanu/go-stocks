package models

import (
	"database/sql"
	"time"
)

// Value represents a value for a symbol.
type Value struct {
	Id     int
	Symbol string
	Ts     time.Time
	Value  float32
}

// ValueModel is the wrapper for the DB.
type ValueModel struct {
	DB *sql.DB
}

// All fetches all the values over time for a symbol.
func (m ValueModel) All(symbol string) ([]Value, error) {
	rows, err := m.DB.Query("SELECT id, symbol, value, ts FROM historical_values WHERE symbol = $1", symbol)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var values []Value

	for rows.Next() {
		var value Value

		err := rows.Scan(&value.Id, &value.Symbol, &value.Value, &value.Ts)
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

// AllByDay fetches all the values over time for a symbol.
func (m ValueModel) AllByDay(symbol string) ([]Value, error) {
	rows, err := m.DB.Query("select symbol, date(ts) as date, avg(value) from historical_values where symbol = $1 group by symbol, date", symbol)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var values []Value

	for rows.Next() {
		var value Value

		err := rows.Scan(&value.Symbol, &value.Ts, &value.Value)
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

// Insert adds a value to the DB.
func (m ValueModel) Insert(symbol string, value float32) error {
	rows, err := m.DB.Query(`INSERT INTO historical_values ("symbol", "value", "ts") values ($1, $2, CURRENT_TIMESTAMP)`, symbol, value)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
