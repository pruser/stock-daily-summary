package stock

import (
	"encoding/csv"
	"net/http"
	"sort"
)

type CSVDataSource struct {
	url    string
	parser func([]string) (DailyData, error)
}

func NewCSVDataSource(url string, parser func([]string) (DailyData, error)) *CSVDataSource {
	return &CSVDataSource{url: url, parser: parser}
}

func (s CSVDataSource) Get(symbols []string) ([]DailyData, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := csv.NewReader(resp.Body)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	sort.Strings(symbols)
	result := make([]DailyData, 0)

	for _, id := range symbols {
		for _, v := range records {
			if v[0] == id {
				dv, err := s.parser(v)
				if err != nil {
					return nil, err
				}
				result = append(result, dv)
			}
		}
	}

	return result, nil
}
