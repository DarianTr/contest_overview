package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func GetUsaco() {
	url := "http://usaco.org/"
	c := colly.NewCollector()

	// Find and visit all links
	//need to make 2023-2024 with time.Year()
	c.OnHTML("div.panel", func(e *colly.HTMLElement) {
		if e.ChildText("h2") == "2023-2024 Schedule" {
			content := e.Text
			lines := strings.Split(content, "\n")
			fmt.Println(lines[3])
			fmt.Println(lines[4])
			fmt.Println(lines[5])
			fmt.Println(lines[6])
		}
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}
