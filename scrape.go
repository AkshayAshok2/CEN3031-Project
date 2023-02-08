package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"strconv"
)

type ForeclosureProperty struct {
	address       string `json: "Address"`
	price         string `json: "Price"`
	bed           string `json: "Bed"`
	bath          string `json: "Bath"`
	squareFootage string `json: "Square Footage"`
}

func connectAndLoadData(w http.ResponseWriter, r *http.Request) {

	property := ForeclosureProperty{
		address: "123",
		price: "1000",
		bed: "3",
		bath: "4",
		squareFootage: "2000",
	}
	var dsn = "go:Gators123@tcp(cen3031-project.mysql.database.azure.com)/listings?charset=utf8&parseTime=True&Ioc=Local"
	var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.Create(property)
	if err := db.Create(property).Error; err != nil {
		log.Fatalln((err))
	}

	json.NewEncoder(w).Encode(property)
	fmt.Println("Fields Added", property)

}


// whatever is inside main runs
func main() {
	http.HandleFunc("/CreateStuff",connectAndLoadData)
	log.Fatal(http.ListenAndServe(":8080", nil))
	/*c := colly.NewCollector(colly.AllowedDomains("alachua.realforeclose.com"))

	/*for i := 1; i < 31; i++ {
	if i < 10{
		var date string = "0" + strconv.Itoa(i)
	} else {
		var date string = strconv.Itoa(i)
	}
	*/

	//Initializes the days and strings to loop through the urls and gather foreclosure data for each day
	//var month string = "02"
	//var day string = "09"

	/*for i := 1; i < 13; i++ {
		if i < 10{
			var month string = "0" + strconv.Itoa(i)
		} else {
			var month string = strconv.Itoa(i)
		}

		for j := 1; j < 28; i++ {
			if j < 10{
				var day string = "0" + strconv.Itoa(i)
			} else {
				var day string = strconv.Itoa(i)
			}
		}
	}
	

	scrapeURL := "https://alachua.realforeclose.com/index.cfm?zaction=AUCTION&Zmethod=PREVIEW&AUCTIONDATE=" + month + "/" + day + "/2022"

	c.OnHTML("div.Sub_Title", func(e *colly.HTMLElement) {

		fmt.Println(e.Text)

		//fmt.Println(e.ChildText(e.Text))
		//fmt.Println(e.ChildText("td.AD_DTA"))

	})

	c.OnHTML("div.Auct_Area", func(e *colly.HTMLElement) {
		selection := e.DOM
		fmt.Println(selection.Find("div.AUCTION_DETAILS > table > tbody > tr:nth-child(3) > td").Text())
	})

	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		fmt.Print("Childe: ")
		fmt.Println(e.ChildText)

	})

	var scrapeAddress string
	var scrapeBath string
	var scrapeBed string
	var scrapePrice string
	var scrapeSquareFootage string

	var property ForeclosureProperty
	property.address = scrapeAddress
	property.bath = scrapeBath
	property.bed = scrapeBed
	property.price = scrapePrice
	property.squareFootage = scrapeSquareFootage

	c.Visit(scrapeURL)
	/*
	   	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "go", "Gators123", host, port, database)
	   	// Open the connection
	   	db, err := sql.Open("mysql", dsn)
	   	if err != nil {
	       		log.Fatalf("impossible to create the connection: %s", err)
	   	}
	   	defer db.Close()
	*/
}

//}
