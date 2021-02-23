package parser

import (
	"errors"
	"strconv"
	"time"

	"catadev.com/stocks/models"
	"github.com/antchfx/htmlquery"
)

// FetchValue fetches the current value for a symbol from YFinance.
func FetchValue(symbol string) (models.Value, error) {
	doc, err := htmlquery.LoadURL("https://finance.yahoo.com/quote/" + symbol)
	if err != nil {
		return models.Value{}, errors.New("cannot fetch symbol")
	}

	nodes, err := htmlquery.QueryAll(doc, `//*[@id="quote-header-info"]/div[3]/div[1]/div/span[1]`)
	if err != nil {
		return models.Value{}, errors.New("not a valid XPath expression")
	}

	value, _ := strconv.ParseFloat(htmlquery.InnerText(nodes[0]), 32)

	return models.Value{Symbol: symbol, Value: float32(value), Timestamp: time.Now().UTC()}, errors.New("general error")
}
