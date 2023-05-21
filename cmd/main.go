package main

import (
	"currency/internal/transport/rest/router"
	//"currency/internal/core/rest"
	//"fmt"
	//"time"

	//"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
	//"github.com/madflojo/tasks"
)

func main()  {
	r := router.New()

	r.Logger.Fatal(r.Start("127.0.0.1:8585"))
	
	//currencyPuller()
}

/*
func currencyPuller() {
	scheduler := tasks.New()
	defer scheduler.Stop()

	_, err := scheduler.Add(&tasks.Task{
		Interval: 30 * time.Second,
		TaskFunc: func() error {
			rates, err := rest.ReadUsdBased("token", "current date", rest.UAH, rest.EUR, rest.JPY)
			fmt.Println(err.Error())
			fmt.Println(rates.Base)

			return err

		},
	})

	if err != nil {
		fmt.Println(err.Error())
	}
}
*/