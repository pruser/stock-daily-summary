package stock

import (
	"fmt"
)

type DailyData struct {
	Symbol string
	Open   string
	Min    string
	Max    string
	Price  string
	Volume string
}

func (d DailyData) String() string {
	return fmt.Sprintf("Symbol:%s,Open:%s,Min:%s,Max:%s,Price:%s,Volume:%s", d.Symbol, d.Open, d.Min, d.Max, d.Price, d.Volume)
}
