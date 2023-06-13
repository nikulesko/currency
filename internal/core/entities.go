package core

type LatestRates struct {
	Motd    AddMessage
	Success bool
	Base    string
	Date    string
	Rates   map[string]float64
}

func (l *LatestRates) Clean() CleanLatestRates {
	return CleanLatestRates{
		l.Base,
		l.Date,
		l.Rates,
	}
}

type CleanLatestRates struct {
	Base  string
	Date  string
	Rates map[string]float64
}

type AddMessage struct {
	Msg string
	Url string
}

func EmptyLatestRates() CleanLatestRates {
	return CleanLatestRates{"", "", nil}
}
