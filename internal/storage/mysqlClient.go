package storage

import (
	"currency/internal/core"

	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

var cfg mysql.Config = mysql.Config{
	User:                 "toor",
	Passwd:               "toor",
	Net:                  "tcp",
	Addr:                 "127.0.0.1:3306",
	DBName:               "currency_db",
	AllowNativePasswords: true,
}

func AddRecord(record core.CleanLatestRates) (bool, error) {
	if !ping() {
		panic("DB connection exception")
	}

	result, err := db.Exec("INSERT INTO currency (based, date) VALUES (?, ?)", record.Base, record.Date)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	result, err = db.Exec("INSERT INTO rates (currency, rate, currency_id) VALUES (?, ?, ?), (?, ?, ?), (?, ?, ?)",
		core.UAH, record.Rates[core.UAH], id,
		core.EUR, record.Rates[core.EUR], id,
		core.JPY, record.Rates[core.JPY], id,
	)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	_, err = result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	return true, nil
}

func GetRecordByDate(date string) (core.CleanLatestRates, error) {
	if !ping() {
		panic("DB connection exception")
	}

	var latestRates core.CleanLatestRates
	latestRates.Rates = make(map[string]float64)

	var currencyId int32

	currency := db.QueryRow("SELECT * FROM currency WHERE date = ?", date)
	if err := currency.Scan(&currencyId, &latestRates.Base, &latestRates.Date); err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return latestRates, err
		}
		log.Println(err)
	}

	rates, err := db.Query("SELECT currency,rate FROM rates WHERE currency_id = ?", currencyId)
	if err != nil {
		log.Println(err)
	}
	defer rates.Close()

	for rates.Next() {
		var key string
		var value float64
		if err := rates.Scan(&key, &value); err != nil {
			log.Println(err)
		}
		latestRates.Rates[key] = value
	}

	return latestRates, nil
}

func ping() bool {
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Println(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Println(err)
	}

	return true
}
