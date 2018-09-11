package stock

type DataSource interface {
	Get(symbols []string) ([]DailyData, error)
}
