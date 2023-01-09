/*
Write a server program in go to generate a pie chart to compare populations of various states.
Use go-echarts package.
(Data can be adjusted if required without losing actual meaning)
Also mention URL for the request.

Uttar Pradesh:  199,812,341

Kerala:  33,387,677

Karnataka:  61,130,704

Sikkim:  607,688
*/
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	states := []opts.PieData{
		{Name: "Uttar Pradesh", Value: 199812341},
		{Name: "Kerala", Value: 33387677},
		{Name: "Karnataka", Value: 61130704},
		{Name: "Sikkim", Value: 607688},
	}

	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Population of various states"}),
	)

	pie.AddSeries("Population", states)
	err := pie.Render(w)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server Port started at 0238..")
	fmt.Println("Open browser and type(Request URL): http://localhost:0238")
	fmt.Println("By Karthik Banjan")
	log.Fatal(http.ListenAndServe(":0238", nil))
}
