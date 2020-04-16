package routes

import (
	"fmt"
	"testing"
)

func TestSymbolNew(t *testing.T) {
	// Save and restore original getSymbolNumbers
	savedSymbolNumbers := getSymbolNumbers
	defer func() { getSymbolNumbers = savedSymbolNumbers }()

	// Save and restore original getSymbolInfo
	savedSymbolInfo := getSymbolInfo
	defer func() { getSymbolInfo = savedSymbolInfo }()

	getSymbolNumbers = func(code string) GlobalQuote {
		quote := GlobalQuote{"PRU", 22.5, 28.2, 22.2, 26.1, 1237986, "10/05/2020", 22.8, "", ""}
		return quote
	}

	getSymbolInfo = func(code string) SymbolSearch {
		info := SymbolSearch{"PRU", "Pru Corporation", "Equity", "United States", "2020-10-05 09:30:20.123455", "2020-10-05 16:30:20.123455", "UTC-5", "Peso Cubano", 1}
		return info
	}

	s := New("PRU")

	cases := []struct {
		got, want string
	}{
		{s.Code, "PRU"},
		{s.Name, "Pru Corporation"},
		{s.Type, "Equity"},
		{s.Region, "United States"},
		{s.Currency, "Peso Cubano"},
		{fmt.Sprintf("%f", s.Price), fmt.Sprintf("%f", 26.1)},
		{fmt.Sprintf("%d", s.Volume), fmt.Sprintf("%d", 1237986)},
		{fmt.Sprintf("%f", s.TodaysHigh.Value), fmt.Sprintf("%f", 28.2)},
		{fmt.Sprintf("%f", s.TodaysLow.Value), fmt.Sprintf("%f", 22.2)},
	}

	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("Error while creating the Symbol. The expected value is %q, but got %q", c.want, c.got)
		}
	}

}
