package storage

import (
	"currency/internal/core/rest"
	
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

var cfg mysql.Config = mysql.Config{
	User:   "toor",
	Passwd: "toor",
	Net:    "tcp",
	Addr:   "127.0.0.1:3306",
	DBName: "currency_db",
	AllowNativePasswords: true,
}

func AddRecord(record rest.CleanLatestRates) (bool, error) {
	if !ping() {
		panic("DB connection exception")
	}

	result, err := db.Exec("INSERT INTO currency (based, date) VALUES (?, ?)", record.Base, record.Date)
	if err != nil {
        panic(err.Error())
    }
	id, err := result.LastInsertId()
    if err != nil {
        panic(err.Error())
    }

	result, err = db.Exec("INSERT INTO rates (currency, rate, currency_id) VALUES (?, ?, ?), (?, ?, ?), (?, ?, ?)", 
	rest.UAH, record.Rates[rest.UAH], id,
	rest.EUR, record.Rates[rest.EUR], id,
	rest.JPY, record.Rates[rest.JPY], id,
	)
	if err != nil {
        panic(err.Error())
    }
	_, err = result.LastInsertId()
    if err != nil {
        panic(err.Error())
    }

	return true, nil
}

func GetRecordByDate(date string) (rest.CleanLatestRates, error) {
	if !ping() {
		panic("DB connection exception")
	}

	var latestRates rest.CleanLatestRates
	latestRates.Rates = make(map[string]float64)

	var currencyId int32

	currency := db.QueryRow("SELECT * FROM currency WHERE date = ?", date)
    if err := currency.Scan(&currencyId, &latestRates.Base, &latestRates.Date); err != nil {
        if err == sql.ErrNoRows {
            fmt.Println(err.Error())
			return latestRates, err
        }
        panic(err.Error())
    }

	rates, err := db.Query("SELECT currency,rate FROM rates WHERE currency_id = ?", currencyId)
	if err != nil {
		panic(err.Error())
	}
	defer rates.Close()

	for rates.Next() {
		var key string
		var value float64
		if err := rates.Scan(&key, &value); err != nil {
            panic(err.Error())
        }
		latestRates.Rates[key]=value
	}

	return latestRates, nil
}

func ping() bool {
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        fmt.Println(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        fmt.Println(err)
    }
    
	return true
}
