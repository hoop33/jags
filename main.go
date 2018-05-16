package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get("https://www.jaguars.com/team/players-roster/")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Failed to get roster from team site: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Player,Number,Position,Height,Weight,Age,Experience,College")
	doc.Find("[summary='Roster']").Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		name, exists := s.Find("a").Attr("title")
		if exists {
			fmt.Printf("%s", name)
			for i := 1; i <= 7; i++ {
				fmt.Printf(",%s", strings.TrimSpace(s.Find("td").Eq(i).Text()))
			}
			fmt.Println()
		}
	})
}
