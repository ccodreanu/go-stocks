package models

import (
	"database/sql"
	"time"
)

type Value struct {
	Symbol    string
	Timestamp time.Time
	Value     float32
}

type ValueModel struct {
	DB *sql.DB
}

func (m ValueModel) All(symbol Symbol) ([]Value, error) {
	rows, err := m.DB.Query("SELECT * FROM values WHERE symbol = $1", symbol.Symbol)
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
