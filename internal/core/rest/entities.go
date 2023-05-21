package rest

type HistoricalRates struct {
	Date  string
	Base  string
	Rates map[string]string
}

func EmptyHistoricalRates() HistoricalRates {
	return HistoricalRates{"", "", nil}
}
